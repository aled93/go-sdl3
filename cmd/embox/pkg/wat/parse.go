package wat

import (
	"fmt"
	"io"
	"sdl3/cmd/embox/pkg/wasm"
)

var (
	ErrUnexpectedEOF = io.ErrUnexpectedEOF
)

type ParseOptions struct {
	Strict bool // throw error on unknown attributes, unknown nodes always throws
}

func ParseModule(r io.Reader, opts *ParseOptions) (*wasm.Module, error) {
	if opts == nil {
		opts = &ParseOptions{}
	}

	astRoot, err := parseAst(r)
	if err != nil {
		return nil, fmt.Errorf("failed parse wat: %w", err)
	}

	parser := Parser{
		astRoot: astRoot,
		opts:    *opts,
	}

	return parser.ParseModule()
}

type Parser struct {
	astRoot *AstNode
	opts    ParseOptions
}

func (p *Parser) ParseModule() (*wasm.Module, error) {
	module := &wasm.Module{}

	panic("not implemented")

	return module, nil
}
