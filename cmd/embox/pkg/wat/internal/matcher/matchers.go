package matcher

import (
	"math"
	"regexp"
	"sdl3/cmd/embox/pkg/wat/internal/ast"
)

func Subnode(inner Matcher) *subnodeMatcher {
	return &subnodeMatcher{
		inner: inner,
	}
}

// NamedSubnode matches subnode with exact name
func NamedSubnode(name string, inner Matcher) *subnodeMatcher {
	return &subnodeMatcher{
		inner:     inner,
		exactName: true,
		name:      name,
	}
}

// NamePatternSubnode matches subnode testing name against regexp pattern
func NamePatternSubnode(namePattern *regexp.Regexp, inner Matcher) *subnodeMatcher {
	if namePattern == nil {
		panic("namePattern is nil")
	}

	return &subnodeMatcher{
		inner: inner,
		ptrn:  namePattern,
	}
}

func Keyword() *attr {
	return &attr{
		kind: ast.NodeKind_AttrKeyword,
	}
}

func KeywordExact(keyword string) *attr {
	return &attr{
		kind:     ast.NodeKind_AttrKeyword,
		exactStr: keyword,
	}
}

func KeywordPattern(pattern *regexp.Regexp) *attr {
	if pattern == nil {
		panic("pattern is nil")
	}

	return &attr{
		kind:    ast.NodeKind_AttrKeyword,
		ptrnStr: pattern,
	}
}

func Identifier() *attr {
	return &attr{
		kind: ast.NodeKind_AttrIdentifier,
	}
}

func IdentifierExact(ident string) *attr {
	return &attr{
		kind:     ast.NodeKind_AttrIdentifier,
		exactStr: ident,
	}
}

func IdentifierPattern(pattern *regexp.Regexp) *attr {
	if pattern == nil {
		panic("pattern is nil")
	}

	return &attr{
		kind:    ast.NodeKind_AttrIdentifier,
		ptrnStr: pattern,
	}
}

func StringExact(ident string) *attr {
	return &attr{
		kind:     ast.NodeKind_AttrString,
		exactStr: ident,
	}
}

func StringPattern(pattern *regexp.Regexp) *attr {
	if pattern == nil {
		panic("pattern is nil")
	}

	return &attr{
		kind:    ast.NodeKind_AttrString,
		ptrnStr: pattern,
	}
}

func Int() *attr {
	return &attr{
		kind:   ast.NodeKind_AttrInteger,
		minInt: math.MinInt64,
		maxInt: math.MaxInt64,
	}
}

func IntExact(value int64) *attr {
	return &attr{
		kind:   ast.NodeKind_AttrInteger,
		minInt: value,
		maxInt: value,
	}
}

func IntRange(min, max int64) *attr {
	return &attr{
		kind:   ast.NodeKind_AttrInteger,
		minInt: min,
		maxInt: max,
	}
}

// Sequence matches seq in exact provided order
func Sequence(seq ...Matcher) *sequence {
	return &sequence{
		seq: seq,
	}
}

// Optional can match m
func Optional(m Matcher) *optional {
	return &optional{
		submatcher: m,
	}
}

// OneOf matches one of variants
func OneOf(variants ...Matcher) *oneOf {
	return &oneOf{
		variants: variants,
	}
}

// All matches all variants in any order
func All(variants ...Matcher) *matchAll {
	return &matchAll{
		submatchers: variants,
	}
}

// Repeat matches m any number of time
func Repeat(m Matcher) *repeat {
	return &repeat{
		matcher:    m,
		minRepeats: 0,
		maxRepeats: math.MaxInt,
	}
}

// RepeatExact matches m exact n times
func RepeatExact(n int, m Matcher) *repeat {
	return &repeat{
		matcher:    m,
		minRepeats: n,
		maxRepeats: n,
	}
}

// RepeatAtLeast matches m at least minRepeats times
func RepeatAtLeast(minRepeats int, m Matcher) *repeat {
	return &repeat{
		matcher:    m,
		minRepeats: minRepeats,
		maxRepeats: math.MaxInt,
	}
}

// RepeatRange matches m at least minRepeats times
// but no more that maxRepeats
func RepeatRange(minRepeats, maxRepeats int, m Matcher) *repeat {
	return &repeat{
		matcher:    m,
		minRepeats: minRepeats,
		maxRepeats: maxRepeats,
	}
}

