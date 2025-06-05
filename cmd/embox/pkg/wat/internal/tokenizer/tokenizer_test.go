package tokenizer

import (
	_ "embed"
	"io"
	"strings"
	"testing"
)

func TestParseSimpleWat(t *testing.T) {
	src := `(module
 (type $0 (func (param i32) (result i32)))
 (import "env" "emscripten_cancel_main_loop" (func $fimport$0))
 (global $global$0 (mut i32) (i32.const 65536))
 (table $0 1531 funcref)
 (func $0
  (call $3016)
  (call $2621)
  (call $3394)
  (call $2860)
  (call $2862)
 )
)`

	tokenizer := NewTokenizer(strings.NewReader(src))
	for {
		tok, err := tokenizer.NextToken()

		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		t.Log(tok.String())
	}

	t.Log("Done")
}

var (
	//go:embed testdata/sdl3.wat
	sdl3Wat []byte
)

func TestParseSDL3Wat(t *testing.T) {
	tokenizer := NewTokenizer(strings.NewReader(string(sdl3Wat)))
	for {
		_, err := tokenizer.NextToken()

		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
	}
}

func TestParseNumbers(t *testing.T) {
	src := `42 3.14 14.1e-5 -0xDEADBEEF 0xDEAD.BEEF 0xDEAD.BEEFp-4`

	tokenizer := NewTokenizer(strings.NewReader(src))
	for {
		tok, err := tokenizer.NextToken()

		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		t.Log(tok.String())
	}

	t.Log("Done")
}
