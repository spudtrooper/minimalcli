package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/spudtrooper/goutil/or"
)

var (
	invalidURLParamCharsRE = regexp.MustCompile(`["'<>]`)
)

type errorResponse struct {
	Error string
}

func respondWithError(w http.ResponseWriter, req *http.Request, err error) {
	respondWithJSON(req, w, errorResponse{
		Error: err.Error(),
	})
}

func respondWithErrorString(w http.ResponseWriter, req *http.Request, tmpl string, args ...interface{}) {
	respondWithJSON(req, w, errorResponse{
		Error: fmt.Sprintf(tmpl, args...),
	})
}

func respondWithJSON(req *http.Request, w http.ResponseWriter, obj interface{}) {
	j, err := json.Marshal(obj)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	responseWithJSONBytes(req, w, j)
}

func responseWithJSONBytes(req *http.Request, w http.ResponseWriter, j []byte) {
	fmt.Fprint(w, string(j))
	fmt.Fprint(w, "\n")
	if debug := getBoolURLParam(req, "debug"); debug {
		log.Printf("respondWithJSON: %s", string(j))
	}
}

func sanitizeURLParam(s string) string {
	return invalidURLParamCharsRE.ReplaceAllString(s, "")
}

func getStringURLParam(req *http.Request, key string) string {
	key = sanitizeURLParam(key)
	vals := req.URL.Query()[key]
	if len(vals) > 0 {
		return vals[0]
	}
	return ""
}

func getStringURLParamOrDie(w http.ResponseWriter, req *http.Request, key string) (string, bool) {
	if res := getStringURLParam(req, key); res != "" {
		return res, true
	}
	respondWithErrorString(w, req, fmt.Sprintf("%s required", key))
	return "", false
}

func getStringListURLParam(req *http.Request, key string) []string {
	key = sanitizeURLParam(key)
	p := getStringURLParam(req, key)
	if p == "" {
		return []string{}
	}
	return strings.Split(p, ",")
}

func getIntURLParam(req *http.Request, key string) int {
	key = sanitizeURLParam(key)
	vals := req.URL.Query()[key]
	if len(vals) > 0 {
		v := vals[0]
		if v != "" {
			res, err := strconv.Atoi(v)
			if err != nil {
				log.Printf("getIntURLParam: %v", err)
				return 0
			}
			return res
		}
	}
	return 0
}

func getBoolURLParam(req *http.Request, key string) bool {
	key = sanitizeURLParam(key)
	vals := req.URL.Query()[key]
	if len(vals) > 0 {
		v := vals[0]
		if v == "0" || strings.ToLower(v) == "false" {
			return false
		}
		return true
	}
	return false
}

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
	var routes []string
	for _, h := range hs {
		h := h.(*handler)
		if h.cliOnly {
			continue
		}
		route := fmt.Sprintf("/%s/%s", prefix, strings.ToLower(h.name))
		routes = append(routes, route)
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
	sort.Strings(routes)

	handleFunc(fmt.Sprintf("/%s/", prefix), func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, genIndex(indexTitle, routes))
	})

	return mux
}

func genIndex(title string, routes []string) string {
	const t = `
<!DOCTYPE html>
<html>
	<head></head>
	<body>
		<h1>{{.Title}}</h1>
		<ul>
			{{range .Routes}}<li><a href="{{.}}">{{.}}</a></li>
			{{end}}
		</ul>
	</body>
</html>
	`
	var buf bytes.Buffer
	var data = struct {
		Title  string
		Routes []string
	}{
		Title:  title,
		Routes: routes,
	}
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
