package handler

import (
	"context"
	"log"
	"net/http"
	"time"
)

type EvalContext interface {
	Context() context.Context
	String(name string) string
	MustString(name string) (string, bool)
	Bool(name string) bool
	Int(name string) int
	Float32(name string) float32
	Duration(name string) time.Duration
	Time(name string) (time.Time, error)
}

type cliEvalContext struct {
	ctx          context.Context
	strFlags     map[string]*string
	boolFlags    map[string]*bool
	intFlags     map[string]*int
	float32Flags map[string]*float32
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

func (c *cliEvalContext) Duration(name string) time.Duration {
	flag, ok := c.durFlags[name]
	if !ok {
		return 0
	}
	return *flag
}

func (c *cliEvalContext) Time(name string) (time.Time, error) {
	flag, ok := c.timeFlags[name]
	if !ok {
		var zero time.Time
		return zero, nil
	}
	return *flag, nil
}

func (c *cliEvalContext) Float32(name string) float32 {
	flag, ok := c.float32Flags[name]
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

func (c *serverEvalContext) MustString(name string) (string, bool) {
	return getStringURLParamOrDie(c.w, c.req, name)
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
	// TODO: This isn't going to work, but keeping for now to maintain the interface
	res, err := time.Parse("2006-01-02 15:04", s)
	return res, err
}
