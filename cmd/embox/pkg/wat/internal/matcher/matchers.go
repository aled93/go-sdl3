package matcher

import (
	"math"
	"regexp"
	"sdl3/cmd/embox/pkg/wat/internal/ast"
	"sdl3/cmd/embox/pkg/wat/internal/tuple"
)

// Subnode matches any subnode
func Subnode[T any](inner Matcher[T]) *subnodeMatcher[T] {
	return &subnodeMatcher[T]{
		inner: inner,
	}
}

// NamedSubnode matches subnode with exact name
func NamedSubnode[T any](name string, inner Matcher[T]) *subnodeMatcher[T] {
	return &subnodeMatcher[T]{
		inner:     inner,
		exactName: true,
		name:      name,
	}
}

// NamePatternSubnode matches subnode testing name against regexp pattern
func NamePatternSubnode[T any](namePattern *regexp.Regexp, inner Matcher[T]) *subnodeMatcher[T] {
	if namePattern == nil {
		panic("namePattern is nil")
	}

	return &subnodeMatcher[T]{
		inner: inner,
		ptrn:  namePattern,
	}
}

// Keyword matches any keyword attribute
func Keyword() *attr {
	return &attr{
		kind: ast.NodeKind_AttrKeyword,
	}
}

// Keyword matches keyword attribute with exact value
func KeywordExact(keyword string) *attr {
	return &attr{
		kind:     ast.NodeKind_AttrKeyword,
		exactStr: keyword,
	}
}

// KeywordPattern matches keyword attribute with value that satisfies regexp pattern
func KeywordPattern(pattern *regexp.Regexp) *attr {
	if pattern == nil {
		panic("pattern is nil")
	}

	return &attr{
		kind:    ast.NodeKind_AttrKeyword,
		ptrnStr: pattern,
	}
}

// Identifier matches any identifier attribute
func Identifier() *attr {
	return &attr{
		kind: ast.NodeKind_AttrIdentifier,
	}
}

// IdentifierExact matches identifier attribute with exact value
func IdentifierExact(ident string) *attr {
	return &attr{
		kind:     ast.NodeKind_AttrIdentifier,
		exactStr: ident,
	}
}

// IdentifierPattern matches identifier attribute with value that satisfies
// regexp pattern
func IdentifierPattern(pattern *regexp.Regexp) *attr {
	if pattern == nil {
		panic("pattern is nil")
	}

	return &attr{
		kind:    ast.NodeKind_AttrIdentifier,
		ptrnStr: pattern,
	}
}

// String matches any string attribute
func String(ident string) *attr {
	return &attr{
		kind: ast.NodeKind_AttrString,
	}
}

// StringExact matches string attribute with exact value
func StringExact(ident string) *attr {
	return &attr{
		kind:     ast.NodeKind_AttrString,
		exactStr: ident,
	}
}

// StringPattern matches string attribute with value satisfies regexp pattern
func StringPattern(pattern *regexp.Regexp) *attr {
	if pattern == nil {
		panic("pattern is nil")
	}

	return &attr{
		kind:    ast.NodeKind_AttrString,
		ptrnStr: pattern,
	}
}

// Int matches any integer attribute
func Int() *attr {
	return &attr{
		kind:   ast.NodeKind_AttrInteger,
		minInt: math.MinInt64,
		maxInt: math.MaxInt64,
	}
}

// IntExact matches integer attribute with exact value
func IntExact(value int64) *attr {
	return &attr{
		kind:   ast.NodeKind_AttrInteger,
		minInt: value,
		maxInt: value,
	}
}

// IntRange matches integer attribute with value within range
func IntRange(min, max int64) *attr {
	return &attr{
		kind:   ast.NodeKind_AttrInteger,
		minInt: min,
		maxInt: max,
	}
}

// Sequence2 matches sequences of 2 submatchers in exact provided order
func Sequence2[T1, T2 any](m1 Matcher[T1], m2 Matcher[T2]) *sequence2[T1, T2] {
	return &sequence2[T1, T2]{
		m1: m1,
		m2: m2,
	}
}

