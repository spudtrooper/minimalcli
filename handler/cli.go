package handler

import (
	"context"
	"fmt"
	"time"

	goutiljson "github.com/spudtrooper/goutil/json"
	minimalcli "github.com/spudtrooper/minimalcli/app"
)

type cliAdapter struct {
	strFlags  map[string]*string
	boolFlags map[string]*bool
	intFlags  map[string]*int
	durFlags  map[string]*time.Duration
}

func NewCLIAdapter() *cliAdapter {
	return &cliAdapter{
		strFlags:  make(map[string]*string),
		boolFlags: make(map[string]*bool),
		intFlags:  make(map[string]*int),
		durFlags:  make(map[string]*time.Duration),
	}
}

func (c *cliAdapter) BindStringFlag(name string, flag *string)          { c.strFlags[name] = flag }
func (c *cliAdapter) BindBoolFlag(name string, flag *bool)              { c.boolFlags[name] = flag }
func (c *cliAdapter) BindIntFlag(name string, flag *int)                { c.intFlags[name] = flag }
func (c *cliAdapter) BindDurationFlag(name string, flag *time.Duration) { c.durFlags[name] = flag }

func CreateCLI(a *cliAdapter, hs ...Handler) *minimalcli.App {
	app := minimalcli.Make()
	app.Init()

	for _, h := range hs {
		h := h.(*handler)
		app.Register(h.name, func(ctx context.Context) error {
			evalCtx := &cliEvalContext{
				ctx:       ctx,
				strFlags:  a.strFlags,
				boolFlags: a.boolFlags,
				intFlags:  a.intFlags,
				durFlags:  a.durFlags,
			}
			res, err := h.fn(evalCtx)
			if err != nil {
				return err
			}
			fmt.Printf("%s: %s", h.name, goutiljson.MustColorMarshal(res))
			return nil
		})
	}

	return app
}
