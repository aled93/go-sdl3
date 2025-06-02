package wat

import (
	"errors"
	"fmt"
	"io"
	"sdl3/cmd/embox/pkg/wasm"
	"sdl3/cmd/embox/pkg/wat/internal/ast"
	"sdl3/cmd/embox/pkg/wat/internal/matcher"
	"sdl3/cmd/embox/pkg/wat/internal/tuple"
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
	module = matcher.OneOf2(
		matcher.NamedSubnode("module", moduleInner),
		moduleInner,
	)

	moduleInner = matcher.RepeatAtLeast(0,
		typedef,
	)

	typedef = matcher.NamedSubnode("type", matcher.Sequence2(
		matcher.Optional(id),
		functype,
	))

	id = matcher.OneOf2(
		matcher.Identifier(),
		matcher.Int(),
	)

	functype = matcher.NamedSubnode("func", matcher.Sequence2(
		matcher.Repeat(param),
		matcher.Repeat(result),
	))

	param = matcher.NamedSubnode("param", matcher.OneOf2(
		matcher.Sequence2(
			id,
			valtype,
		),
		matcher.RepeatAtLeast(1, valtype),
	))

	result = matcher.NamedSubnode("result", matcher.RepeatAtLeast(1, valtype))

	numtype = matcher.Transform(
		matcher.OneOf4(
			matcher.KeywordExact("i32"),
			matcher.KeywordExact("i64"),
			matcher.KeywordExact("f32"),
			matcher.KeywordExact("f64"),
		),
		func(in tuple.Of5[*ast.Node, *ast.Node, *ast.Node, *ast.Node, int]) wasm.NumberType {
			switch in.M5 {
			case 1:
				return wasm.Type_i32
			case 2:
				return wasm.Type_i64
			case 3:
				return wasm.Type_f32
			}
			return wasm.Type_f64
		},
	)

	vectype = matcher.Transform(
		matcher.KeywordExact("v128"),
		func(in *ast.Node) wasm.VectorType {
			return wasm.Type_v128
		},
	)

	reftype = matcher.Transform(
		matcher.OneOf2(
			matcher.KeywordExact("funcref"),
			matcher.KeywordExact("externref"),
		),
		func(in tuple.Of3[*ast.Node, *ast.Node, int]) wasm.RefType {
			if in.M3 == 1 {
				return wasm.Type_FuncRef
			} else {
				return wasm.Type_ExternRef
			}
		},
	)

	valtype = matcher.OneOf3(
		numtype,
		vectype,
		reftype,
	)
)
