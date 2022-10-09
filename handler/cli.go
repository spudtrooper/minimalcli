package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/spudtrooper/goutil/json"
	minimalcli "github.com/spudtrooper/minimalcli/app"
)

type cliAdapter struct {
	strFlags     map[string]*string
	boolFlags    map[string]*bool
	intFlags     map[string]*int
	durFlags     map[string]*time.Duration
	float32Flags map[string]*float32
	float64Flags map[string]*float64
}

func NewCLIAdapter() *cliAdapter {
	return &cliAdapter{
		strFlags:     make(map[string]*string),
		boolFlags:    make(map[string]*bool),
		intFlags:     make(map[string]*int),
		durFlags:     make(map[string]*time.Duration),
		float32Flags: make(map[string]*float32),
		float64Flags: make(map[string]*float64),
	}
}

func (c *cliAdapter) BindStringFlag(name string, flag *string)          { c.strFlags[name] = flag }
func (c *cliAdapter) BindBoolFlag(name string, flag *bool)              { c.boolFlags[name] = flag }
func (c *cliAdapter) BindIntFlag(name string, flag *int)                { c.intFlags[name] = flag }
func (c *cliAdapter) BindDurationFlag(name string, flag *time.Duration) { c.durFlags[name] = flag }
func (c *cliAdapter) BindFloat32Flag(name string, flag *float32)        { c.float32Flags[name] = flag }
func (c *cliAdapter) BindFloat64Flag(name string, flag *float64)        { c.float64Flags[name] = flag }

func CreateCLI(a *cliAdapter, hs ...Handler) *minimalcli.App {
	app := minimalcli.Make()
	app.Init()

	for _, h := range hs {
		h := h.(*handler)
		app.Register(h.name, func(ctx context.Context) error {
			evalCtx := &cliEvalContext{
				ctx:          ctx,
				strFlags:     a.strFlags,
				boolFlags:    a.boolFlags,
				intFlags:     a.intFlags,
				durFlags:     a.durFlags,
				float32Flags: a.float32Flags,
				float64Flags: a.float64Flags,
			}
			res, err := h.fn(evalCtx)
			if err != nil {
				return err
			}
			fmt.Printf("%s: %s", h.name, json.MustColorMarshal(res))
			return nil
		})
	}

	return app
}
