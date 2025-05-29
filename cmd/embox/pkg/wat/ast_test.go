package wat

import (
	"strings"
	"testing"
)

func TestAstParseSimpleWat(t *testing.T) {
	src := `(module (import "js" "print" (func (param i32) ) ) )`

	astRoot, err := parseAst(strings.NewReader(src))
	if err != nil {
		if ut, ok := err.(*UnexpectedToken); ok {
			t.Fatalf("Parse error: %s\n| %s\n| %s", err.Error(), src, strings.Repeat(" ", ut.GotToken.Pos)+"^")
		} else {
			t.Fatal(err)
		}
	}

	printNode(astRoot, t)
}

func TestAstParseSDL3Wat(t *testing.T) {
	src := string(sdl3Wat)

	astRoot, err := parseAst(strings.NewReader(src))
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

func printNode(curNode *AstNode, t *testing.T) {
	t.Logf("(%s", curNode.Name)
	for attr := curNode.FirstAttr; attr != nil; attr = attr.NextAttr {
		switch attr.Kind {
		case AttrFloat:
			t.Logf(" %f", attr.FloatValue)
		case AttrIdentifier:
			t.Logf(" $%s", attr.StrValue)
		case AttrInteger:
			t.Logf(" %d", attr.IntValue)
		case AttrKeyword:
			t.Logf(" %s", attr.StrValue)
		case AttrString:
			t.Logf(" %#v", attr.StrValue)
		}
	}

	for child := curNode.FirstChild; child != nil; child = child.NextSibling {
		t.Log("\n")
		printNode(child, t)
	}

	t.Log(")")
}