type Matcher interface {
	TryMatch(c *Cursor) (any, bool)
}

type MatcherBase struct{}

type subnodeMatcher struct {
	MatcherBase
	exactName bool
	name      string
	ptrn      *regexp.Regexp
	inner     Matcher
}

type attr struct {
	MatcherBase
	kind     ast.NodeKind
	exactStr string
	ptrnStr  *regexp.Regexp
	minInt   int64
	maxInt   int64
}

type sequence struct {
	MatcherBase
	seq []Matcher
}

type optional struct {
	MatcherBase
	submatcher Matcher
}

type oneOf struct {
	MatcherBase
	variants []Matcher
}

type matchAll struct {
	MatcherBase
	submatchers []Matcher
}

type repeat struct {
	MatcherBase
	matcher    Matcher
	minRepeats int
	maxRepeats int
}

type transform struct {
	f func(any) any
}

func (m *MatcherBase) Transform(transformer func(in any) any) *transform {
	return &transform{
		f: transformer,
	}
}

func (m *subnodeMatcher) TryMatch(c *Cursor) (any, bool) {
	node := c.Node()

	if node.Kind != ast.NodeKind_SubNode {
		return nil, false
	}

	if m.exactName && node.Name != m.name {
		return nil, false
	} else if m.ptrn != nil && !m.ptrn.MatchString(node.Name) {
		return nil, false
	}

	cp := c.Mark()
	c.EnterChildren()
	var res any

	if m.inner != nil {
		var ok bool
		if res, ok = m.inner.TryMatch(c); !ok {
			c.Reset(cp)
			return nil, false
		}
	}

	c.ExitChildren()
	c.GotoNextSibling()

	return res, true
}

func (a *attr) TryMatch(c *Cursor) (any, bool) {
	node := c.Node()

	if node.Kind != a.kind {
		return nil, false
	}

	switch node.Kind {
	case ast.NodeKind_AttrString:
		fallthrough
	case ast.NodeKind_AttrIdentifier:
		fallthrough
	case ast.NodeKind_AttrKeyword:
		if (a.ptrnStr != nil && a.ptrnStr.MatchString(node.StrValue)) ||
			(len(a.exactStr) > 0 && a.exactStr == node.StrValue) ||
			(len(a.exactStr) == 0) {
			return node, true
		}

	case ast.NodeKind_AttrInteger:
		if node.IntValue >= a.minInt && node.IntValue <= a.maxInt {
			return node, true
		}
	}

	return nil, false
}

func (s *sequence) TryMatch(c *Cursor) (any, bool) {
	cp := c.Mark()
	ress := []any{}

	for i, m := range s.seq {
		if res, ok := m.TryMatch(c); ok {
			ress = append(ress, res)
			if !c.GotoNextSibling() && i < len(s.seq)-1 {
				c.Reset(cp)
				return nil, false
			}
		} else {
			c.Reset(cp)
			return nil, false
		}
	}

	return ress, true
}

func (o *optional) TryMatch(c *Cursor) (any, bool) {
	if res, ok := o.submatcher.TryMatch(c); ok {
		return res, true
	}

	return nil, true
}

func (a *oneOf) TryMatch(c *Cursor) (any, bool) {
	cp := c.Mark()

	for _, m := range a.variants {
		if res, ok := m.TryMatch(c); ok {
			return res, true
		}

		c.Reset(cp)
	}

	return nil, false
}

func (a *matchAll) TryMatch(c *Cursor) (any, bool) {
	cp := c.Mark()
	ress := []any{}

outter:
	for len(ress) < len(a.submatchers) {
		for _, m := range a.submatchers {
			if res, ok := m.TryMatch(c); ok {
				ress = append(ress, res)
				c.GotoNextSibling()
				continue outter
			}
		}

		c.Reset(cp)
		return nil, false
	}

	return ress, true
}

func (r *repeat) TryMatch(c *Cursor) (any, bool) {
	cp := c.Mark()

	var ress []any

	for range r.maxRepeats {
		res, ok := r.matcher.TryMatch(c)

		if !ok {
			break
		}

		ress = append(ress, res)

		if len(ress) > r.maxRepeats {
			break
		}
	}

	if len(ress) < r.minRepeats {
		c.Reset(cp)
		return nil, false
	}

	return ress, true
}

func (t *transform) TryMatch(c *Cursor) (any, bool) {
	return t.f(c.Node()), true
}