// Sequence3 matches sequences of 3 submatchers in exact provided order
func Sequence3[T1, T2, T3 any](m1 Matcher[T1], m2 Matcher[T2], m3 Matcher[T3]) *sequence3[T1, T2, T3] {
	return &sequence3[T1, T2, T3]{
		m1: m1,
		m2: m2,
		m3: m3,
	}
}

// Sequence4 matches sequences of 4 submatchers in exact provided order
func Sequence4[T1, T2, T3, T4 any](m1 Matcher[T1], m2 Matcher[T2], m3 Matcher[T3], m4 Matcher[T4]) *sequence4[T1, T2, T3, T4] {
	return &sequence4[T1, T2, T3, T4]{
		m1: m1,
		m2: m2,
		m3: m3,
		m4: m4,
	}
}

// Optional can match m
func Optional[T any](m Matcher[T]) *optional[T] {
	return &optional[T]{
		submatcher: m,
	}
}

// OneOfSame matches one of variants with same output type
func OneOfSame[T any](variants ...Matcher[T]) *oneOfSame[T] {
	return &oneOfSame[T]{
		variants: variants,
	}
}

// OneOf2 matches one of 2 variants
func OneOf2[T1, T2 any](v1 Matcher[T1], v2 Matcher[T2]) *oneOf2[T1, T2] {
	return &oneOf2[T1, T2]{
		v1: v1,
		v2: v2,
	}
}

// OneOf3 matches one of 3 variants
func OneOf3[T1, T2, T3 any](v1 Matcher[T1], v2 Matcher[T2], v3 Matcher[T3]) *oneOf3[T1, T2, T3] {
	return &oneOf3[T1, T2, T3]{
		v1: v1,
		v2: v2,
		v3: v3,
	}
}

// OneOf4 matches one of 4 variants
func OneOf4[T1, T2, T3, T4 any](v1 Matcher[T1], v2 Matcher[T2], v3 Matcher[T3], v4 Matcher[T4]) *oneOf4[T1, T2, T3, T4] {
	return &oneOf4[T1, T2, T3, T4]{
		v1: v1,
		v2: v2,
		v3: v3,
		v4: v4,
	}
}

// All2 matches all variants in any order
func All2[T1, T2 any](m1 Matcher[T1], m2 Matcher[T2]) *matchAll2[T1, T2] {
	return &matchAll2[T1, T2]{
		m1: m1,
		m2: m2,
	}
}

// All3 matches all variants in any order
func All3[T1, T2, T3 any](m1 Matcher[T1], m2 Matcher[T2], m3 Matcher[T3]) *matchAll3[T1, T2, T3] {
	return &matchAll3[T1, T2, T3]{
		m1: m1,
		m2: m2,
		m3: m3,
	}
}

// All4 matches all variants in any order
func All4[T1, T2, T3, T4 any](m1 Matcher[T1], m2 Matcher[T2], m3 Matcher[T3], m4 Matcher[T4]) *matchAll4[T1, T2, T3, T4] {
	return &matchAll4[T1, T2, T3, T4]{
		m1: m1,
		m2: m2,
		m3: m3,
		m4: m4,
	}
}

// Repeat matches m any number of time
func Repeat[T any](m Matcher[T]) *repeat[T] {
	return &repeat[T]{
		matcher:    m,
		minRepeats: 0,
		maxRepeats: math.MaxInt,
	}
}

// RepeatExact matches m exact n times
func RepeatExact[T any](n int, m Matcher[T]) *repeat[T] {
	return &repeat[T]{
		matcher:    m,
		minRepeats: n,
		maxRepeats: n,
	}
}

// RepeatAtLeast matches m at least minRepeats times
func RepeatAtLeast[T any](minRepeats int, m Matcher[T]) *repeat[T] {
	return &repeat[T]{
		matcher:    m,
		minRepeats: minRepeats,
		maxRepeats: math.MaxInt,
	}
}

// RepeatRange matches m at least minRepeats times
// but no more that maxRepeats
func RepeatRange[T any](minRepeats, maxRepeats int, m Matcher[T]) *repeat[T] {
	return &repeat[T]{
		matcher:    m,
		minRepeats: minRepeats,
		maxRepeats: maxRepeats,
	}
}

