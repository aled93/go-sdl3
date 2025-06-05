package ast

import (
	"errors"
	"io"
	"sdl3/cmd/embox/pkg/wat/internal/tokenizer"
	"strings"
)

var (
	ErrExcessNodeClosing = errors.New("excess node closing")
)

func Parse(r io.Reader) (root *Node, err error) {
	type State int
	const (
		State_PreNode State = iota
		State_NodeName
		State_NodeInner
	)

	toker := tokenizer.NewTokenizer(r)
	state := State_PreNode
	var curNode *Node
	var prevNode *Node
	var lastChild *Node
	var nodeStack []*Node
	var lastChildStack []*Node

	for {
		tok, err := toker.NextToken()

		switch state {
		case State_PreNode:
			// (import "env" ...)
			// ^

			if tok.Kind == tokenizer.Invalid {
				// done
				return root, nil
			}

			if tok.Kind != tokenizer.ParenOpen {
				return nil, &UnexpectedToken{
					GotToken: tok,
					ExpectedTokens: []tokenizer.TokenKind{
						tokenizer.ParenOpen,
					},
				}
			}

			curNode = &Node{}
			if root == nil {
				root = curNode
			}
			if prevNode != nil {
				prevNode.NextSibling = curNode
				prevNode = nil
			}

			curNode.ParenOpenToken = &tok

			state = State_NodeName

		case State_NodeName:
			// (import "env" ...)
			//  ^
			if tok.Kind != tokenizer.Keyword {
				return nil, &UnexpectedToken{
					GotToken: tok,
					ExpectedTokens: []tokenizer.TokenKind{
						tokenizer.Keyword,
					},
				}
			}

			curNode.NameToken = &tok
			curNode.Name = tok.Content
			state = State_NodeInner

		case State_NodeInner:
			// (import "env" ...)
			//         ^

			var newAttr *Node
			switch tok.Kind {
			case tokenizer.ParenOpen:
				nodeStack = append(nodeStack, curNode)

				newNode := &Node{
					ParenOpenToken: &tok,
				}
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

			case tokenizer.ParenClose:
				curNode.ParenCloseToken = &tok

				if len(nodeStack) == 0 {
					prevNode = curNode
					curNode = nil
					lastChild = nil

					state = State_PreNode
				} else {
					curNode = nodeStack[len(nodeStack)-1]
					nodeStack = nodeStack[:len(nodeStack)-1]
					lastChild = lastChildStack[len(lastChildStack)-1]
					lastChildStack = lastChildStack[:len(lastChildStack)-1]

					state = State_NodeInner
				}

			case tokenizer.Keyword:
				newAttr = &Node{
					Kind:       NodeKind_AttrKeyword,
					StrValue:   tok.Content,
					ValueToken: &tok,
				}
				fallthrough

			case tokenizer.Identifier:
				if newAttr == nil {
					newAttr = &Node{
						Kind:       NodeKind_AttrIdentifier,
						StrValue:   tok.Content,
						ValueToken: &tok,
					}
				}
				fallthrough

			case tokenizer.String:
				if newAttr == nil {
					newAttr = &Node{
						Kind:       NodeKind_AttrString,
						StrValue:   tok.Content,
						ValueToken: &tok,
					}
				}
				fallthrough

			case tokenizer.DecimalNumber:
				if newAttr == nil {
					newAttr = &Node{
						Kind:       NodeKind_AttrInteger,
						StrValue:   tok.Content,
						ValueToken: &tok,
					}
				}
				fallthrough

			case tokenizer.FloatNumber:
				if newAttr == nil {
					newAttr = &Node{
						Kind:       NodeKind_AttrFloat,
						StrValue:   tok.Content,
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

type Node struct {
	Kind            NodeKind
	Name            string
	StrValue        string
	FirstChild      *Node
	NextSibling     *Node
	NameToken       *tokenizer.Token
	ParenOpenToken  *tokenizer.Token
	ParenCloseToken *tokenizer.Token
	ValueToken      *tokenizer.Token
}

var escapeStr = strings.NewReplacer(
	"\t", "\\t",
	"\n", "\\n",
	"\r", "\\r",
	"\\", "\\\\",
	"\"", "\\\"",
)

func (n *Node) FirstToken() *tokenizer.Token {
	if n.ParenOpenToken != nil {
		return n.ParenOpenToken
	}

	return n.ValueToken
}

func (n *Node) LastToken() *tokenizer.Token {
	if n.ParenCloseToken != nil {
		return n.ParenCloseToken
	}

	return n.ValueToken
}

func (n *Node) String() string {
	if n == nil {
		return "nothing"
	}

	switch n.Kind {
	case NodeKind_AttrFloat:
		return n.StrValue
	case NodeKind_AttrIdentifier:
		return "$" + n.StrValue
	case NodeKind_AttrInteger:
		return n.StrValue
	case NodeKind_AttrKeyword:
		return n.StrValue
	case NodeKind_AttrString:
		return `"` + escapeStr.Replace(n.StrValue) + `"`
	case NodeKind_SubNode:
		return "(" + n.Name
	}

	panic("unhandled NodeKind variant")
}

func (node *Node) WriteTo(w io.Writer) (n int64, err error) {
	switch node.Kind {
	case NodeKind_AttrFloat:
		wn, werr := w.Write([]byte(node.StrValue))
		return int64(wn), werr

	case NodeKind_AttrIdentifier:
		wn, werr := w.Write([]byte("$" + node.StrValue))
		return int64(wn), werr

	case NodeKind_AttrInteger:
		wn, werr := w.Write([]byte(node.StrValue))
		return int64(wn), werr

	case NodeKind_AttrKeyword:
		wn, werr := w.Write([]byte(node.StrValue))
		return int64(wn), werr

	case NodeKind_AttrString:
		wn, werr := escapeStr.WriteString(w, node.StrValue)
		return int64(wn), werr

	case NodeKind_SubNode:
		wn, werr := w.Write([]byte("(" + node.Name))
		n += int64(wn)
		if werr != nil {
			return n, werr
		}

		for subnode := node.FirstChild; subnode != nil; subnode = subnode.NextSibling {
			var sn int64

			wn, werr = w.Write([]byte(" "))
			n += int64(wn)
			if werr != nil {
				return n, werr
			}

			sn, werr = subnode.WriteTo(w)
			n += sn
			if werr != nil {
				return n, werr
			}
		}

		wn, werr = w.Write([]byte(")"))
		n += int64(wn)
		if werr != nil {
			return n, werr
		}

	default:
		panic("unknown node kind")
	}

	return n, nil
}

type NodeKind int

const (
	NodeKind_SubNode NodeKind = iota
	NodeKind_AttrKeyword
	NodeKind_AttrIdentifier
	NodeKind_AttrString
	NodeKind_AttrInteger
	NodeKind_AttrFloat
)

func (k NodeKind) String() string {
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

	panic("unhandled NodeKind variant")
}

type UnexpectedToken struct {
	GotToken       tokenizer.Token
	ExpectedTokens []tokenizer.TokenKind
	ExpectedWhat   string
}

func (ut *UnexpectedToken) Error() string {
	var sb strings.Builder

	sb.WriteString("unexpected token")
	if ut.GotToken.Kind != tokenizer.Invalid {
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
