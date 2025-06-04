package ast

import (
	_ "embed"
	"fmt"
	"sdl3/cmd/embox/pkg/wat/internal/tokenizer"
	"strings"
	"testing"
)

var (
	//go:embed testdata/sdl3.wat
	sdl3Wat []byte
)

func TestAstParseSimpleWat(t *testing.T) {
	src := `(module (import "js" "print" (func (param i32) (result i32) ) ) )`

	astRoot, err := Parse(strings.NewReader(src))
	if err != nil {
		if ut, ok := err.(*UnexpectedToken); ok {
			t.Fatalf("Parse error: %s\n| %s\n| %s", err.Error(), src, strings.Repeat(" ", ut.GotToken.Pos)+"^")
		} else {
			t.Fatal(err)
		}
	}

	expectedAst := &Node{
		Kind: NodeKind_SubNode,
		Name: "module",
		FirstChild: &Node{
			Kind: NodeKind_SubNode,
			Name: "import",
			FirstChild: astNodeChain(
				&Node{
					Kind:     NodeKind_AttrString,
					StrValue: "js",
				},
				&Node{
					Kind:     NodeKind_AttrString,
					StrValue: "print",
				},
				&Node{
					Kind: NodeKind_SubNode,
					Name: "func",
					FirstChild: astNodeChain(
						&Node{
							Kind: NodeKind_SubNode,
							Name: "param",
							FirstChild: &Node{
								Kind:     NodeKind_AttrKeyword,
								StrValue: "i32",
							},
						},
						&Node{
							Kind: NodeKind_SubNode,
							Name: "result",
							FirstChild: &Node{
								Kind:     NodeKind_AttrKeyword,
								StrValue: "i32",
							},
						},
					),
				},
			),
		},
	}

	if err, tok := compareNodes(expectedAst, astRoot); err != nil {
		code, arrow := getTokenContext(src, tok)
		t.Error("got wrong ast:", err, "\n| "+code+"\n| "+arrow)
	}
}

func TestAstParseMalformedInput(t *testing.T) {
	src := `(module`

	_, err := Parse(strings.NewReader(src))
	if err == nil {
		t.Fatal("expected parse error, got nil error")
	}
}

func TestAstParseShortenedModule(t *testing.T) {
	src := `(data $d1 "one") (data $d2 "two")`

	astRoot, err := Parse(strings.NewReader(src))
	if err != nil {
		if ut, ok := err.(*UnexpectedToken); ok {
			t.Fatalf("Parse error: %s\n| %s\n| %s", err.Error(), src, strings.Repeat(" ", ut.GotToken.Pos)+"^")
		} else {
			t.Fatal(err)
		}
	}

	expectedAst := astNodeChain(
		&Node{
			Kind: NodeKind_SubNode,
			Name: "data",
			FirstChild: astNodeChain(
				&Node{
					Kind:     NodeKind_AttrIdentifier,
					StrValue: "d1",
				},
				&Node{
					Kind:     NodeKind_AttrString,
					StrValue: "one",
				},
			),
		},
		&Node{
			Kind: NodeKind_SubNode,
			Name: "data",
			FirstChild: astNodeChain(
				&Node{
					Kind:     NodeKind_AttrIdentifier,
					StrValue: "d2",
				},
				&Node{
					Kind:     NodeKind_AttrString,
					StrValue: "two",
				},
			),
		},
	)

	if err, tok := compareNodes(expectedAst, astRoot); err != nil {
		code, arrow := getTokenContext(src, tok)
		t.Error("got wrong ast:", err, "\n| "+code+"\n| "+arrow)
	}
}

func TestAstParseSDL3Wat(t *testing.T) {
	src := string(sdl3Wat)

	astRoot, err := Parse(strings.NewReader(src))
	if err != nil {
		if ut, ok := err.(*UnexpectedToken); ok {
			offset := ut.GotToken.Pos - 16
			length := 32
			t.Fatalf("Parse error: %s\n| %s\n| %s", err.Error(), src[offset:offset+length], strings.Repeat(" ", ut.GotToken.Pos-offset)+"^")
		} else {
			t.Fatal(err)
		}
	}

	printNode(astRoot, t)
}

func astNodeChain(nodes ...*Node) *Node {
	if len(nodes) == 0 {
		return nil
	}

	for i := 0; i+1 < len(nodes); i++ {
		nodes[i].NextSibling = nodes[i+1]
	}

	return nodes[0]
}

