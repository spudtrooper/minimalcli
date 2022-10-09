package cli

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
	"github.com/spudtrooper/minimalcli/handler"
)

func genSourceLocations(input, sourceLinkURIRoot, output string) error {
	m, err := handler.FindHandlerSourceLocations([]string{input}, sourceLinkURIRoot)
	if err != nil {
		return errors.Wrapf(err, "failed to find handler source locations")
	}
	var locs []handler.HandlerSourceLocation
	for handler, loc := range m {
		locs = append(locs, handler.HandlerSourceLocation{
			Handler: handler,
			Loc:     loc,
		})
	}
	if output == "" {
		log.Printf("handlerSourceLocations: %v\n", locs)
	} else {
		b, err := json.MarshalIndent(locs, "", "  ")
		if err != nil {
			return errors.Wrapf(err, "failed to marshal handlerSourceLocations")
		}
		if err := ioutil.WriteFile(output, []byte(b), 0644); err != nil {
			return errors.Wrapf(err, "failed to write output file")
		}
		log.Printf("wrote %d source locations to %s", len(locs), output)
	}

	return nil
}
