// These functions are a drop in replacement for the standard `flag` package that
// allow you to easily bind them all to a `cliAdapter`.
package handler

import (
	"flag"
	"time"
)

type cache struct {
	strFlags     map[string]*string
	boolFlags    map[string]*bool
	intFlags     map[string]*int
	float64Flags map[string]*float64
	durFlags     map[string]*time.Duration
}

func newCache() *cache {
	return &cache{
		strFlags:     make(map[string]*string),
		boolFlags:    make(map[string]*bool),
		intFlags:     make(map[string]*int),
		float64Flags: make(map[string]*float64),
		durFlags:     make(map[string]*time.Duration),
	}
}

var globalFlagCache = newCache()

func String(name, def, desc string) *string {
	res := flag.String(name, def, desc)
	globalFlagCache.strFlags[name] = res
	return res
}

func Bool(name string, def bool, desc string) *bool {
	res := flag.Bool(name, def, desc)
	globalFlagCache.boolFlags[name] = res
	return res
}

func Int(name string, def int, desc string) *int {
	res := flag.Int(name, def, desc)
	globalFlagCache.intFlags[name] = res
	return res
}

func Float64(name string, def float64, desc string) *float64 {
	res := flag.Float64(name, def, desc)
	globalFlagCache.float64Flags[name] = res
	return res
}

func Duration(name string, def time.Duration, desc string) *time.Duration {
	res := flag.Duration(name, def, desc)
	globalFlagCache.durFlags[name] = res
	return res
}
