package wat

import (
	"errors"
	"fmt"
	"io"
	"os"
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

	print("Restored AST: ")
	_, err = astRoot.WriteTo(os.Stdout)
	if err != nil {
		panic(err)
	}
	println()

	cursor := matcher.NewCursor(astRoot)

	res, err := module.TryMatch(cursor)
	if err != nil {
		return nil, err
	}

	if !cursor.EndOfSubnode() {
		return nil, fmt.Errorf("unexpected node %s", cursor.Node())
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
		matcher.Repeat(
			matcher.OneOf3(
				typedef,
				table,
				elem,
			),
		),
		func(in []tuple.Of4[
			wasm.TypeDef,
			wasm.Table,
			[]wasm.TableElem,
			int,
		]) wasm.Module {
			module := wasm.Module{}

			for _, d := range in {
				switch d.M4 {
				case 1:
					module.Typedefs = append(module.Typedefs, d.M1)

				case 2:
					module.Tables = append(module.Tables, d.M2)

				case 3:
					module.TableElems = append(module.TableElems, d.M3...)
				}
			}

			panic("unreachable")
		},
	)

	typedef = matcher.Transform(
		matcher.NamedSubnode("type", matcher.Sequence2(
			matcher.Optional(id),
			funcsignature,
		)),
		func(in tuple.Of2[tuple.Of2[wasm.ElementId, bool], wasm.FuncSignature]) wasm.TypeDef {
			return wasm.TypeDef{
				Id:       in.M1.M1,
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
					Name: in.M1.StrValue,
				}
			}
			return wasm.ElementId{
				Index: uint32(in.M1.IntValue),
			}
		},
	)

	funcsignature = matcher.Transform(
		matcher.NamedSubnode("func", matcher.Sequence2(
			matcher.Repeat(param),
			matcher.Repeat(result),
		)),
		func(in tuple.Of2[[][]wasm.Param, []wasm.Result]) wasm.FuncSignature {
			allParamsLen := 0
			for _, a := range in.M1 {
				allParamsLen += len(a)
			}

			allParams := make([]wasm.Param, 0, allParamsLen)
			for _, p := range in.M1 {
				allParams = append(allParams, p...)
			}

			return wasm.FuncSignature{
				Params:  allParams,
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
		func(in tuple.Of3[tuple.Of2[wasm.ElementId, wasm.ValueType], []wasm.ValueType, int]) []wasm.Param {
			if in.M3 == 1 {
				return []wasm.Param{
					{
						Name: in.M1.M1.Name,
						Type: in.M1.M2,
					},
				}
			}

			res := make([]wasm.Param, len(in.M2))
			for i, v := range in.M2 {
				res[i] = wasm.Param{
					Type: v,
				}
			}
			return res
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

	limits = matcher.Transform(
		matcher.Sequence2(
			matcher.Int(),
			matcher.Optional(matcher.Int()),
		),
		func(in tuple.Of2[*ast.Node, tuple.Of2[*ast.Node, bool]]) (res wasm.Limits) {
			res.Min = uint32(in.M1.IntValue)
			res.HaveMax = in.M2.M2
			if res.HaveMax {
				res.Max = uint32(in.M2.M1.IntValue)
			}
			return
		},
	)

	tableType = matcher.Transform(
		matcher.Sequence2(
			limits,
			reftype,
		),
		func(in tuple.Of2[wasm.Limits, wasm.RefType]) wasm.TableType {
			return wasm.TableType{
				Limits:  in.M1,
				Element: in.M2,
			}
		},
	)

	table = matcher.Transform(
		matcher.NamedSubnode("table",
			matcher.Sequence2(
				matcher.Optional(id),
				tableType,
			),
		),
		func(in tuple.Of2[tuple.Of2[wasm.ElementId, bool], wasm.TableType]) wasm.Table {
			return wasm.Table{
				Id:        in.M1.M1,
				TableType: in.M2,
			}
		},
	)

	elem = matcher.Transform(
		matcher.NamedSubnode("elem",
			matcher.OneOf3(
				matcher.Sequence2(
					matcher.Optional(id),
					elemList,
				),
				matcher.Sequence4(
					matcher.Optional(id),
					tableUse,
					matcher.NamedSubnode("offset",
						expr,
					),
					elemList,
				),
				matcher.Sequence3(
					matcher.Optional(id),
					matcher.KeywordExact("declare"),
					elemList,
				),
			),
		),
		func(in tuple.Of4[
			tuple.Of2[tuple.Of2[wasm.ElementId, bool], []wasm.TableElem],
			tuple.Of4[tuple.Of2[wasm.ElementId, bool], wasm.ElementId, wasm.Expression, []wasm.TableElem],
			tuple.Of3[tuple.Of2[wasm.ElementId, bool], *ast.Node, []wasm.TableElem],
			int,
		]) []wasm.TableElem {
			switch in.M4 {
			case 1:
				for i := range in.M1.M2 {
					in.M1.M2[i].Id = in.M1.M1.M1
				}

				return in.M1.M2

			case 2:
				for i := range in.M2.M4 {
					in.M2.M4[i].Id = in.M2.M1.M1
					in.M2.M4[i].TableId = in.M2.M2
					in.M2.M4[i].Offset = in.M2.M3
				}

				return in.M2.M4
			}

			for i := range in.M3.M3 {
				in.M3.M3[i].Id = in.M3.M1.M1
				in.M3.M3[i].IsDeclarative = true
			}

			return in.M3.M3
		},
	)

	elemList = matcher.Transform(
		matcher.Sequence2(
			reftype,
			matcher.NamedSubnode("item", matcher.Repeat(expr)),
		),
		func(in tuple.Of2[wasm.RefType, []wasm.Expression]) []wasm.TableElem {
			return []wasm.TableElem{
				{
					RefType: in.M1,
					Items:   in.M2,
				},
			}
		},
	)

	tableUse = matcher.NamedSubnode("table", id)

	expr = matcher.Transform(
		matcher.OneOfSame(
			matcher.KeywordExact("i32.const"),
			matcher.KeywordExact("i64.const"),
			matcher.KeywordExact("f32.const"),
			matcher.KeywordExact("f64.const"),
		),
		func(in tuple.Of2[*ast.Node, int]) wasm.Expression {
			return wasm.Expression{} // TODO
		},
	)
)
