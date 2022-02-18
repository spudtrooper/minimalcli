package app

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/pkg/errors"
	"github.com/thomaso-mirodin/intmath/intgr"
)

var (
	actions = flag.String("actions", "", "comma-delimited list of calls to make")
)

type del func(context.Context, []string) error

type cmd struct {
	name   string
	abbrev string
	fn     del
}

type app struct {
	cmds      []*cmd
	actions   []string
	extraHelp del
	args      []string
}

func Make() *app {
	return &app{}
}

func (a *app) Register(name string, fn del) {
	c := &cmd{
		name: name,
		fn:   fn,
	}
	a.cmds = append(a.cmds, c)
}

func (a *app) init(args []string) error {
	var actionList []string

	// Pull arguments before flags into the acton map
	if len(args) > 1 && !strings.HasPrefix(args[1], "-") {
		lastCommand := 1
		for lastCommand < len(args) {
			if action := strings.TrimSpace(strings.ToLower(args[lastCommand])); action != "" && !strings.HasPrefix(action, "-") {
				actionList = append(actionList, action)
				lastCommand++
			} else {
				break
			}
		}
		var newArgs []string
		for i, arg := range args {
			if i == 0 || i >= lastCommand {
				newArgs = append(newArgs, arg)
			}
		}
		args = newArgs
	}

	// Parse flags after modifying args. Call this directly so we can test with a know set of args.
	flag.CommandLine.Parse(args[1:])

	if *actions != "" {
		for _, c := range strings.Split(*actions, ",") {
			if action := strings.TrimSpace(strings.ToLower(c)); action != "" {
				actionList = append(actionList, action)
			}
		}
	}
	for _, c := range flag.Args() {
		if action := strings.TrimSpace(strings.ToLower(c)); action != "" {
			actionList = append(actionList, action)
		}
	}

	if len(actionList) == 0 {
		return errors.Errorf("you need to specify at least one call")
	}

	a.actions = actionList
	a.args = args

	return nil
}

func (a *app) AddExtraHelp(fn del) {
	a.extraHelp = fn
}

func (a *app) showHelp(ctx context.Context) error {
	repeat := func(n int) string {
		var res string
		for i := 0; i < n; i++ {
			res += "="
		}
		return res
	}
	var namePad int
	{
		maxNameLength := math.MinInt
		for _, c := range a.cmds {
			maxNameLength = intgr.Max(maxNameLength, len(c.name))
		}
		namePad = maxNameLength + 2
	}

	fmt.Println("The following commands are available:")
	fmt.Println()
	fmt.Printf("  %"+fmt.Sprintf("%d", namePad)+"s - %s\n", "Action", "Abbreviation")
	fmt.Printf("  %"+fmt.Sprintf("%d", namePad)+"s - %s\n", repeat(namePad), repeat(len("Abbreviation")))
	for _, c := range a.cmds {
		fmt.Printf("  %"+fmt.Sprintf("%d", namePad)+"s - %s\n", c.name, c.abbrev)
	}

	if a.extraHelp != nil {
		if err := a.extraHelp(ctx, a.args); err != nil {
			return err
		}
	}

	return nil
}

func (a *app) findCmd(s string) *cmd {
	for _, c := range a.cmds {
		if strings.EqualFold(s, c.name) {
			return c
		}
		if strings.EqualFold(s, c.abbrev) {
			return c
		}
	}
	return nil
}

func (a *app) preRun() {
	sort.Slice(a.cmds, func(i, j int) bool {
		return a.cmds[i].name < a.cmds[j].name
	})
	getAbbrev := func(name string) string {
		var buf bytes.Buffer
		for _, s := range name {
			if s := string(s); strings.ToUpper(s) == s {
				buf.WriteString(strings.ToLower(s))
			}
		}
		return buf.String()
	}
	isUnique := func(s string) bool {
		for _, c := range a.cmds {
			if c.abbrev == s {
				return false
			}
		}
		return true
	}
	for _, c := range a.cmds {
		abbrev := getAbbrev(c.name)
		if !isUnique(abbrev) {
			for i := 0; i <= len(c.name); i++ {
				sub := strings.ToLower(string(c.name[0:i]))
				if isUnique(sub) {
					abbrev = sub
					break
				}
			}
		}
		c.abbrev = abbrev
	}
}

func (a *app) Run(ctx context.Context, args []string) error {
	a.init(args)
	a.Register("Help", func(ctx context.Context, args []string) error {
		if err := a.showHelp(ctx); err != nil {
			return err
		}
		return nil
	})
	a.preRun()
	if len(a.actions) == 0 {
		a.showHelp(ctx)
		return nil
	}
	for _, s := range a.actions {
		c := a.findCmd(s)
		if c == nil {
			return errors.Errorf("no action for %q", s)
		}
		if err := c.fn(ctx, a.args); err != nil {
			return errors.Errorf("running %q: %v", c.name, err)
		}
	}
	return nil
}
