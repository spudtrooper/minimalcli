package handler

import (
	"bytes"
	"context"
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
)

type sourceLocation struct {
	uri  string
	line int
}

//go:generate genopts --function CreateHandler indexTitle:string prefix:string indexName:string editName:string footerHTML:string sourceLinks handlersFiles:[]string sourceLinkURIRoot:string
func CreateHandler(ctx context.Context, hs []Handler, optss ...CreateHandlerOption) (*http.ServeMux, error) {
	opts := MakeCreateHandlerOptions(optss...)

	if len(opts.HandlersFiles()) > 1 && !opts.SourceLinks() {
		return nil, errors.Errorf("if handlersFile is set, sourceLinks should be true; you probably made a mistake")
	}
	if len(opts.HandlersFiles()) == 0 && opts.SourceLinks() {
		return nil, errors.Errorf("if sourceLinks is true, handlersFile should be set; you probably made a mistake")
	}

	indexTitle := or.String(opts.IndexTitle(), "API")
	indexName := opts.IndexName()
	editName := or.String(opts.EditName(), "_edit")
	prefix := or.String(opts.Prefix(), "api")
	footerHTML := opts.FooterHTML()
	handlersFiles := opts.HandlersFiles()
	sourceLinkURIRoot := opts.SourceLinkURIRoot()

	mux := http.NewServeMux()

	log.Printf("Initializing server...")

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

	var handlerToSource map[string]sourceLocation
	if len(handlersFiles) > 0 {
		m, err := findHandlerSourceLocations(handlersFiles, sourceLinkURIRoot)
		if err != nil {
			return nil, err
		}
		handlerToSource = m
		if len(handlerToSource) != len(hs) {
			log.Printf("warning: when generating source links, we found %d handlers and %d source locations", len(hs), len(handlerToSource))
		}
	}

	handleFunc(fmt.Sprintf("/%s/%s", prefix, indexName), func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		s, err := genIndex(indexTitle, prefix, editName, routesToHandlers, footerHTML, handlerToSource)
		if err != nil {
			respondWithError(w, req, err)
			return
		}
		fmt.Fprint(w, s)
	})

	handleFunc(fmt.Sprintf("/%s/%s", prefix, editName), func(w http.ResponseWriter, req *http.Request) {
		route, ok := getStringURLParamOrDie(w, req, "route")
		if !ok {
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		s, err := genEdit(indexTitle, route, prefix, indexName, routesToHandlers)
		if err != nil {
			respondWithError(w, req, err)
			return
		}
		fmt.Fprint(w, s)
	})

	return mux, nil
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
	return
}

var (
	// NewHandlerFromParams("SaveRawRestaurantDetailsFromID"
	newHandlerFromParamsRE = regexp.MustCompile(`NewHandlerFromParams\("([^"]+)"`)
	// NewHandler("AddRestaurantsToSearchByURIs"
	newHandlerRE = regexp.MustCompile(`NewHandler\("([^"]+)"`)
)

func findHandlerSourceLocations(handlersFiles []string, sourceLinkURIRoot string) (map[string]sourceLocation, error) {
	res := map[string]sourceLocation{}
	for _, f := range handlersFiles {
		c, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, err
		}
		for i, line := range strings.Split(string(c), "\n") {
			if m := newHandlerFromParamsRE.FindStringSubmatch(line); len(m) == 2 {
				h := m[1]
				uri := path.Join(sourceLinkURIRoot, f)
				loc := sourceLocation{uri, i + 1}
				res[h] = loc
				continue
			}
			if m := newHandlerRE.FindStringSubmatch(line); len(m) == 2 {
				h := m[1]
				uri := path.Join(sourceLinkURIRoot, f)
				loc := sourceLocation{uri, i + 1}
				res[h] = loc
				continue
			}
		}
	}
	if len(res) == 0 {
		return nil, errors.Errorf("no handlers found in any of the files: %s", handlersFiles)
	}
	return res, nil
}

