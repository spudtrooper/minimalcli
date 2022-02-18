package app

import (
	"context"
	"strings"
	"testing"

	"github.com/spudtrooper/goutil/or"
)

func TestApp(t *testing.T) {
	var tests = []struct {
		name                                                         string
		args                                                         []string
		wanntFooCalled, wantBarCalled, wantBazCalled, wantHelpCalled bool
		wantErr                                                      bool
	}{
		{
			args:           []string{"app-name", "Foo", "Bar"},
			wanntFooCalled: true,
			wantBarCalled:  true,
		},
		{
			name:           "lowercase",
			args:           []string{"app-name", "foo", "bar"},
			wanntFooCalled: true,
			wantBarCalled:  true,
		},
		{
			name:           "abbrev",
			args:           []string{"app-name", "f", "b"},
			wanntFooCalled: true,
			wantBarCalled:  true,
		},
		{
			name:           "abbrev 2",
			args:           []string{"app-name", "f", "b", "ba"},
			wanntFooCalled: true,
			wantBarCalled:  true,
			wantBazCalled:  true,
		},
		{
			name:           "uppercase",
			args:           []string{"app-name", "FOO", "BAR"},
			wanntFooCalled: true,
			wantBarCalled:  true,
		},
		{
			name:           "help no args",
			args:           []string{"app-name"},
			wantHelpCalled: true,
		},
		{
			name:           "help",
			args:           []string{"app-name", "help"},
			wantHelpCalled: true,
		},
		{
			name:           "help abbrev",
			args:           []string{"app-name", "h"},
			wantHelpCalled: true,
		},
	}
	for _, test := range tests {
		name := or.String(test.name, strings.Join(test.args, " "))
		t.Run(name, func(t *testing.T) {
			app := Make()

			fooCalled := false
			barCalled := false
			bazCalled := false
			helpCalled := false

			app.AddExtraHelp(func(context.Context) error {
				helpCalled = true
				return nil
			})

			app.Register("Foo", func(ctx context.Context) error {
				fooCalled = true
				return nil
			})
			app.Register("Bar", func(ctx context.Context) error {
				barCalled = true
				return nil
			})
			app.Register("Baz", func(ctx context.Context) error {
				bazCalled = true
				return nil
			})

			err := app.Run(context.Background(), test.args)
			if test.wantErr && err == nil {
				t.Fatalf("want error and got none")
			}
			if !test.wantErr && err != nil {
				t.Fatalf("Run: %v", err)
			}

			if want, got := fooCalled, test.wanntFooCalled; want != got {
				t.Errorf("fooCalled: got %v, wanted %v", got, want)
			}
			if want, got := barCalled, test.wantBarCalled; want != got {
				t.Errorf("barCalled: got %v, wanted %v", got, want)
			}
			if want, got := bazCalled, test.wantBazCalled; want != got {
				t.Errorf("bazCalled: got %v, wanted %v", got, want)
			}
			if want, got := helpCalled, test.wantHelpCalled; want != got {
				t.Errorf("helpCalled: got %v, wanted %v", got, want)
			}
		})
	}
}
