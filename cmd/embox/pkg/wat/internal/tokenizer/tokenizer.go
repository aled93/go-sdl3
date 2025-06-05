package tokenizer

import (
	"bufio"
	"errors"
	"io"
	"math"
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	ErrInvalidInputUtf8 = errors.New("invalid input utf8 sequence")
	ErrNumberOverflow   = errors.New("number value overflows")
)

var (
	idRunes = []rune{
		'!', '#', '$', '%', '&', '\'', '*', '+', '-', '.', '/',
		':', '<', '=', '>', '?', '@', '\\', '^', '_', '`', '|', '~',
	}
)

type Tokenizer struct {
	r   bufio.Reader
	pos int
}

func NewTokenizer(r io.Reader) *Tokenizer {
	return &Tokenizer{
		r: *bufio.NewReader(r),
	}
}

func (t *Tokenizer) NextToken() (tok Token, err error) {
	defer func() {
		if r := recover(); r == ErrInvalidInputUtf8 {
			err = ErrInvalidInputUtf8
		} else if r != nil {
			panic(r)
		}
	}()

	// skip spaces
	for {
		r := t.peekRune()

		if r == '\u0000' {
			return tok, io.EOF
		}
		if !(r == ' ' || r == '\t' || r == '\r' || r == '\n') {
			break
		}

		t.readRune()
	}

	tok.Pos = t.pos

	r := t.peekRune()

	switch {
	case r == '(':
		t.readRune()

		if t.expectRune(';') {
			tok.Kind = BlockComment
			t.readBlockComment(&tok, &err)
		} else {
			tok.Kind = ParenOpen
		}

	case r == ')':
		t.readRune()
		tok.Kind = ParenClose

	case r == '+' || r == '-' || (r >= '0' && r <= '9'):
		t.readNumber(&tok, &err)

	case r >= 'a' && r <= 'z':
		tok.Kind = Keyword
		t.readKeyword(&tok, &err)

	case r == '$':
		tok.Kind = Identifier
		t.readRune()
		t.readId(&tok, &err)

	case r == ';':
		tok.Kind = LineComment
		t.readRune()
		if !t.expectRune(';') {
			return tok, &UnexpectedRune{
				UnparsedToken: tok,
				GotRune:       t.peekRune(),
				ExpectedRunes: []rune{';'},
			}
		}
		t.readLineComment(&tok, &err)

	case r == '"':
		tok.Kind = String
		t.readRune()
		t.readString(&tok, &err)

	default:
		return tok, &UnexpectedRune{
			UnparsedToken: tok,
			GotRune:       t.peekRune(),
		}
	}

	return
}

func (t *Tokenizer) readRune() rune {
	r, n, err := t.r.ReadRune()
	t.pos += n

	if r == utf8.RuneError {
		panic(ErrInvalidInputUtf8)
	}
	if err == io.EOF {
		return '\u0000'
	}
	if err != nil {
		panic(err)
	}

	return r
}

func (t *Tokenizer) peekRune() rune {
	r := t.readRune()
	if r != '\u0000' {
		t.r.UnreadRune()
		t.pos -= utf8.RuneLen(r)
	}
	return r
}

func (t *Tokenizer) expectRune(r rune) bool {
	rr, _, err := t.r.ReadRune()

	if rr == utf8.RuneError || err != nil {
		return false
	}

	if r == rr {
		return true
	}

	t.r.UnreadRune()
	return false
}

func (t *Tokenizer) readBlockComment(tok *Token, err *error) {
	var content []byte

	for {
		r := t.readRune()
		if r == ';' && t.expectRune(')') {
			break
		} else {
			content = utf8.AppendRune(content, r)
		}
	}

	tok.Content = string(content)
}

func (t *Tokenizer) readLineComment(tok *Token, err *error) {
	var content []byte

	for {
		r := t.readRune()
		if r == '\u0000' || r == '\n' || r == '\r' {
			if r == '\r' {
				t.expectRune('\n')
			}
			break
		}
	}

	tok.Content = string(content)
}

