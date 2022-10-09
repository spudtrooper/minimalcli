package handler

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/or"
	"github.com/yosssi/gohtml"
)

type SourceLocation struct {
	Uri  string `json:"uri"`
	Line int    `json:"line"`
}

// TODO: github-specfic hash
func (s SourceLocation) URI() string { return fmt.Sprintf("%s#L%d", s.Uri, s.Line) }

type HandlerSourceLocation struct {
	Handler string
	Loc     SourceLocation
}

type Section struct {
	Prefix     string
	Title      string
	Handlers   []Handler
	FooterHTML string
}

//go:generate genopts --function AddSection indexName:string editName:string footerHTML:string sourceLinks handlersFiles:[]string handlersFilesRoot:string sourceLinkURIRoot:string formatHTML serializedSourceLocations:[]byte
func AddSection(ctx context.Context, mux *http.ServeMux, hs []Handler, prefix, title string, optss ...AddSectionOption) (*Section, error) {
	opts := MakeAddSectionOptions(optss...)

	if opts.SourceLinks() {
		if len(opts.HandlersFiles()) == 0 && len(opts.SerializedSourceLocations()) == 0 {
			return nil, errors.Errorf("if sourceLinks is true, HandlersFiles or SerializedSourceLocationsshould be set; you probably made a mistake")
		}
	} else {
		if len(opts.HandlersFiles()) > 0 || len(opts.SerializedSourceLocations()) > 0 {
			return nil, errors.Errorf("if sourceLinks isn't set, HandlersFiles or SerializedSourceLocationsshould should not set; you probably made a mistake")
		}
	}

	indexTitle := title // TODO: Don't need to alias. Convert indexTitle -> title
	indexName := opts.IndexName()
	editName := or.String(opts.EditName(), "_edit")
	footerHTML := opts.FooterHTML()
	handlersFiles := opts.HandlersFiles()
	handlersFilesRoot := opts.HandlersFilesRoot()
	sourceLinkURIRoot := opts.SourceLinkURIRoot()
	formatHTML := opts.FormatHTML()

	handleFunc := func(route string, fn func(w http.ResponseWriter, req *http.Request)) {
		log.Printf("adding route %s", route)
		mux.HandleFunc(route, fn)
	}
	routesToHandlers := map[string]*handler{}
	for _, h := range hs {
		h := h.(*handler)
		if h.cliOnly {
			continue
		}
		route := fmt.Sprintf("/%s/%s", prefix, strings.ToLower(h.name))
		routesToHandlers[route] = h
		handleFunc(route, func(w http.ResponseWriter, req *http.Request) {
			handle(ctx, h, w, req)
		})
	}

	var handlerToSource map[string]SourceLocation
	if opts.SourceLinks() && len(handlersFiles) > 0 {
		m, err := findHandlerSourceLocations(handlersFiles, handlersFilesRoot, sourceLinkURIRoot)
		if err != nil {
			return nil, errors.Errorf("failed to find source locations in files %s: %w", handlersFiles, err)
		}
		handlerToSource = m
		if len(handlerToSource) != len(hs) {
			log.Printf("warning: when generating source links, we found %d handlers and %d source locations", len(hs), len(handlerToSource))
		}
	}

	if opts.SourceLinks() && len(opts.SerializedSourceLocations()) > 0 {
		var locs []HandlerSourceLocation
		if err := json.Unmarshal(opts.SerializedSourceLocations(), &locs); err != nil {
			return nil, errors.Errorf("failed to unmarshal source locations: %w", err)
		}
		m := map[string]SourceLocation{}
		for _, loc := range locs {
			m[loc.Handler] = loc.Loc
		}
		handlerToSource = m
	}

	{
		index, err := genIndex(indexTitle, prefix, editName, routesToHandlers, footerHTML, handlerToSource, formatHTML)
		if err != nil {
			return nil, errors.Errorf("error generating index page: %w", err)
		}
		handleFunc(fmt.Sprintf("/%s/%s", prefix, indexName), func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprint(w, index)
		})
	}

	{
		routesToEdits := map[string]string{}
		for route, h := range routesToHandlers {
			var sourceURI string
			if s, ok := handlerToSource[h.name]; ok {
				sourceURI = s.URI()
			}
			edit, err := genEdit(indexTitle, route, prefix, indexName, h, formatHTML, sourceURI)
			if err != nil {
				return nil, errors.Errorf("error generating edit page for %s: %w", h.name, err)
			}
			routesToEdits[route] = edit

		}
		handleFunc(fmt.Sprintf("/%s/%s", prefix, editName), func(w http.ResponseWriter, req *http.Request) {
			route, ok := getStringURLParamOrDie(w, req, "route")
			if !ok {
				return
			}
			edit, ok := routesToEdits[route]
			if !ok {
				respondWithErrorString(w, req, "no edit found for route %s", route)
				return
			}
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprint(w, edit)
		})
	}

	sec := &Section{
		Prefix:     prefix,
		Title:      title,
		Handlers:   hs,
		FooterHTML: opts.FooterHTML(),
	}
	return sec, nil
}

