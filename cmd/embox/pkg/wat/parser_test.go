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
  (import "js" "print" (func (param i32)))
)`

	_, err := ParseModule(strings.NewReader(src), &ParseOptions{
		Strict: true,
	})
	if err != nil {
		panic(err)
	}
}
