package tokenizer

import "strconv"

type Token struct {
	Kind       TokenKind
	Pos        int // byte position in input stream
	Line, Col  int
	Len        int    // number of bytes in input stream
	Content    string // name for idents, text for comments
	IntValue   int64
	FloatValue float64
}

type TokenKind int

const (
	Invalid       TokenKind = iota
	Keyword                 // import
	Identifier              // $global0
	String                  // "env"
	DecimalNumber           // -42
	FloatNumber             // -3.14
	ParenOpen               // (
	ParenClose              // )
	LineComment             // ;; comment
	BlockComment            // (; comment ;)
)

func (tok *Token) String() string {
	switch tok.Kind {
	case BlockComment:
		return "(; " + tok.Content + " ;)"
	case DecimalNumber:
		return strconv.FormatInt(tok.IntValue, 10)
	case FloatNumber:
		return strconv.FormatFloat(tok.FloatValue, 'f', 1, 64)
	case Identifier:
		return "$" + tok.Content
	case Keyword:
		return tok.Content
	case LineComment:
		return ";;" + tok.Content
	case ParenClose:
		return ")"
	case ParenOpen:
		return "("
	case String:
		return "\"" + tok.Content + "\""
	}

	return "???"
}

func (tk TokenKind) String() string {
	switch tk {
	case Invalid:
		return "Invalid"
	case BlockComment:
		return "BlockComment"
	case DecimalNumber:
		return "DecimalNumber"
	case FloatNumber:
		return "FloatNumber"
	case Identifier:
		return "Identifier"
	case Keyword:
		return "Keyword"
	case LineComment:
		return "LineComment"
	case ParenClose:
		return "ParenClose"
	case ParenOpen:
		return "ParenOpen"
	case String:
		return "String"
	}

	return "???"
}