func (t *Tokenizer) readKeyword(tok *Token, err *error) {
	var content []byte

	for {
		r := t.peekRune()
		if (r >= '0' && r <= '9') ||
			(r >= 'a' && r <= 'z') ||
			(r >= 'A' && r < 'Z') ||
			(slices.Contains(idRunes, r)) {
			t.readRune()
			content = utf8.AppendRune(content, r)
		} else {
			break
		}
	}

	tok.Content = string(content)
}

func (t *Tokenizer) readId(tok *Token, err *error) {
	var content []byte

	for r := t.peekRune(); (r >= '0' && r <= '9') ||
		(r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') ||
		slices.Contains(idRunes, r); r = t.peekRune() {
		content = utf8.AppendRune(content, t.readRune())
	}

	tok.Content = string(content)
}

func (t *Tokenizer) readNumber(tok *Token, err *error) {
	neg := false
	if t.expectRune('-') {
		neg = true
	} else if t.expectRune('+') {
		neg = false
	}

	if t.expectRune('i') {
		if t.expectRune('n') && t.expectRune('f') {
			tok.Kind = FloatNumber
			if neg {
				tok.FloatValue = math.Inf(-1)
			} else {
				tok.FloatValue = math.Inf(1)
			}
			return
		}

		panic(&UnexpectedRune{
			UnparsedToken: *tok,
			GotRune:       t.peekRune(),
			ExpectedWhat:  "expected \"inf\"",
		})
	}

	if t.expectRune('n') {
		if t.expectRune('a') && t.expectRune('n') {
			tok.Kind = FloatNumber
			tok.FloatValue = math.NaN()
			return
		}

		panic(&UnexpectedRune{
			UnparsedToken: *tok,
			GotRune:       t.peekRune(),
			ExpectedWhat:  "expected \"nan\"",
		})
	}

	var digits []byte
	for r := t.peekRune(); r == '_' || (r >= '0' && r <= '9'); r = t.peekRune() {
		digits = append(digits, byte(t.readRune()))
	}

	if len(digits) == 0 {
		panic(&UnexpectedRune{
			UnparsedToken: *tok,
			ExpectedWhat:  "missing integer part of number",
		})
	}

	var intPart, fracPart, expPart uint64
	isHex := false
	hasExp := false
	expNeg := false
	isFloat := false

	expPart = 1

	if v, err := strconv.Atoi(string(digits)); err != nil {
		panic(err)
	} else {
		intPart = uint64(v)
	}

	if t.expectRune('x') {
		isHex = true
		if len(digits) != 1 || digits[0] != '0' {
			panic(&UnexpectedRune{
				UnparsedToken: *tok,
				ExpectedWhat:  "hex must be prefixed with exact \"0x\"",
			})
		}

		intPart = t.readAndParseUint(tok, digits, 1)
		if t.expectRune('.') {
			isFloat = true
			fracPart = t.readAndParseUint(tok, nil, 1)
		}
		if t.expectRune('p') || t.expectRune('P') {
			hasExp = true
			if t.expectRune('+') {
				expNeg = false
			} else if t.expectRune('-') {
				expNeg = true
			} else {
				panic(&UnexpectedRune{
					UnparsedToken: *tok,
					GotRune:       t.peekRune(),
					ExpectedRunes: []rune{'+', '-'},
				})
			}

			expPart = t.readAndParseUint(tok, nil, 0)
		}
	} else {
		if t.expectRune('.') {
			isFloat = true
			fracPart = t.readAndParseUint(tok, nil, 0)
		}
		if t.expectRune('e') || t.expectRune('E') {
			hasExp = true
			if t.expectRune('+') {
				expNeg = false
			} else if t.expectRune('-') {
				expNeg = true
			}

			expPart = t.readAndParseUint(tok, nil, 0)
		}
	}

	if isFloat {
		tok.Kind = FloatNumber
	} else {
		tok.Kind = DecimalNumber
	}

	if !isFloat && !hasExp {
		tok.Kind = DecimalNumber
		tok.IntValue = int64(intPart)
		if neg {
			tok.IntValue = -tok.IntValue
		}
	} else {
		// delegate to checked solutions :)
		var str []byte
		if neg {
			str = append(str, byte('-'))
		}

		if isHex {
			str = append(str, '0', 'x')
		}

		base := 10
		if isHex {
			base = 16
		}
		str = strconv.AppendUint(str, intPart, base)

		if isFloat {
			str = append(str, byte('.'))
			str = strconv.AppendUint(str, fracPart, base)
		}

		if hasExp || isHex {
			if isHex {
				str = append(str, byte('p'))
			} else {
				str = append(str, byte('e'))
			}
			if expNeg {
				str = append(str, byte('-'))
			} else {
				str = append(str, byte('+'))
			}
			str = strconv.AppendUint(str, expPart, 10)
		}

		if isFloat {
			tok.Kind = FloatNumber
			f, e := strconv.ParseFloat(string(str), 64)
			if e != nil {
				panic(e)
			}
			tok.FloatValue = f
		}
	}
}

func (t *Tokenizer) readAndParseUint(tok *Token, digits []byte, hex int) (val uint64) {
	base := 10
	if hex > 1 {
		base = 16
	}

	if hex == 0 {
		for r := t.peekRune(); r == '_' || (r >= '0' && r <= '9'); r = t.peekRune() {
			digits = append(digits, byte(t.readRune()))
		}
	} else {
		for r := t.peekRune(); r == '_' ||
			(r >= '0' && r <= '9') ||
			(r >= 'A' && r <= 'F') ||
			(r >= 'a' && r <= 'F'); r = t.peekRune() {
			if r < '0' || r > '9' {
				base = 16
			}
			digits = append(digits, byte(t.readRune()))
		}
	}

	if len(digits) > 0 && (digits[0] == '_' || digits[len(digits)-1] == '_') {
		panic(&UnexpectedRune{
			UnparsedToken: *tok,
			GotRune:       t.peekRune(),
			ExpectedWhat:  "underscores must be between digits",
		})
	}

	v, e := strconv.ParseUint(string(digits), base, 64)
	if e == strconv.ErrRange {
		panic(ErrNumberOverflow)
	} else if e != nil {
		panic(e)
	}

	val = v

	return
}

func (t *Tokenizer) readString(tok *Token, err *error) {
	var content []byte

	for {
		r := t.readRune()

		if r == '\\' {
			r = t.readRune()

			switch r {
			case 't':
				content = append(content, '\t')
			case 'n':
				content = append(content, '\n')
			case 'r':
				content = append(content, '\r')
			case '\'':
				content = append(content, '\'')
			case '"':
				content = append(content, '"')
			case '\\':
				content = append(content, '\\')
			case 'u':
				cp := t.readAndParseUint(nil, nil, 1)
				if cp > utf8.MaxRune {
					panic(UnexpectedRune{
						UnparsedToken: *tok,
						ExpectedWhat:  "unicode code point out of range",
					})
				}
			}
		} else {
			if r == '"' {
				break
			}

			content = utf8.AppendRune(content, r)
		}
	}

	tok.Content = string(content)
}

type UnexpectedRune struct {
	UnparsedToken Token
	GotRune       rune
	ExpectedRunes []rune
	ExpectedWhat  string
}

func (ur *UnexpectedRune) Error() string {
	var str strings.Builder

	str.WriteString("unexpected rune")
	if ur.GotRune != '\u0000' {
		str.WriteRune(' ')
		str.WriteRune('"')
		str.WriteRune(ur.GotRune)
		str.WriteRune('"')
	}
	str.WriteString(" while parsing token ")
	str.WriteString(ur.UnparsedToken.Kind.String())

	if len(ur.ExpectedWhat) > 0 {
		str.WriteString(", ")
		str.WriteString(ur.ExpectedWhat)
	} else if len(ur.ExpectedRunes) > 0 {
		str.WriteString(", expected one of: ")
		for i, r := range ur.ExpectedRunes {
			if i > 0 {
				str.WriteString(", ")
			}
			str.WriteRune('"')
			str.WriteRune(r)
			str.WriteRune('"')
		}
	}

	return str.String()
}
