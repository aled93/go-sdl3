package wat

import (
	"errors"
	"io"
	"strconv"
	"strings"
)

var (
	ErrExcessNodeClosing = errors.New("excess node closing")
)

func parseAst(r io.Reader) (root *AstNode, err error) {
	type State int
	const (
		State_Initial State = iota
		State_PreNode
		State_NodeName
		State_NodeInner
	)

	toker := NewTokenizer(r)
	state := State_Initial
	root = &AstNode{}
	curNode := root
	var lastChild *AstNode
	var nodeStack []*AstNode
	var lastChildStack []*AstNode

	for {
		tok, err := toker.NextToken()

		switch state {
		case State_PreNode:
			// (import "env" ...)
			// ^
			curNode.NextSibling = &AstNode{}
			curNode = curNode.NextSibling
			fallthrough

		case State_Initial:
			if tok.Kind != ParenOpen {
				return nil, &UnexpectedToken{
					GotToken: tok,
					ExpectedTokens: []TokenKind{
						ParenOpen,
					},
				}
			}

			curNode.ParenOpenToken = &tok

			state = State_NodeName

		case State_NodeName:
			// (import "env" ...)
			//  ^
			if tok.Kind != Keyword {
				return nil, &UnexpectedToken{
					GotToken: tok,
					ExpectedTokens: []TokenKind{
						Keyword,
					},
				}
			}

			curNode.NameToken = &tok
			curNode.Name = tok.Content
			state = State_NodeInner

		case State_NodeInner:
			// (import "env" ...)
			//         ^

			var newAttr *AstNode
			switch tok.Kind {
			case ParenOpen:
				nodeStack = append(nodeStack, curNode)

				newNode := &AstNode{}
				if lastChild != nil {
					lastChild.NextSibling = newNode
				} else {
					curNode.FirstChild = newNode
				}
				lastChild = newNode
				curNode = newNode

				lastChildStack = append(lastChildStack, lastChild)
				lastChild = nil

				state = State_NodeName

			case ParenClose:
				if len(nodeStack) == 0 {
					// done
					return root, nil
				}

				curNode.ParenCloseToken = &tok

				curNode = nodeStack[len(nodeStack)-1]
				nodeStack = nodeStack[:len(nodeStack)-1]
				lastChild = lastChildStack[len(lastChildStack)-1]
				lastChildStack = lastChildStack[:len(lastChildStack)-1]

				state = State_NodeInner

			case Keyword:
				newAttr = &AstNode{
					Kind:       NodeKind_AttrKeyword,
					StrValue:   tok.Content,
					ValueToken: &tok,
				}
				fallthrough

			case Identifier:
				if newAttr == nil {
					newAttr = &AstNode{
						Kind:       NodeKind_AttrIdentifier,
						StrValue:   tok.Content,
						ValueToken: &tok,
					}
				}
				fallthrough

			case String:
				if newAttr == nil {
					newAttr = &AstNode{
						Kind:       NodeKind_AttrString,
						StrValue:   tok.Content,
						ValueToken: &tok,
					}
				}
				fallthrough

			case DecimalNumber:
				if newAttr == nil {
					newAttr = &AstNode{
						Kind:       NodeKind_AttrInteger,
						IntValue:   tok.IntValue,
						ValueToken: &tok,
					}
				}
				fallthrough

			case FloatNumber:
				if newAttr == nil {
					newAttr = &AstNode{
						Kind:       NodeKind_AttrFloat,
						FloatValue: tok.FloatValue,
						ValueToken: &tok,
					}
				}

				if lastChild == nil {
					curNode.FirstChild = newAttr
				} else {
					lastChild.NextSibling = newAttr
				}
				lastChild = newAttr
			}

		}

		if err != nil {
			return nil, err
		}
	}
}

type AstNode struct {
	Kind            AstNodeKind
	Name            string
	StrValue        string
	IntValue        int64
	FloatValue      float64
	FirstChild      *AstNode
	NextSibling     *AstNode
	NameToken       *Token
	ParenOpenToken  *Token
	ParenCloseToken *Token
	ValueToken      *Token
}

var escapeStr = strings.NewReplacer(
	"\t", "\\t",
	"\n", "\\n",
	"\r", "\\r",
	"\\", "\\\\",
	"\"", "\\\"",
)

func (n *AstNode) String() string {
	switch n.Kind {
	case NodeKind_AttrFloat:
		return strconv.FormatFloat(n.FloatValue, 'f', 4, 64)
	case NodeKind_AttrIdentifier:
		return "$" + n.StrValue
	case NodeKind_AttrInteger:
		return strconv.FormatInt(n.IntValue, 10)
	case NodeKind_AttrKeyword:
		return n.StrValue
	case NodeKind_AttrString:
		return `"` + escapeStr.Replace(n.StrValue) + `"`
	case NodeKind_SubNode:
		return "(" + n.Name
	}

	panic("unhandled AstNodeKind variant")
}

type AstNodeKind int

const (
	NodeKind_SubNode AstNodeKind = iota
	NodeKind_AttrKeyword
	NodeKind_AttrIdentifier
	NodeKind_AttrString
	NodeKind_AttrInteger
	NodeKind_AttrFloat
)

func (k AstNodeKind) String() string {
	switch k {
	case NodeKind_AttrFloat:
		return "AttrFloat"
	case NodeKind_AttrIdentifier:
		return "AttrIdent"
	case NodeKind_AttrInteger:
		return "AttrInt"
	case NodeKind_AttrKeyword:
		return "AttrKeyword"
	case NodeKind_AttrString:
		return "AttrString"
	case NodeKind_SubNode:
		return "Subnode"
	}

	panic("unhandled AstNodeKind variant")
}

type UnexpectedToken struct {
	GotToken       Token
	ExpectedTokens []TokenKind
	ExpectedWhat   string
}

func (ut *UnexpectedToken) Error() string {
	var sb strings.Builder

	sb.WriteString("unexpected token")
	if ut.GotToken.Kind != Invalid {
		sb.WriteRune(' ')
		sb.WriteString(ut.GotToken.Kind.String())
		sb.WriteString(" \"")
		sb.WriteString(ut.GotToken.String())
		sb.WriteString("\"")
	}

	if len(ut.ExpectedTokens) == 1 {
		sb.WriteString(", expected ")
		sb.WriteString(ut.ExpectedTokens[0].String())
	} else if len(ut.ExpectedTokens) > 1 {
		sb.WriteString(", expected one of: ")
		for i, tok := range ut.ExpectedTokens {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(tok.String())
		}
	} else if len(ut.ExpectedWhat) > 0 {
		sb.WriteString(", ")
		sb.WriteString(ut.ExpectedWhat)
	}

	return sb.String()
}