//go:generate genopts --function GenIndex title:string footerHTML:string formatHTML route:string
func GenIndex(ctx context.Context, mux *http.ServeMux, secs []Section, optss ...GenIndexOption) error {
	opts := MakeGenIndexOptions(optss...)

	title := or.String(opts.Title(), "Index")
	route := or.String(opts.Route(), "/")

	handleFunc := func(route string, fn func(w http.ResponseWriter, req *http.Request)) {
		log.Printf("adding route %s", route)
		mux.HandleFunc(route, fn)
	}

	all, err := genAllIndex(title, opts.FooterHTML(), opts.FormatHTML(), secs)
	if err != nil {
		return err
	}

	handleFunc(route, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, all)
	})
	return nil
}

var (
	//go:embed tmpl/all.html
	allHTMLTemplate string
)

func genAllIndex(title, footerHTML string, format bool, secs []Section) (string, error) {
	type Sec struct {
		Route string
		Title string
	}
	var sections []Sec
	for _, s := range secs {
		sections = append(sections, Sec{
			Route: s.Prefix,
			Title: s.Title,
		})
	}
	sort.Slice(sections, func(i, j int) bool {
		return sections[i].Route < sections[j].Route
	})
	var data = struct {
		Title      string
		Sections   []Sec
		FooterHTML string
	}{
		Title:      title,
		Sections:   sections,
		FooterHTML: footerHTML,
	}

	var buf bytes.Buffer
	if err := renderTemplate(&buf, allHTMLTemplate, "all", data); err != nil {
		return "", err
	}
	s := buf.String()
	if format {
		s = gohtml.Format(s)
	}
	return s, nil
}

//go:generate genopts --function AddHandlers --extends AddSection indexTitle:string prefix:string
func AddHandlers(ctx context.Context, mux *http.ServeMux, hs []Handler, optss ...AddHandlersOption) error {
	opts := MakeAddHandlersOptions(optss...)
	if _, err := AddSection(ctx, mux, hs, opts.Prefix(), opts.IndexTitle(), opts.ToAddSectionOptions()...); err != nil {
		return err
	}
	return nil
}

func handle(ctx context.Context, h *handler, w http.ResponseWriter, req *http.Request) {
	if !strings.EqualFold(req.Method, h.method) {
		respondWithErrorString(w, req, "method %s not supported, only %s supported", req.Method, h.method)
		return
	}
	evalCtx := &serverEvalContext{ctx, w, req}
	res, err := h.fn(evalCtx)
	if err != nil {
		respondWithError(w, req, err)
		return
	}
	respondWithJSON(req, w, res)
}

var (
	// NewHandlerFromHandlerFn("SaveRawRestaurantDetailsFromID"
	newHandlerFromHandlerFnRE = regexp.MustCompile(`NewHandlerFromHandlerFn\("([^"]+)"`)
	// NewHandler("AddRestaurantsToSearchByURIs"
	newHandlerRE = regexp.MustCompile(`NewHandler\("([^"]+)"`)
)

func FindHandlerSourceLocations(handlersFiles []string, sourceLinkURIRoot string) (map[string]SourceLocation, error) {
	return findHandlerSourceLocations(handlersFiles, "", sourceLinkURIRoot)
}