// Transform transforms output of matcher m using provided function
func Transform[In, Out any](m Matcher[In], transformer func(in In) Out) *transform[In, Out] {
	return &transform[In, Out]{
		m: m,
		f: transformer,
	}
}

type Matcher[Out any] interface {
	TryMatch(c *Cursor) (Out, bool)
}

type MatcherBase struct{}

type subnodeMatcher[T any] struct {
	MatcherBase
	exactName bool
	name      string
	ptrn      *regexp.Regexp
	inner     Matcher[T]
}

type attr struct {
	MatcherBase
	kind     ast.NodeKind
	exactStr string
	ptrnStr  *regexp.Regexp
	minInt   int64
	maxInt   int64
}

type sequence2[T1, T2 any] struct {
	MatcherBase
	m1 Matcher[T1]
	m2 Matcher[T2]
}

type sequence3[T1, T2, T3 any] struct {
	MatcherBase
	m1 Matcher[T1]
	m2 Matcher[T2]
	m3 Matcher[T3]
}

type sequence4[T1, T2, T3, T4 any] struct {
	MatcherBase
	m1 Matcher[T1]
	m2 Matcher[T2]
	m3 Matcher[T3]
	m4 Matcher[T4]
}

type optional[T any] struct {
	MatcherBase
	submatcher Matcher[T]
}

type oneOfSame[T any] struct {
	MatcherBase
	variants []Matcher[T]
}

type oneOf2[T1, T2 any] struct {
	MatcherBase
	v1 Matcher[T1]
	v2 Matcher[T2]
}

type oneOf3[T1, T2, T3 any] struct {
	MatcherBase
	v1 Matcher[T1]
	v2 Matcher[T2]
	v3 Matcher[T3]
}

type oneOf4[T1, T2, T3, T4 any] struct {
	MatcherBase
	v1 Matcher[T1]
	v2 Matcher[T2]
	v3 Matcher[T3]
	v4 Matcher[T4]
}

type matchAll2[T1, T2 any] struct {
	MatcherBase
	m1 Matcher[T1]
	m2 Matcher[T2]
}

type matchAll3[T1, T2, T3 any] struct {
	MatcherBase
	m1 Matcher[T1]
	m2 Matcher[T2]
	m3 Matcher[T3]
}

type matchAll4[T1, T2, T3, T4 any] struct {
	MatcherBase
	m1 Matcher[T1]
	m2 Matcher[T2]
	m3 Matcher[T3]
	m4 Matcher[T4]
}

type repeat[T any] struct {
	MatcherBase
	matcher    Matcher[T]
	minRepeats int
	maxRepeats int
}

type transform[In, Out any] struct {
	m Matcher[In]
	f func(In) Out
}

func (m *subnodeMatcher[T]) TryMatch(c *Cursor) (res T, ok bool) {
	node := c.Node()

	if node.Kind != ast.NodeKind_SubNode {
		return res, false
	}

	if m.exactName && node.Name != m.name {
		return res, false
	} else if m.ptrn != nil && !m.ptrn.MatchString(node.Name) {
		return res, false
	}

	cp := c.Mark()
	c.EnterChildren()

	if m.inner != nil {
		if innerRes, ok := m.inner.TryMatch(c); !ok {
			c.Reset(cp)
			return res, false
		} else {
			res = innerRes
		}
	}

	c.ExitChildren()
	c.GotoNextSibling()

	return res, true
}

