package cli

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/fatih/color"
	"github.com/spudtrooper/goutil/io"
	"github.com/spudtrooper/goutil/task"
)

const username = "spudtrooper" // TODO

func initDirectory(outDir string) error {
	log.Printf("initializing directory %s", outDir)

	var rootDir, scriptsDir, cliDir, frontendDir, apiDir, handlersDir string

	absDot, err := filepath.Abs(".")
	if err != nil {
		return err
	}
	pkg := filepath.Base(absDot)

	tb := task.MakeBuilder(task.Color(color.New(color.FgYellow)))

	tb.Add("initializing directories", func() error {
		r, err := io.MkdirAll(outDir)
		if err != nil {
			return err
		}
		rootDir = r

		s, err := io.MkdirAll(rootDir, "scripts")
		if err != nil {
			return err
		}
		scriptsDir = s

		c, err := io.MkdirAll(rootDir, "cli")
		if err != nil {
			return err
		}
		cliDir = c

		f, err := io.MkdirAll(rootDir, "frontend")
		if err != nil {
			return err
		}
		frontendDir = f

		a, err := io.MkdirAll(rootDir, "api")
		if err != nil {
			return err
		}
		apiDir = a

		h, err := io.MkdirAll(rootDir, "handlers")
		if err != nil {
			return err
		}
		handlersDir = h

		return nil
	})

	tb.Add("creating files", func() error {
		if _, err := writeFile(`
# {{.Pkg}}

A TODO API for go.

## Usage

### Running local front end

`+"```"+`bash
./scripts/frontend_local.sh
`+"```"+`
	`, struct {
			Pkg string
		}{
			Pkg: pkg,
		}, rootDir, "README.md"); err != nil {
			return err
		}

		if _, err := writeFile(`
		package main

		import (
			"context"
		
			"github.com/spudtrooper/goutil/check"
			"github.com/{{.Username}}/{{.Pkg}}/cli"
		)
		
		func main() {
			check.Err(cli.Main(context.Background()))
		}
			`, struct {
			Pkg      string
			Username string
		}{
			Pkg:      pkg,
			Username: username,
		}, rootDir, "main.go"); err != nil {
			return err
		}

		if _, err := writeFile(`
		package main

		import (
			"context"
		
			"github.com/spudtrooper/goutil/check"
			"github.com/{{.Username}}/{{.Pkg}}/cli"
		)
		
		func main() {
			check.Err(cli.Main(context.Background()))
		}
			`, struct {
			Pkg      string
			Username string
		}{
			Pkg:      pkg,
			Username: username,
		}, cliDir, "main.go"); err != nil {
			return err
		}

		if _, err := writeFile(`
		package main

		import (
			"context"
			"flag"
			"log"
			"os"
			"strconv"
		
			"github.com/spudtrooper/goutil/check"
			"github.com/{{.Username}}/{{.Pkg}}/api"
			"github.com/{{.Username}}/{{.Pkg}}/frontend"
		)
		
		var (
			portForTesting = flag.Int("port_for_testing", 0, "port to listen on")
			host           = flag.String("host", "localhost", "host name")
		)
		
		func main() {
			flag.Parse()
			var port int
			if *portForTesting != 0 {
				port = *portForTesting
			} else {
				p, err := strconv.Atoi(os.Getenv("PORT"))
				if err != nil {
					log.Fatalf("invalid port: %v", err)
				}
				port = p
			}
			if port == 0 {
				log.Fatalf("port is required")
			}
		
			ctx := context.Background()
			client := api.NewClient()
			check.Err(frontend.ListenAndServe(ctx, client, port, *host))
		}
			`, struct {
			Pkg      string
			Username string
		}{
			Pkg:      pkg,
			Username: username,
		}, rootDir, "frontend_main.go"); err != nil {
			return err
		}

		if _, err := writeFile(`
		package cli

		import (
			"context"
		
			"github.com/spudtrooper/minimalcli/handler"
			flag "github.com/spudtrooper/minimalcli/handler"
			"github.com/{{.Username}}/{{.Pkg}}/api"
			"github.com/{{.Username}}/{{.Pkg}}/handlers"
		)
		
		func Main(ctx context.Context) error {
			// TODO: Add flags here.
			if false {
				flag.String("todo", "", "todo")
			}
		
			client, err := api.NewClientFromFlags()
			if err != nil {
				return err
			}
		
			return handler.RunCLI(ctx, handlers.CreateHandlers(client)...)
		}
			`, struct {
			Pkg      string
			Username string
		}{
			Pkg:      pkg,
			Username: username,
		}, cliDir, "main.go"); err != nil {
			return err
		}

		if _, err := writeFile(`
		package frontend

		import (
			"context"
			"fmt"
			"log"
			"net/http"
		
			"github.com/spudtrooper/minimalcli/handler"
			"github.com/{{.Username}}/{{.Pkg}}/api"
			"github.com/{{.Username}}/{{.Pkg}}/handlers"
		)
		
		func ListenAndServe(ctx context.Context, client *api.Client, port int, host string) error {
			mux := http.NewServeMux()
			handler.Init(mux)
			if err := handler.AddHandlers(ctx, mux, handlers.CreateHandlers(client),
				//TODO: Add handlers here.
			); err != nil {
				return err
			}
			mux.Handle("/", http.RedirectHandler("/api", http.StatusSeeOther))
		
			log.Printf("listening on http://%s:%d", host, port)
		
			if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
				return err
			}
		
			return nil
		}
			`, struct {
			Pkg      string
			Username string
		}{
			Pkg:      pkg,
			Username: username,
		}, frontendDir, "server.go"); err != nil {
			return err
		}

		if _, err := writeFile(`
		package api

		import (
			"encoding/json"
			"flag"
			"io/ioutil"
			"log"
		
			"github.com/spudtrooper/goutil/io"
		)
		
		var (
			userCreds   = flag.String("{{.Pkg}}_user_creds", ".user_creds.json", "file with user credentials")
			noUserCreds = flag.Bool("{{.Pkg}}_no_user_creds", false, "Don't use user creds event if it exists")
		)
		
		//go:generate genopts --function Base
		
		// Client is a client for spotifydown.com
		type Client struct {
		}
		
		// NewClientFromFlags creates a new client from command line flags
		func NewClientFromFlags() (*Client, error) {
			client := NewClient()
			return client, nil
		}
		
		// NewClient creates a new client directly from the API Key
		func NewClient() *Client {
			return &Client{}
		}
		
		// NewClientFromFile creates a new client from a JSON file like user_creds-example.json
		func NewClientFromFile(credsFile string) (*Client, error) {
			return &Client{}, nil
		}
		
		type Creds struct {
			// TODO: Add creds here.
		}
		
		func ReadCredsFromFlags() (Creds, error) {
			return readCreds(*userCreds)
		}
		
		func WriteCredsFromFlags(creds Creds) error {
			b, err := json.Marshal(&creds)
			if err != nil {
				return err
			}
			if err := ioutil.WriteFile(*userCreds, b, 0755); err != nil {
				return err
			}
			log.Printf("wrote to %s", *userCreds)
			return nil
		}
		
		func readCreds(credsFile string) (creds Creds, ret error) {
			if !io.FileExists(credsFile) {
				return
			}
			credsBytes, err := ioutil.ReadFile(credsFile)
			if err != nil {
				ret = err
				return
			}
			if err := json.Unmarshal(credsBytes, &creds); err != nil {
				ret = err
				return
			}
			return
		}
			`, struct {
			Pkg      string
			Username string
		}{
			Pkg:      pkg,
			Username: username,
		}, apiDir, "client.go"); err != nil {
			return err
		}

		if _, err := writeFile(`
		package handlers

		import (
			_ "embed"
		
			"github.com/spudtrooper/minimalcli/handler"
			"github.com/{{.Username}}/{{.Pkg}}/api"
		)
		
		//go:generate minimalcli gsl --input handlers.go --uri_root "github.com/{{.Username}}/{{.Pkg}}/blob/main/handlers" --output handlers.go.json
		//go:embed handlers.go.json
		var SourceLocations []byte
		
		func CreateHandlers(client *api.Client) []handler.Handler {
			b := handler.NewHandlerBuilder()
		
			/*
			TODO: Example
			b.NewHandler("Metadata",
				func(ctx context.Context, ip any) (any, error) {
					p := ip.(api.MetadataParams)
					return client.Metadata(p.Options()...)
				},
				api.MetadataParams{},
				handler.NewHandlerExtraRequiredFields([]string{"track"}),
			)
			*/
				
			return b.Build()
		}		
			`, struct {
			Pkg      string
			Username string
		}{
			Pkg:      pkg,
			Username: username,
		}, handlersDir, "handlers.go"); err != nil {
			return err
		}

		if _, err := writeFile(`
	#!/bin/sh
	#
	# Runs the frontend server locally.
	#
	set -e
	
	SCRIPTS="$(dirname "$0")"
	
	go generate ./...
	
	$SCRIPTS/just_frontend_local.sh "$@"
				`, struct {
		}{}, scriptsDir, "frontend_local.sh"); err != nil {
			return err
		}

		if _, err := writeFile(`
	#!/bin/sh
	#
	# Runs the frontend server locally.
	#
	set -e
	
	go mod tidy
	go run frontend_main.go --port_for_testing 8080 --host localhost "$@"
				`, struct {
		}{}, scriptsDir, "just_frontend_local.sh"); err != nil {
			return err
		}

		if _, err := writeFile(`
	#!/bin/sh
	#
	# Runs main
	#
	set -e
	
	go mod tidy
	go run main.go "$@"
				`, struct {
		}{}, scriptsDir, "just_run.sh"); err != nil {
			return err
		}

		if _, err := writeFile(`
	#!/bin/sh
	#
	# Runs main
	#
	set -e
	
	SCRIPTS="$(dirname "$0")"
	
	go generate ./...
	
	$SCRIPTS/just_run.sh "$@"
				`, struct {
		}{}, scriptsDir, "run.sh"); err != nil {
			return err
		}

		if err := run(scriptsDir, "chmod", "+x", "frontend_local.sh"); err != nil {
			return err
		}
		if err := run(scriptsDir, "chmod", "+x", "just_frontend_local.sh"); err != nil {
			return err
		}
		if err := run(scriptsDir, "chmod", "+x", "just_run.sh"); err != nil {
			return err
		}
		if err := run(scriptsDir, "chmod", "+x", "run.sh"); err != nil {
			return err
		}

		return nil
	})

	tb.Add("formating files", func() error {
		if err := run(rootDir, "gofmt", "-s", "-w", "."); err != nil {
			return err
		}
		return nil
	})

	tb.Add("updating dependencies", func() error {
		if err := run(rootDir, "go", "mod", "init", fmt.Sprintf("github.com/%s/%s", username, pkg)); err != nil {
			return err
		}
		if err := run(rootDir, "go", "mod", "tidy"); err != nil {
			return err
		}
		return nil
	})

	tb.Add("generating handlers JSON", func() error {
		if err := run(rootDir, "go", "generate", "./..."); err != nil {
			return err
		}
		return nil
	})

	tb.Add("testing", func() error {
		if err := run(rootDir, path.Join(scriptsDir, "run.sh")); err != nil {
			return err
		}
		return nil
	})

	return tb.Build().Go()
}

func writeFile(t string, data interface{}, dir string, outFileName string) (string, error) {
	b, err := renderTemplate(t, outFileName, data)
	if err != nil {
		return "", err
	}
	outFile := path.Join(dir, outFileName)
	if err := ioutil.WriteFile(outFile, b, 7055); err != nil {
		return "", err
	}
	subPrintf("wrote to %s", outFile)
	return outFile, nil
}

func run(dir, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func renderTemplate(t string, name string, data interface{}) ([]byte, error) {
	tmpl, err := template.New(name).Parse(strings.TrimSpace(t))
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func subPrintf(tmpl string, args ...interface{}) {
	log.Println(color.GreenString(fmt.Sprintf("  "+tmpl, args...)))
}
