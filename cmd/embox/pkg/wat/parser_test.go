package wat

import (
	"sdl3/cmd/embox/pkg/wasm"
	"strings"
	"testing"
)

func TestParseEmptyModule(t *testing.T) {
	src := "(module )"

	_, err := ParseModule(strings.NewReader(src), &ParseOptions{
		Strict: true,
	})
	if err != nil {
		panic(err)
	}
}

func TestParseModuleWithForeignNodes(t *testing.T) {
	src := "(module foreign)"

	_, err := ParseModule(strings.NewReader(src), &ParseOptions{
		Strict: true,
	})
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestParseSimpleModule(t *testing.T) {
	src := `(module
  (type $dotest (func (param $a i32) (param $b i64) (result f64)))
)`
	// FIXME: attribute matchers doesn't advance cursor
	mod, err := ParseModule(strings.NewReader(src), &ParseOptions{
		Strict: true,
	})
	if err != nil {
		panic(err)
	}

	if len(mod.Typedefs) != 1 {
		t.Fatalf("types not parsed")
	}

	if mod.Typedefs[0].Id.Name != "dotest" {
		t.Fatalf(`type name expected "dotest", got "%s"`, mod.Typedefs[0].Id.Name)
	}

	var zero wasm.ElementId
	if mod.Typedefs[0].FuncType.Id != zero {
		t.Errorf(`type's func id expected be empty, got %v`, mod.Typedefs[0].FuncType.Id)
	}

	if len(mod.Typedefs[0].FuncType.Params) != 2 {
		t.Errorf(`type's func params expected length of 2, got %d`, len(mod.Typedefs[0].FuncType.Params))
	}
}
