package tokenizer

import (
	"bufio"
	"errors"
	"io"
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
	trk *LineEndTracker
	pos int
}

func NewTokenizer(r io.Reader) *Tokenizer {
	trk := NewLineEndTracker(r)
	return &Tokenizer{
		r:   *bufio.NewReader(trk),
		trk: trk,
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
	tok.Line, tok.Col = t.trk.LineCol(t.pos)

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
	var chs []byte
	digitsPos := 0
	tok.Kind = DecimalNumber

	neg := false
	if t.expectRune('-') {
		neg = true
		chs = append(chs, '-')
		digitsPos = 1
	} else if t.expectRune('+') {
		neg = false
		chs = append(chs, '+')
		digitsPos = 1
	}

	if t.expectRune('i') {
		if t.expectRune('n') && t.expectRune('f') {
			tok.Kind = FloatNumber
			if neg {
				tok.Content = "-inf"
			} else {
				tok.Content = "inf"
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
			tok.Content = "nan"
			return
		}

		panic(&UnexpectedRune{
			UnparsedToken: *tok,
			GotRune:       t.peekRune(),
			ExpectedWhat:  "expected \"nan\"",
		})
	}

	for r := t.peekRune(); r == '_' || (r >= '0' && r <= '9'); r = t.peekRune() {
		chs = append(chs, byte(t.readRune()))
	}

	if len(chs)-digitsPos == 0 {
		panic(&UnexpectedRune{
			UnparsedToken: *tok,
			ExpectedWhat:  "missing integer part of number",
		})
	}

	if t.expectRune('x') {
		tok.Flags |= IsHex

		if len(chs)-digitsPos != 1 || chs[digitsPos] != '0' {
			panic(&UnexpectedRune{
				UnparsedToken: *tok,
				ExpectedWhat:  "hex must be prefixed with exact \"0x\"",
			})
		}

		chs = append(chs, 'x')
		chs = t.readUint(tok, chs, 1)
		if t.expectRune('.') {
			chs = append(chs, '.')
			t.readUint(tok, nil, 1)
		}
		if t.expectRune('p') || t.expectRune('P') {
			chs = append(chs, 'p', byte(t.peekRune()))
			if !t.expectRune('+') && !t.expectRune('-') {
				panic(&UnexpectedRune{
					UnparsedToken: *tok,
					GotRune:       t.peekRune(),
					ExpectedRunes: []rune{'+', '-'},
				})
			}

			chs = t.readUint(tok, chs, 0)
		}
	} else {
		if t.expectRune('.') {
			tok.Kind = FloatNumber
			chs = append(chs, '.')
			chs = t.readUint(tok, chs, 0)
		}
		if t.expectRune('e') || t.expectRune('E') {
			chs = append(chs, 'e')
			if t.expectRune('+') {
				chs = append(chs, '+')
			} else if t.expectRune('-') {
				chs = append(chs, '-')
			}

			chs = t.readUint(tok, chs, 0)
		}
	}
}

func (t *Tokenizer) readUint(tok *Token, digits []byte, hex int) []byte {
	if hex == 0 {
		for r := t.peekRune(); r == '_' || (r >= '0' && r <= '9'); r = t.peekRune() {
			digits = append(digits, byte(t.readRune()))
		}
	} else {
		for r := t.peekRune(); r == '_' ||
			(r >= '0' && r <= '9') ||
			(r >= 'A' && r <= 'F') ||
			(r >= 'a' && r <= 'F'); r = t.peekRune() {
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

	return digits
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
				cp := t.readUint(tok, nil, 1)
				i, err := strconv.Atoi(string(cp))
				if err != nil {
					panic(err)
				}

				if i > utf8.MaxRune {
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
