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

	_, err := ParseModule(strings.NewReader(src), &ParseOptions{
		Strict: true,
	})
	if err != nil {
		panic(err)
	}
}