func (a *attr) TryMatch(c *Cursor) (*ast.Node, bool) {
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

func (s *sequence2[T1, T2]) TryMatch(c *Cursor) (res tuple.Of2[T1, T2], ok bool) {
	cp := c.Mark()

	if m1Res, ok := s.m1.TryMatch(c); ok {
		res.M1 = m1Res
		if !c.GotoNextSibling() {
			c.Reset(cp)
			return res, false
		}
	} else {
		c.Reset(cp)
		return res, false
	}

	if m2Res, ok := s.m2.TryMatch(c); ok {
		res.M2 = m2Res
		c.GotoNextSibling()
	} else {
		c.Reset(cp)
		return res, false
	}

	return res, true
}

func (s *sequence3[T1, T2, T3]) TryMatch(c *Cursor) (res tuple.Of3[T1, T2, T3], ok bool) {
	cp := c.Mark()

	if m1Res, ok := s.m1.TryMatch(c); ok {
		res.M1 = m1Res
		if !c.GotoNextSibling() {
			c.Reset(cp)
			return res, false
		}
	} else {
		c.Reset(cp)
		return res, false
	}

	if m2Res, ok := s.m2.TryMatch(c); ok {
		res.M2 = m2Res
		if !c.GotoNextSibling() {
			c.Reset(cp)
			return res, false
		}
	} else {
		c.Reset(cp)
		return res, false
	}

	if m3Res, ok := s.m3.TryMatch(c); ok {
		res.M3 = m3Res
		c.GotoNextSibling()
	} else {
		c.Reset(cp)
		return res, false
	}

	return res, true
}

func (s *sequence4[T1, T2, T3, T4]) TryMatch(c *Cursor) (res tuple.Of4[T1, T2, T3, T4], ok bool) {
	cp := c.Mark()

	if m1Res, ok := s.m1.TryMatch(c); ok {
		res.M1 = m1Res
		if !c.GotoNextSibling() {
			c.Reset(cp)
			return res, false
		}
	} else {
		c.Reset(cp)
		return res, false
	}

	if m2Res, ok := s.m2.TryMatch(c); ok {
		res.M2 = m2Res
		if !c.GotoNextSibling() {
			c.Reset(cp)
			return res, false
		}
	} else {
		c.Reset(cp)
		return res, false
	}

	if m3Res, ok := s.m3.TryMatch(c); ok {
		res.M3 = m3Res
		if !c.GotoNextSibling() {
			c.Reset(cp)
			return res, false
		}
	} else {
		c.Reset(cp)
		return res, false
	}

	if m4Res, ok := s.m4.TryMatch(c); ok {
		res.M4 = m4Res
		c.GotoNextSibling()
	} else {
		c.Reset(cp)
		return res, false
	}

	return res, true
}

func (o *optional[T]) TryMatch(c *Cursor) (T, bool) {
	if res, ok := o.submatcher.TryMatch(c); ok {
		return res, true
	}

	var zeroT T
	return zeroT, true
}

func (a *oneOfSame[T]) TryMatch(c *Cursor) (res tuple.Of2[T, int], ok bool) {
	cp := c.Mark()

	for i, v := range a.variants {
		if mRes, ok := v.TryMatch(c); ok {
			return tuple.New2(mRes, i), true
		}
	}

	c.Reset(cp)

	return res, false
}

func (a *oneOf2[T1, T2]) TryMatch(c *Cursor) (res tuple.Of3[T1, T2, int], ok bool) {
	cp := c.Mark()

	if mRes, ok := a.v1.TryMatch(c); ok {
		return tuple.New3(mRes, *new(T2), 1), true
	}

	if mRes, ok := a.v2.TryMatch(c); ok {
		return tuple.New3(*new(T1), mRes, 2), true
	}

	c.Reset(cp)

	return res, false
}

func (a *oneOf3[T1, T2, T3]) TryMatch(c *Cursor) (res tuple.Of4[T1, T2, T3, int], ok bool) {
	cp := c.Mark()

	if mRes, ok := a.v1.TryMatch(c); ok {
		return tuple.New4(mRes, *new(T2), *new(T3), 1), true
	}

	if mRes, ok := a.v2.TryMatch(c); ok {
		return tuple.New4(*new(T1), mRes, *new(T3), 2), true
	}

	if mRes, ok := a.v3.TryMatch(c); ok {
		return tuple.New4(*new(T1), *new(T2), mRes, 3), true
	}

	c.Reset(cp)

	return res, false
}

func (a *oneOf4[T1, T2, T3, T4]) TryMatch(c *Cursor) (res tuple.Of5[T1, T2, T3, T4, int], ok bool) {
	cp := c.Mark()

	if mRes, ok := a.v1.TryMatch(c); ok {
		return tuple.New5(mRes, *new(T2), *new(T3), *new(T4), 1), true
	}

	if mRes, ok := a.v2.TryMatch(c); ok {
		return tuple.New5(*new(T1), mRes, *new(T3), *new(T4), 2), true
	}

	if mRes, ok := a.v3.TryMatch(c); ok {
		return tuple.New5(*new(T1), *new(T2), mRes, *new(T4), 3), true
	}

	if mRes, ok := a.v4.TryMatch(c); ok {
		return tuple.New5(*new(T1), *new(T2), *new(T3), mRes, 4), true
	}

	c.Reset(cp)

	return res, false
}

func (a *matchAll2[T1, T2]) TryMatch(c *Cursor) (res tuple.Of2[T1, T2], ok bool) {
	cp := c.Mark()

	for i := range res.Size() {
		isLast := i == res.Size()-1

		if mRes, ok := a.m1.TryMatch(c); ok {
			res.M1 = mRes
			if !c.GotoNextSibling() && !isLast {
				c.Reset(cp)
				return res, false
			}
			continue
		}

		if mRes, ok := a.m2.TryMatch(c); ok {
			res.M2 = mRes
			if !c.GotoNextSibling() && !isLast {
				c.Reset(cp)
				return res, false
			}
			continue
		}

		c.Reset(cp)
		return res, false
	}

	return res, true
}

func (a *matchAll3[T1, T2, T3]) TryMatch(c *Cursor) (res tuple.Of3[T1, T2, T3], ok bool) {
	cp := c.Mark()

	for i := range res.Size() {
		isLast := i == res.Size()-1

		if mRes, ok := a.m1.TryMatch(c); ok {
			res.M1 = mRes
			if !c.GotoNextSibling() && !isLast {
				c.Reset(cp)
				return res, false
			}
			continue
		}

		if mRes, ok := a.m2.TryMatch(c); ok {
			res.M2 = mRes
			if !c.GotoNextSibling() && !isLast {
				c.Reset(cp)
				return res, false
			}
			continue
		}

		if mRes, ok := a.m3.TryMatch(c); ok {
			res.M3 = mRes
			if !c.GotoNextSibling() && !isLast {
				c.Reset(cp)
				return res, false
			}
			continue
		}

		c.Reset(cp)
		return res, false
	}

	return res, true
}

func (a *matchAll4[T1, T2, T3, T4]) TryMatch(c *Cursor) (res tuple.Of4[T1, T2, T3, T4], ok bool) {
	cp := c.Mark()

	for i := range res.Size() {
		isLast := i == res.Size()-1

		if mRes, ok := a.m1.TryMatch(c); ok {
			res.M1 = mRes
			if !c.GotoNextSibling() && !isLast {
				c.Reset(cp)
				return res, false
			}
			continue
		}

		if mRes, ok := a.m2.TryMatch(c); ok {
			res.M2 = mRes
			if !c.GotoNextSibling() && !isLast {
				c.Reset(cp)
				return res, false
			}
			continue
		}

		if mRes, ok := a.m3.TryMatch(c); ok {
			res.M3 = mRes
			if !c.GotoNextSibling() && !isLast {
				c.Reset(cp)
				return res, false
			}
			continue
		}

		if mRes, ok := a.m4.TryMatch(c); ok {
			res.M4 = mRes
			if !c.GotoNextSibling() && !isLast {
				c.Reset(cp)
				return res, false
			}
			continue
		}

		c.Reset(cp)
		return res, false
	}

	return res, true
}

func (r *repeat[T]) TryMatch(c *Cursor) (res []T, ok bool) {
	cp := c.Mark()

	for range r.maxRepeats {
		mRes, ok := r.matcher.TryMatch(c)

		if !ok {
			break
		}

		res = append(res, mRes)

		if len(res) > r.maxRepeats {
			break
		}
	}

	if len(res) < r.minRepeats {
		c.Reset(cp)
		return nil, false
	}

	return res, true
}

func (t *transform[In, Out]) TryMatch(c *Cursor) (res Out, ok bool) {
	if mRes, ok := t.m.TryMatch(c); !ok {
		return res, false
	} else {
		return t.f(mRes), true
	}
}