func findHandlerSourceLocations(handlersFiles []string, handlersFilesRoot string, sourceLinkURIRoot string) (map[string]SourceLocation, error) {
	res := map[string]SourceLocation{}
	root := sourceLinkURIRoot
	if strings.HasSuffix(root, "/") {
		root = strings.TrimRight(root, "/")
	}
	if strings.HasPrefix(root, "/") {
		root = strings.TrimLeft(root, "/")
	}
	for _, f := range handlersFiles {
		log.Printf("reading %s...", f)
		c, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, err
		}
		startLen := len(res)
		for i, line := range strings.Split(string(c), "\n") {
			if m := newHandlerFromHandlerFnRE.FindStringSubmatch(line); len(m) == 2 {
				h := m[1]
				f = strings.TrimPrefix(f, handlersFilesRoot)
				// Can't use path.Join, because it will remove the leading double slashes
				uri := root + "/" + f
				loc := SourceLocation{uri, i + 1}
				res[h] = loc
				continue
			}
			if m := newHandlerRE.FindStringSubmatch(line); len(m) == 2 {
				h := m[1]
				uri := path.Join(sourceLinkURIRoot, f)
				loc := SourceLocation{uri, i + 1}
				res[h] = loc
				continue
			}
		}
		stopLen := len(res)
		log.Printf("found %d source locations from %s...", stopLen-startLen, f)
	}
	if len(res) == 0 {
		return nil, errors.Errorf("no handlers found in any of the files: %s", handlersFiles)
	}
	return res, nil
}

//go:embed tmpl/index.html
var indexHTMLTemplate string

func genIndex(title, prefix, editName string, routesToHandlers map[string]*handler, footerHTML string, handlerToSource map[string]SourceLocation, format bool) (string, error) {
	type route struct {
		Route     string
		HasParams bool
		SourceURI string
	}
	var routePaths []string
	for r := range routesToHandlers {
		routePaths = append(routePaths, r)
	}
	sort.Strings(routePaths)
	var routes []route
	for _, r := range routePaths {
		route := route{
			Route:     r,
			HasParams: !routesToHandlers[r].metadata.Empty(),
		}
		h := routesToHandlers[r]
		if s, ok := handlerToSource[h.name]; ok {
			route.SourceURI = s.URI()
		}
		routes = append(routes, route)
	}
	var data = struct {
		Title      string
		Prefix     string
		EditName   string
		Routes     []route
		FooterHTML string
	}{
		Title:      title,
		Prefix:     prefix,
		EditName:   editName,
		Routes:     routes,
		FooterHTML: footerHTML,
	}

	var buf bytes.Buffer
	if err := renderTemplate(&buf, indexHTMLTemplate, "index", data); err != nil {
		return "", err
	}
	s := buf.String()
	if format {
		s = gohtml.Format(s)
	}
	return s, nil
}

func renderTemplate(buf io.Writer, t string, name string, data interface{}) error {
	tmpl, err := template.New(name).Parse(strings.TrimSpace(t))
	if err != nil {
		return err
	}
	if err := tmpl.Execute(buf, data); err != nil {
		return err
	}
	return nil
}

//go:embed tmpl/edit.html
var editHTMLTemplate string

func genEdit(title, route, prefix, indexName string, h *handler, format bool, sourceURI string) (string, error) {
	type form struct {
		Name     string
		Required bool
		Type     HandlerMetadataParamType
		Default  string
	}
	var forms []form
	for _, p := range h.metadata.Params {
		f := form{
			Name:     p.Name,
			Type:     p.Type,
			Required: p.Required,
			Default:  p.Default,
		}
		forms = append(forms, f)
	}
	var data = struct {
		Title     string
		Route     string
		Prefix    string
		IndexName string
		SourceURI string
		Forms     []form
	}{
		Title:     title,
		Route:     route,
		Prefix:    prefix,
		IndexName: indexName,
		SourceURI: sourceURI,
		Forms:     forms,
	}

	var buf bytes.Buffer
	if err := renderTemplate(&buf, editHTMLTemplate, "edit", data); err != nil {
		return "", err
	}

	s := buf.String()
	if format {
		s = gohtml.Format(s)
	}
	return s, nil
}
