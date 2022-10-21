package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/pkg/errors"
	goutiltime "github.com/spudtrooper/goutil/time"
)

type EvalContext interface {
	Context() context.Context
	String(name string) string
	MustString(name string) (string, bool)
	Bool(name string) bool
	Int(name string) int
	MustInt(name string) (int, bool)
	Float32(name string) float32
	Float64(name string) float64
	Duration(name string) time.Duration
	Time(name string) (time.Time, error)
}

type cliEvalContext struct {
	ctx          context.Context
	strFlags     map[string]*string
	boolFlags    map[string]*bool
	intFlags     map[string]*int
	float32Flags map[string]*float32
	float64Flags map[string]*float64
	durFlags     map[string]*time.Duration
	timeFlags    map[string]*time.Time
}

func (c *cliEvalContext) Context() context.Context { return c.ctx }

func (c *cliEvalContext) String(name string) string {
	flag, ok := c.strFlags[name]
	if !ok {
		return ""
	}
	return *flag
}

func (c *cliEvalContext) MustString(name string) (string, bool) {
	val := c.String(name)
	if val == "" {
		log.Fatalf("--%s required", name)
	}
	return val, true
}

func (c *cliEvalContext) Bool(name string) bool {
	flag, ok := c.boolFlags[name]
	if !ok {
		return false
	}
	return *flag
}

func (c *cliEvalContext) Int(name string) int {
	flag, ok := c.intFlags[name]
	if !ok {
		return 0
	}
	return *flag
}

func (c *cliEvalContext) MustInt(name string) (int, bool) {
	val := c.Int(name)
	if val == 0 {
		log.Fatalf("--%s required", name)
	}
	return val, true
}

func (c *cliEvalContext) Duration(name string) time.Duration {
	flag, ok := c.durFlags[name]
	if !ok {
		return 0
	}
	return *flag
}

func (c *cliEvalContext) Time(name string) (time.Time, error) {
	if flag, ok := c.timeFlags[name]; ok {
		return *flag, nil
	}
	flag, ok := c.strFlags[name]
	if !ok {
		var zero time.Time
		return zero, nil
	}
	t, err := goutiltime.Parse(*flag)
	if err != nil {
		return time.Time{}, errors.Errorf("failed to parse time: %v", err)
	}
	return t, nil
}

func (c *cliEvalContext) Float32(name string) float32 {
	flag, ok := c.float32Flags[name]
	if !ok {
		return 0
	}
	return *flag
}

func (c *cliEvalContext) Float64(name string) float64 {
	flag, ok := c.float64Flags[name]
	if !ok {
		return 0
	}
	return *flag
}

type serverEvalContext struct {
	ctx context.Context
	w   http.ResponseWriter
	req *http.Request
}

func (c *serverEvalContext) Context() context.Context    { return c.ctx }
func (c *serverEvalContext) String(name string) string   { return getStringURLParam(c.req, name) }
func (c *serverEvalContext) Bool(name string) bool       { return getBoolURLParam(c.req, name) }
func (c *serverEvalContext) Int(name string) int         { return getIntURLParam(c.req, name) }
func (c *serverEvalContext) Float32(name string) float32 { return getFloat32URLParam(c.req, name) }
func (c *serverEvalContext) Float64(name string) float64 { return getFloat64URLParam(c.req, name) }

func (c *serverEvalContext) MustString(name string) (string, bool) {
	return getStringURLParamOrDie(c.w, c.req, name)
}

func (c *serverEvalContext) MustInt(name string) (int, bool) {
	return getIntURLParamOrDie(c.w, c.req, name)
}

func (c *serverEvalContext) Duration(name string) time.Duration {
	return time.Duration(getIntURLParam(c.req, name))
}

func (c *serverEvalContext) Time(name string) (time.Time, error) {
	s := getStringURLParam(c.req, name)
	if s == "" {
		var t time.Time
		return t, nil
	}
	return goutiltime.Parse(s)
}
