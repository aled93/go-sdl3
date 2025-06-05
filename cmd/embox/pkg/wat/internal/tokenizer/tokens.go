package tokenizer

import (
	"errors"
	"strconv"
)

type Token struct {
	Kind      TokenKind
	Flags     TokenFlags
	Pos       int // byte position in input stream
	Line, Col int
	Len       int    // number of bytes in input stream
	Content   string // name for idents, text for comments, digits for numbers
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

type TokenFlags uint8

const (
	IsHex TokenFlags = 1 << iota
)

func (tok *Token) String() string {
	switch tok.Kind {
	case BlockComment:
		return "(; " + tok.Content + " ;)"
	case DecimalNumber:
		return tok.Content
	case FloatNumber:
		return tok.Content
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

func (tok *Token) AsInt32() (int32, bool) {
	if tok.Kind != DecimalNumber {
		return 0, false
	}

	base := 10
	if (tok.Flags & IsHex) != 0 {
		base = 16
	}

	i, err := strconv.ParseInt(tok.Content, base, 32)
	if errors.Is(err, strconv.ErrRange) {
		return 0, false
	}

	if err != nil {
		panic(err)
	}

	return int32(i), true
}

func (tok *Token) AsInt64() (int64, bool) {
	if tok.Kind != DecimalNumber {
		return 0, false
	}

	base := 10
	if (tok.Flags & IsHex) != 0 {
		base = 16
	}

	i, err := strconv.ParseInt(tok.Content, base, 32)
	if errors.Is(err, strconv.ErrRange) {
		return 0, false
	}

	if err != nil {
		panic(err)
	}

	return i, true
}

func (tok *Token) AsFloat32() (float32, bool) {
	if tok.Kind != FloatNumber {
		return 0, false
	}

	f, err := strconv.ParseFloat(tok.Content, 32)
	if errors.Is(err, strconv.ErrRange) {
		return 0, false
	}

	if err != nil {
		panic(err)
	}

	return float32(f), true
}

func (tok *Token) AsFloat64() (float64, bool) {
	if tok.Kind != FloatNumber {
		return 0, false
	}

	f, err := strconv.ParseFloat(tok.Content, 64)
	if errors.Is(err, strconv.ErrRange) {
		return 0, false
	}

	if err != nil {
		panic(err)
	}

	return f, true
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
