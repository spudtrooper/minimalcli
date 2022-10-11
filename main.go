package main

import (
	"context"

	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/minimalcli/cli"
	"github.com/spudtrooper/minimalcli/gitversion"
)

func main() {
	if gitversion.CheckVersionFlag() {
		return nil
	}
	check.Err(cli.Main(context.Background()))
}
