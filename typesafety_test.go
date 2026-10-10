package fulus_test

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"strings"
	"testing"
)

// typeCheck type-checks src as a file of package main that can import this module.
func typeCheck(t *testing.T, src string) error {
	t.Helper()
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "main.go", src, 0)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	conf := types.Config{Importer: importer.ForCompiler(fset, "source", nil)}
	_, err = conf.Check("main", fset, []*ast.File{f}, nil)
	return err
}

func TestTypeSafety(t *testing.T) {
	if testing.Short() {
		t.Skip("type-checks the module from source")
	}
	t.Parallel()

	tests := []struct {
		name    string
		body    string
		wantErr string
	}{
		{
			name:    "different currencies do not add",
			body:    "_, _ = fulus.NewMoney[currency.USD](1).Add(fulus.NewMoney[currency.EUR](1))",
			wantErr: "cannot use",
		},
		{
			name:    "interface type is not a currency",
			body:    "var m fulus.Money[currency.Currency]; _ = m",
			wantErr: "does not satisfy currency.Unit",
		},
		{
			name:    "currency type with state is not a currency",
			body:    "var m fulus.Money[stateful]; _ = m",
			wantErr: "does not satisfy currency.Unit",
		},
		{
			name:    "cross rate middle currencies must agree",
			body:    `_, _ = fulus.Cross(fulus.MustParseRate[currency.USD, currency.JPY]("150"), fulus.MustParseRate[currency.EUR, currency.USD]("1.1"))`,
			wantErr: "does not match inferred type",
		},
		{
			name:    "rate base must be the money currency",
			body:    `_, _ = fulus.Convert(fulus.NewMoney[currency.USD](1), fulus.MustParseRate[currency.EUR, currency.JPY]("160"), fulus.RoundHalfEven)`,
			wantErr: "does not match inferred type",
		},
		{
			name: "empty struct currency is a currency",
			body: "var m fulus.Money[currency.USD]; _ = m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			src := `package main

import (
	"github.com/khatibomar/fulus"
	"github.com/khatibomar/fulus/currency"
)

type stateful struct{ currency.USD; code string }

func main() {
	` + tt.body + `
}
`
			err := typeCheck(t, src)
			switch {
			case tt.wantErr == "" && err != nil:
				t.Fatalf("unexpected error: %v", err)
			case tt.wantErr != "" && (err == nil || !strings.Contains(err.Error(), tt.wantErr)):
				t.Fatalf("error = %v, expected it to contain %q", err, tt.wantErr)
			}
		})
	}
}