func getTokenContext(input string, tok *tokenizer.Token) (line, arrow string) {
	lo := max(0, tok.Pos-32)
	hi := min(len(input), tok.Pos+32)

	return input[lo:hi], strings.Repeat(" ", tok.Pos-lo) + "^"
}

func compareNodes(expected, got *Node) (error, *tokenizer.Token) {
	if expected.Kind != got.Kind {
		return fmt.Errorf("expected node %s (%s), got %s (%s)", expected, expected.Kind, got, got.Kind), got.ParenOpenToken
	}

	switch got.Kind {
	case NodeKind_AttrFloat:
		if expected.Kind != NodeKind_AttrFloat {
			return fmt.Errorf("expected float attribute with value \"%f\", got \"%s\"", expected.FloatValue, got), got.ValueToken
		}

		if expected.FloatValue != got.FloatValue {
			return fmt.Errorf("expected float attribute with value %f, got %f", expected.FloatValue, got.FloatValue), got.ValueToken
		}
	case NodeKind_AttrIdentifier:
		if expected.Kind != NodeKind_AttrIdentifier {
			return fmt.Errorf("expected identifier attribute \"%s\", got \"%s\"", expected, got), got.ValueToken
		}

		if expected.StrValue != got.StrValue {
			return fmt.Errorf("expected identifier %s, got %s", expected, got), got.ValueToken
		}
	case NodeKind_AttrInteger:
		if expected.Kind != NodeKind_AttrInteger {
			return fmt.Errorf("expected integer attribute with value %d, got %s", expected.IntValue, got), got.ValueToken
		}

		if expected.IntValue != got.IntValue {
			return fmt.Errorf("expected integer attribute with value %s, got %s", expected, got), got.ValueToken
		}

	case NodeKind_AttrKeyword:
		if expected.Kind != NodeKind_AttrKeyword {
			return fmt.Errorf("expected keyword attribute %s, got %s", expected, got), got.ValueToken
		}

		if expected.StrValue != got.StrValue {
			return fmt.Errorf("expected keyword %s, got %s", expected.StrValue, got.StrValue), got.ValueToken
		}
	case NodeKind_AttrString:
		if expected.Kind != NodeKind_AttrString {
			return fmt.Errorf("expected string attribute %s, got %s", expected, got), got.ValueToken
		}

		if expected.StrValue != got.StrValue {
			return fmt.Errorf("expected string %s, got %s", expected.StrValue, got.StrValue), got.ValueToken
		}

	case NodeKind_SubNode:
		if expected.Name != got.Name {
			return fmt.Errorf("expected node name \"%s\", got \"%s\"", expected.Name, got.Name), got.NameToken
		}

		if err, tok := compareNodes(expected.FirstChild, got.FirstChild); err != nil {
			return err, tok
		}

	default:
		panic(fmt.Sprintf("unexpected wat.AstNodeKind: %#v", got.Kind))
	}

	tok := got.ValueToken
	if tok == nil {
		tok = got.NameToken
	}

	if expected.NextSibling != nil && got.NextSibling == nil {
		return fmt.Errorf("expected node \"%s\", got nothing", expected.NextSibling), tok
	} else if expected.NextSibling == nil && got.NextSibling != nil {
		return fmt.Errorf("unexpected node \"%s\"", got.NextSibling), tok
	} else if expected.NextSibling != nil {
		return compareNodes(expected.NextSibling, got.NextSibling)
	}

	return nil, nil
}

func printNode(curNode *Node, t *testing.T) {
	t.Logf("(%s", curNode.Name)
	for child := curNode.FirstChild; child != nil; child = child.NextSibling {
		switch child.Kind {
		case NodeKind_SubNode:
			printNode(child, t)

		case NodeKind_AttrFloat:
			t.Logf(" %f", child.FloatValue)
		case NodeKind_AttrIdentifier:
			t.Logf(" $%s", child.StrValue)
		case NodeKind_AttrInteger:
			t.Logf(" %d", child.IntValue)
		case NodeKind_AttrKeyword:
			t.Logf(" %s", child.StrValue)
		case NodeKind_AttrString:
			t.Logf(" %#v", child.StrValue)
		}
	}

	t.Log(")")
}
