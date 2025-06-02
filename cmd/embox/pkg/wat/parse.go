package wat

import (
	"errors"
	"fmt"
	"io"
	"sdl3/cmd/embox/pkg/wasm"
	"sdl3/cmd/embox/pkg/wat/internal/ast"
	"sdl3/cmd/embox/pkg/wat/internal/matcher"
)

var (
	ErrUnexpectedEOF       = io.ErrUnexpectedEOF
	ErrElementIdOutOfRange = errors.New("element")
)

type ParseOptions struct {
	Strict bool // throw error on unknown attributes, unknown nodes always throws
}

func ParseModule(r io.Reader, opts *ParseOptions) (*wasm.Module, error) {
	if opts == nil {
		opts = &ParseOptions{}
	}

	astRoot, err := ast.Parse(r)
	if err != nil {
		return nil, fmt.Errorf("failed parse wat: %w", err)
	}

	cursor := matcher.NewCursor(astRoot)

	res, ok := module.TryMatch(cursor)
	if !ok {
		panic("not ok")
	}

	panic(res)
}

var (
	module = matcher.OneOf(
		matcher.NamedSubnode("module", moduleInner),
		moduleInner,
	)

	moduleInner = matcher.RepeatAtLeast(0,
		typedef,
	)

	typedef = matcher.NamedSubnode("type", matcher.Sequence(
		matcher.Optional(id),
		functype,
	))

	id = matcher.OneOf(
		matcher.Identifier(),
		matcher.Int(),
	)

	functype = matcher.NamedSubnode("func", matcher.Sequence(
		matcher.Repeat(param),
		matcher.Repeat(result),
	))

	param = matcher.NamedSubnode("param", matcher.OneOf(
		matcher.Sequence(
			id,
			valtype,
		),
		matcher.RepeatAtLeast(1, valtype),
	))

	result = matcher.NamedSubnode("result", matcher.RepeatAtLeast(1, valtype))

	numtype = matcher.OneOf(
		matcher.KeywordExact("i32"),
		matcher.KeywordExact("i64"),
		matcher.KeywordExact("f32"),
		matcher.KeywordExact("f64"),
	).Transform(func(in any) any {
		switch in.(*ast.Node).StrValue {
		case "i32":
			return wasm.Type_i32
		case "i64":
			return wasm.Type_i64
		case "f32":
			return wasm.Type_f32
		}
		return wasm.Type_f64
	})

	vectype = matcher.KeywordExact("v128").Transform(func(in any) any {
		return wasm.Type_v128
	})

	reftype = matcher.OneOf(
		matcher.KeywordExact("funcref"),
		matcher.KeywordExact("externref"),
	).Transform(func(in any) any {
		if in.(*ast.Node).StrValue == "funcref" {
			return wasm.Type_FuncRef
		} else {
			return wasm.Type_ExternRef
		}
	})

	valtype = matcher.OneOf(
		numtype,
		vectype,
		reftype,
	)
)
