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

	res, err := module.TryMatch(cursor)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

var (
	module = matcher.Transform(
		matcher.OneOf2(
			matcher.NamedSubnode("module", moduleInner),
			moduleInner,
		),
		func(in tuple.Of3[wasm.Module, wasm.Module, int]) wasm.Module {
			if in.M3 == 1 {
				return in.M1
			}
			return in.M2
		},
	)

	moduleInner = matcher.Transform(
		matcher.RepeatAtLeast(0,
			typedef,
		),
		func(in []wasm.TypeDef) wasm.Module {
			return wasm.Module{
				Typedefs: in,
			}
		},
	)

	typedef = matcher.Transform(
		matcher.NamedSubnode("type", matcher.Sequence2(
			matcher.Optional(id),
			funcsignature,
		)),
		func(in tuple.Of2[wasm.ElementId, wasm.FuncSignature]) wasm.TypeDef {
			return wasm.TypeDef{
				Id:       in.M1,
				FuncType: in.M2,
			}
		},
	)

	id = matcher.Transform(
		matcher.OneOfSame(
			matcher.Identifier(),
			matcher.Int(),
		),
		func(in tuple.Of2[*ast.Node, int]) wasm.ElementId {
			if in.M2 == 0 {
				return wasm.ElementId{
					HasValue: true,
					Name:     in.M1.StrValue,
				}
			}
			return wasm.ElementId{
				HasValue: true,
				Index:    uint32(in.M1.IntValue),
			}
		},
	)

	funcsignature = matcher.Transform(
		matcher.NamedSubnode("func", matcher.Sequence2(
			matcher.Repeat(param),
			matcher.Repeat(result),
		)),
		func(in tuple.Of2[[]wasm.Param, []wasm.Result]) wasm.FuncSignature {
			return wasm.FuncSignature{
				Params:  in.M1,
				Results: in.M2,
			}
		},
	)

	param = matcher.Transform(
		matcher.NamedSubnode("param", matcher.OneOf2(
			matcher.Sequence2(
				id,
				valtype,
			),
			matcher.RepeatAtLeast(1, valtype),
		)),
		func(in tuple.Of3[tuple.Of2[wasm.ElementId, wasm.ValueType], []wasm.ValueType, int]) wasm.Param {
			if in.M3 == 1 {
				return wasm.Param{
					Name: in.M1.M1.Name,
					Type: in.M1.M2,
				}
			}
			panic("TODO multiple param abbreviation")
		},
	)

	result = matcher.Transform(
		matcher.NamedSubnode("result", matcher.RepeatAtLeast(1, valtype)),
		func(in []wasm.ValueType) wasm.Result {
			if len(in) > 1 {
				panic("multiple returns not supported yet")
			}

			return wasm.Result{
				Type: in[0],
			}
		},
	)

	numtype = matcher.Transform(
		matcher.OneOfSame(
			matcher.KeywordExact("i32"),
			matcher.KeywordExact("i64"),
			matcher.KeywordExact("f32"),
			matcher.KeywordExact("f64"),
		),
		func(in tuple.Of2[*ast.Node, int]) wasm.NumberType {
			switch in.M2 {
			case 0:
				return wasm.Type_i32
			case 1:
				return wasm.Type_i64
			case 2:
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
		matcher.OneOfSame(
			matcher.KeywordExact("funcref"),
			matcher.KeywordExact("externref"),
		),
		func(in tuple.Of2[*ast.Node, int]) wasm.RefType {
			if in.M2 == 0 {
				return wasm.Type_FuncRef
			} else {
				return wasm.Type_ExternRef
			}
		},
	)

	valtype = matcher.Transform(
		matcher.OneOf3(
			numtype,
			vectype,
			reftype,
		),
		func(in tuple.Of4[wasm.NumberType, wasm.VectorType, wasm.RefType, int]) wasm.ValueType {
			switch in.M4 {
			case 1:
				return in.M1
			case 2:
				return in.M2
			}
			return in.M3
		},
	)
)
