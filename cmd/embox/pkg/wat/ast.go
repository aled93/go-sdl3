package wat

import (
	"errors"
	"io"
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
	var lastAttr *AstAttr
	var lastChild *AstNode
	var nodeStack []*AstNode
	var lastAttrStack []*AstAttr
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

			curNode.Name = tok.Content
			state = State_NodeInner

		case State_NodeInner:
			// (import "env" ...)
			//         ^

			var newAttr *AstAttr
			switch tok.Kind {
			case ParenOpen:
				nodeStack = append(nodeStack, curNode)

				newNode := &AstNode{}
				if lastChild != nil {
					lastChild.NextSibling = newNode
				} else {
					curNode.FirstChild = newNode
				}
				curNode = newNode

				lastChildStack = append(lastChildStack, lastChild)
				lastChild = nil
				lastAttrStack = append(lastAttrStack, lastAttr)
				lastAttr = nil

				state = State_NodeName

			case ParenClose:
				if len(nodeStack) == 0 {
					// done
					return root, nil
				}

				curNode = nodeStack[len(nodeStack)-1]
				nodeStack = nodeStack[:len(nodeStack)-1]
				lastChild = lastChildStack[len(lastChildStack)-1]
				lastChildStack = lastChildStack[:len(lastChildStack)-1]
				lastAttr = lastAttrStack[len(lastAttrStack)-1]
				lastAttrStack = lastAttrStack[:len(lastAttrStack)-1]

				state = State_NodeInner

			case Keyword:
				newAttr = &AstAttr{
					Kind:     AttrKeyword,
					StrValue: tok.Content,
				}
				fallthrough

			case Identifier:
				if newAttr == nil {
					newAttr = &AstAttr{
						Kind:     AttrIdentifier,
						StrValue: tok.Content,
					}
				}
				fallthrough

			case String:
				if newAttr == nil {
					newAttr = &AstAttr{
						Kind:     AttrString,
						StrValue: tok.Content,
					}
				}
				fallthrough

			case DecimalNumber:
				if newAttr == nil {
					newAttr = &AstAttr{
						Kind:     AttrInteger,
						IntValue: tok.IntValue,
					}
				}
				fallthrough

			case FloatNumber:
				if newAttr == nil {
					newAttr = &AstAttr{
						Kind:       AttrFloat,
						FloatValue: tok.FloatValue,
					}
				}

				if lastAttr == nil {
					curNode.FirstAttr = newAttr
				} else {
					lastAttr.NextAttr = newAttr
				}
				lastAttr = newAttr
			}

		}

		if err != nil {
			return nil, err
		}
	}
}

type AstNode struct {
	Name        string
	FirstAttr   *AstAttr
	FirstChild  *AstNode
	NextSibling *AstNode
}

type AstAttrKind int

const (
	AttrKeyword AstAttrKind = iota
	AttrIdentifier
	AttrString
	AttrInteger
	AttrFloat
)

type AstAttr struct {
	NextAttr   *AstAttr
	Kind       AstAttrKind
	StrValue   string
	IntValue   int64
	FloatValue float64
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