func genIndex(title, prefix, editName string, routesToHandlers map[string]*handler, footerHTML string, handlerToSource map[string]sourceLocation) (string, error) {
	const t = `
{{$prefix := .Prefix}}
{{$editName := .EditName}}
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">

		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css" rel="stylesheet"
			integrity="sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC" crossorigin="anonymous">
		<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js"
			integrity="sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM" crossorigin="anonymous">
		</script>

		<link rel="apple-touch-icon" sizes="180x180" href="/img/apple-touch-icon.png">
		<link rel="icon" type="image/png" sizes="32x32" href="/img/favicon-32x32.png">
		<link rel="icon" type="image/png" sizes="16x16" href="/img/favicon-16x16.png">
		<link rel="manifest" href="/img/site.webmanifest">	
	</head>
	<body>
		<div class="container-fluid">
			<h1>{{.Title}}</h1>
			<ul>
				{{range .Routes}}
					<li>
						<a href="{{.Route}}">{{.Route}}</a>					
						{{if .HasParams}}
							(<a href="/{{$prefix}}/{{$editName}}?route={{.Route}}">edit</a>)
						{{end}}
						{{if ne .SourceURI ""}}
							[<a target="_" href="{{.SourceURI}}">source</a>]
						{{end}}
					</li>
				{{end}}
			</ul>
			<div>
				{{.FooterHTML}}
			</div>
		</div>
	</body>
</html>
	`
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
			sourceURI := fmt.Sprintf("%s#L%d", s.uri, s.line)
			route.SourceURI = sourceURI
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
	if err := renderTemplate(&buf, t, "index", data); err != nil {
		return "", err
	}
	return buf.String(), nil
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

func genEdit(title, route, prefix, indexName string, routesToHandlers map[string]*handler) (string, error) {
	const t = `
{{$prefix := .Prefix}}
{{$indexName := .IndexName}}
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">

		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css" rel="stylesheet"
			integrity="sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC" crossorigin="anonymous">
		<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js"
			integrity="sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM" crossorigin="anonymous">
		</script>

		<link rel="apple-touch-icon" sizes="180x180" href="/img/apple-touch-icon.png">
		<link rel="icon" type="image/png" sizes="32x32" href="/img/favicon-32x32.png">
		<link rel="icon" type="image/png" sizes="16x16" href="/img/favicon-16x16.png">
		<link rel="manifest" href="/img/site.webmanifest">	
	</head>
	<body>
		<div class="container-fluid">
			<h1>{{.Title}}</h1>
			<h2>{{.Route}}</h2>
			<form class="row g-1 needs-validation" action="{{.Route}}" method="GET">
				{{range $i, $f := .Forms}}
					{{if eq $f.Type "string"}}
						<div class="col-md">
							<label for="validationCustom{{$i}}" class="form-label">{{$f.Name}} <code>string</code>
							{{if $f.Required}} (required){{end}}
							</label>
							<input type="text" name="{{$f.Name}}" class="form-control" id="validationCustom{{$i}}" value=""
								{{if $f.Required}}required{{end}}
							>
					{{end}}
					{{if eq $f.Type "duration"}}
						<div class="col-md">
							<label for="validationCustom{{$i}}" class="form-label">{{$f.Name}} <code>time.Duration</code>
							{{if $f.Required}} (required){{end}}
							</label>
							<input type="text" name="{{$f.Name}}" class="form-control" id="validationCustom{{$i}}" value=""
								{{if $f.Required}}required{{end}}
							>
					{{end}}					
					{{if eq $f.Type "duration"}}
						<div class="col-md">
							<label for="validationCustom{{$i}}" class="form-label">{{$f.Name}} <code>time.Time</code>
							{{if $f.Required}} (required){{end}}
							</label>
							<input type="text" name="{{$f.Name}}" class="form-control" id="validationCustom{{$i}}" value=""
								{{if $f.Required}}required{{end}}
							>
					{{end}}					
					{{if eq $f.Type "int"}}
						<div class="col-md">
							<label for="validationCustom{{$i}}" class="form-label">{{$f.Name}} <code>int</code>
							{{if $f.Required}} (required){{end}}
							</label>
							<input type="number" name="{{$f.Name}}" class="form-control" id="validationCustom{{$i}}" value=""
								{{if $f.Required}}required{{end}}
							>
					{{end}}
					{{if eq $f.Type "float32"}}
						<div class="col-md">
							<label for="validationCustom{{$i}}" class="form-label">{{$f.Name}} <code>float32</code>
							{{if $f.Required}} (required){{end}}
							</label>
							<input type="number" name="{{$f.Name}}" class="form-control" id="validationCustom{{$i}}" value=""
								{{if $f.Required}}required{{end}}
							>
					{{end}}
					{{if eq $f.Type "bool"}}
						<div class="col-md">
							<label for="validationCustom{{$i}}" class="form-label">{{$f.Name}} <code>bool</code>
							{{if $f.Required}} (required){{end}}
							</label>
							<select name="{{$f.Name}}" class="form-select" id="validationCustom{{$i}}">
								<option selected disabled value="">Choose...</option>
								<option></option>
								<option value="true">True</option>
								<option value="false">False</option>
							</select>
							<div class="invalid-feedback">
								Please select a valid {{$f.Name}}.
							</div>
						</div>					
					{{end}}
				{{end}}
				<br/>
				<div class="col-12">
					<button class="btn btn-primary" type="submit">Submit</button>
				</div>
				</form>
			<br/>
			<a href="/{{$prefix}}/{{$indexName}}">Home</a>
		</div>
	</body>
</html>
	`
	type form struct {
		Name     string
		Required bool
		Type     HandlerMetadataParamType
	}
	var forms []form
	for r, h := range routesToHandlers {
		if r != route {
			continue
		}
		if h.metadata.Empty() {
			continue
		}
		for _, p := range h.metadata.Params {
			f := form{
				Name:     p.Name,
				Type:     p.Type,
				Required: p.Required,
			}
			forms = append(forms, f)
		}
		break
	}
	var data = struct {
		Title     string
		Route     string
		Prefix    string
		IndexName string
		Forms     []form
	}{
		Title:     title,
		Route:     route,
		Prefix:    prefix,
		IndexName: indexName,
		Forms:     forms,
	}

	var buf bytes.Buffer
	if err := renderTemplate(&buf, t, "edit", data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
