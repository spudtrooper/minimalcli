package cli

import (
	"context"
	"log"

	"github.com/spudtrooper/goutil/flags"
	"github.com/spudtrooper/minimalcli/app"
)

var (
	input             = flags.String("input", "input go file")
	output            = flags.String("output", "output JSON file")
	sourceLinkURIRoot = flags.String("uri_root", "source link uri root")
)

func Main(ctx context.Context) error {
	app := app.Make()
	app.Init()

	app.Register("GenSourceLocations", func(context.Context) error {
		requireStringFlag(input, "input")
		requireStringFlag(sourceLinkURIRoot, "uri_root")
		if err := genSourceLocations(*input, *sourceLinkURIRoot, *output); err != nil {
			return err
		}
		return nil
	})

	if err := app.Run(ctx); err != nil {
		return err
	}

	return nil
}

func requireStringFlag(flag *string, name string) {
	if *flag == "" {
		log.Fatalf("--%s required", name)
	}
}
