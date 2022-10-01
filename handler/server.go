package handler

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
	"text/template"

	"github.com/spudtrooper/goutil/or"
)

//go:generate genopts --function CreateHandler indexTitle:string prefix:string
func CreateHandler(ctx context.Context, hs []Handler, optss ...CreateHandlerOption) *http.ServeMux {
	opts := MakeCreateHandlerOptions(optss...)
	indexTitle := or.String(opts.IndexTitle(), "API")
	prefix := or.String(opts.Prefix(), "api")

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
			evalCtx := &serverEvalContext{
				ctx: ctx,
				w:   w,
				req: req,
			}
			res, err := h.fn(evalCtx)
			if err != nil {
				respondWithError(w, req, err)
				return
			}
			respondWithJSON(req, w, res)
		})
	}

	handleFunc(fmt.Sprintf("/%s/", prefix), func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, genIndex(indexTitle, prefix, routesToHandlers))
	})

	handleFunc(fmt.Sprintf("/%s/_edit", prefix), func(w http.ResponseWriter, req *http.Request) {
		route, ok := getStringURLParamOrDie(w, req, "route")
		if !ok {
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, genEdit(indexTitle, route, routesToHandlers))
	})

	return mux
}

func genEdit(title string, route string, routesToHandlers map[string]*handler) string {
	const t = `
<!DOCTYPE html>
<html>
	<head></head>
	<body>
		<h1>{{.Title}}</h1>
		<h2>{{.Route}}</h2>
		<form action="{{.Route}}" method="GET">
			{{range .Forms}}
				<div>
					{{.Name}}: <input type="text" name="{{.Name}}" size="50"></input>
				</div>
			{{end}}
			<br/>
			<input type="submit"></input>
		</form>
	</body>
</html>
	`
	type form struct {
		Name string
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
				Name: p.Name,
			}
			forms = append(forms, f)
		}
		break
	}
	var data = struct {
		Title string
		Route string
		Forms []form
	}{
		Title: title,
		Route: route,
		Forms: forms,
	}

	var buf bytes.Buffer
	renderTemplate(&buf, t, "index", data)
	return buf.String()
}

func genIndex(title string, prefix string, routesToHandlers map[string]*handler) string {
	const t = `
{{$prefix := .Prefix}}
<!DOCTYPE html>
<html>
	<head></head>
	<body>
		<h1>{{.Title}}</h1>
		<ul>
			{{range .Routes}}
				<li>
					<a href="{{.Route}}">{{.Route}}</a>					
					{{if .HasParams}}
						(<a href="/{{$prefix}}/_edit?route={{.Route}}">edit</a>)
					{{end}}
				</li>
			{{end}}
		</ul>
	</body>
</html>
	`
	type route struct {
		Route     string
		HasParams bool
	}
	var routePaths []string
	for r := range routesToHandlers {
		routePaths = append(routePaths, r)
	}
	sort.Strings(routePaths)
	var routes []route
	for _, r := range routePaths {
		routes = append(routes, route{
			Route:     r,
			HasParams: !routesToHandlers[r].metadata.Empty(),
		})
	}
	var data = struct {
		Title  string
		Prefix string
		Routes []route
	}{
		Title:  title,
		Prefix: prefix,
		Routes: routes,
	}

	var buf bytes.Buffer
	renderTemplate(&buf, t, "index", data)
	return buf.String()
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
