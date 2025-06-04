package wat

import (
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
}
