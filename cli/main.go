package cli

import (
	"context"
	"log"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/flags"
	"github.com/spudtrooper/minimalcli/app"
)

var (
	input             = flags.String("input", "input go file")
	output            = flags.String("output", "output JSON file")
	sourceLinkURIRoot = flags.String("uri_root", "source link uri root")
	iniiDir           = flags.String("init_dir", "directory to which we dump the output of initializing a directory")
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

	app.Register("Init", func(context.Context) error {
		dir := "."
		if *iniiDir != "" {
			dir = *iniiDir
		}
		abs, err := filepath.Abs(dir)
		if err != nil {
			return errors.Errorf("getting absolute path of %s: %v", dir, err)
		}
		return initDirectory(abs)
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
