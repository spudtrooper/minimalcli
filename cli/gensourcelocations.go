package cli

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"

	"github.com/pkg/errors"
	"github.com/spudtrooper/minimalcli/handler"
)

func genSourceLocations(input, sourceLinkURIRoot, output string) error {
	m, err := handler.FindHandlerSourceLocations([]string{input}, sourceLinkURIRoot)
	if err != nil {
		return errors.Wrapf(err, "failed to find handler source locations")
	}
	var locs []handler.HandlerSourceLocation
	for h, loc := range m {
		locs = append(locs, handler.HandlerSourceLocation{
			Handler: h,
			Loc:     loc,
		})
	}
	// Sort by handler name to make this deterministic
	sort.Slice(locs, func(i, j int) bool { return locs[i].Handler < locs[j].Handler })
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
