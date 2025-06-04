package matcher

import (
	"fmt"
	"math"
	"regexp"
	"sdl3/cmd/embox/pkg/wat/internal/ast"
	"sdl3/cmd/embox/pkg/wat/internal/tuple"
	"strings"
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
	TryMatch(c *Cursor) (Out, error)
	expects() string
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

func (m *subnodeMatcher[T]) TryMatch(c *Cursor) (res T, err error) {
	node := c.Node()

	if node == nil {
		return res, newSyntaxError(c, m)
	}

	if node.Kind != ast.NodeKind_SubNode {
		return res, newSyntaxError(c, m)
	}

	if m.exactName && node.Name != m.name {
		return res, newSyntaxError(c, m)
	} else if m.ptrn != nil && !m.ptrn.MatchString(node.Name) {
		return res, newSyntaxError(c, m)
	}

	cp := c.Mark()
	c.EnterChildren()

	if m.inner != nil {
		if innerRes, err := m.inner.TryMatch(c); err != nil {
			c.Reset(cp)
			return res, err
		} else {
			res = innerRes
		}
	}

	// TODO: test this only in strict parser mode
	if !c.EndOfSubnode() {
		return res, newSyntaxErrorCustom(c, fmt.Errorf("expected end of subnode, but extra node found (%s)", c.Node()))
	}

	c.ExitChildren()
	c.GotoNextSibling()

	return res, nil
}

func (m *subnodeMatcher[T]) expects() string {
	if m.exactName {
		return fmt.Sprintf(`subnode with name "%s"`, m.name)
	} else if m.ptrn != nil {
		return fmt.Sprintf(`subnode with name matching pattern "%s"`, m.ptrn)
	} else {
		return `subnode`
	}
}

func (a *attr) TryMatch(c *Cursor) (*ast.Node, error) {
	node := c.Node()

	if node == nil {
		return nil, newSyntaxError(c, a)
	}

	if node.Kind != a.kind {
		return nil, newSyntaxError(c, a)
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
			c.GotoNextSibling()
			return node, nil
		}

	case ast.NodeKind_AttrInteger:
		if node.IntValue >= a.minInt && node.IntValue <= a.maxInt {
			c.GotoNextSibling()
			return node, nil
		}
	}

	return nil, newSyntaxErrorCustom(c, fmt.Errorf("unexpected attribute type %s, expected %s", node.Kind, a.kind))
}

func (a *attr) expects() string {
	switch a.kind {
	case ast.NodeKind_AttrFloat:
		return "float attribute"

	case ast.NodeKind_AttrIdentifier:
		if len(a.exactStr) > 0 {
			return fmt.Sprintf(`identifier attribute with value "$%s"`, a.exactStr)
		} else if a.ptrnStr != nil {
			return fmt.Sprintf(`identifier attribute with value matching regexp "%s"`, a.ptrnStr)
		} else {
			return `identifier attribute`
		}

	case ast.NodeKind_AttrInteger:
		if a.minInt == 0 && a.maxInt == math.MaxInt64 {
			return `integer attribute`
		} else if a.minInt == a.maxInt {
			return fmt.Sprintf(`integer attribute with value %d`, a.minInt)
		} else if a.minInt > 0 && a.maxInt == math.MaxInt64 {
			return fmt.Sprintf(`integer attribute with value >= %d`, a.minInt)
		} else if a.maxInt < math.MaxInt64 && a.minInt == 0 {
			return fmt.Sprintf(`integer attribute with value <= %d`, a.minInt)
		} else {
			return fmt.Sprintf(`integer attribute with value within range [%d..%d]`, a.minInt, a.maxInt)
		}

	case ast.NodeKind_AttrKeyword:
		if len(a.exactStr) > 0 {
			return fmt.Sprintf(`keyword "%s"`, a.exactStr)
		} else if a.ptrnStr != nil {
			return fmt.Sprintf(`keyword matching regexp "%s"`, a.ptrnStr)
		} else {
			return `keyword`
		}

	case ast.NodeKind_AttrString:
		if len(a.exactStr) > 0 {
			return fmt.Sprintf(`string attribute with value "%s"`, a.exactStr)
		} else if a.ptrnStr != nil {
			return fmt.Sprintf(`string attribute with value matching regexp "%s"`, a.ptrnStr)
		} else {
			return `string attribute`
		}
	}

	panic("unhandled attribute kind")
}

func (s *sequence2[T1, T2]) TryMatch(c *Cursor) (res tuple.Of2[T1, T2], err error) {
	if c.EndOfSubnode() {
		return res, newSyntaxError(c, s)
	}

	cp := c.Mark()

	if m1Res, err := s.m1.TryMatch(c); err == nil {
		res.M1 = m1Res
		if c.EndOfSubnode() {
			c.Reset(cp)
			return res, newSyntaxErrorCustom(c, fmt.Errorf("unexpected end of subnode, expected 1 more element"))
		}
	} else {
		c.Reset(cp)
		return res, err
	}

	if m2Res, err := s.m2.TryMatch(c); err == nil {
		res.M2 = m2Res
	} else {
		c.Reset(cp)
		return res, err
	}

	return res, nil
}

func (s *sequence2[T1, T2]) expects() string {
	return fmt.Sprintf(`%s followed by %s`, s.m1.expects(), s.m2.expects())
}

func (s *sequence3[T1, T2, T3]) TryMatch(c *Cursor) (res tuple.Of3[T1, T2, T3], err error) {
	if c.EndOfSubnode() {
		return res, newSyntaxError(c, s)
	}

	cp := c.Mark()

	if m1Res, err := s.m1.TryMatch(c); err == nil {
		res.M1 = m1Res
		if c.EndOfSubnode() {
			c.Reset(cp)
			return res, newSyntaxErrorCustom(c, fmt.Errorf("unexpected end of subnode, expected 2 more elements"))
		}
	} else {
		c.Reset(cp)
		return res, err
	}

	if m2Res, err := s.m2.TryMatch(c); err == nil {
		res.M2 = m2Res
		if c.EndOfSubnode() {
			c.Reset(cp)
			return res, newSyntaxErrorCustom(c, fmt.Errorf("unexpected end of subnode, expected 1 more element"))
		}
	} else {
		c.Reset(cp)
		return res, err
	}

	if m3Res, err := s.m3.TryMatch(c); err == nil {
		res.M3 = m3Res
	} else {
		c.Reset(cp)
		return res, err
	}

	return res, nil
}

func (s *sequence3[T1, T2, T3]) expects() string {
	return fmt.Sprintf(`sequence of %s, %s and %s`, s.m1.expects(), s.m2.expects(), s.m3.expects())
}

func (s *sequence4[T1, T2, T3, T4]) TryMatch(c *Cursor) (res tuple.Of4[T1, T2, T3, T4], err error) {
	if c.EndOfSubnode() {
		return res, newSyntaxError(c, s)
	}

	cp := c.Mark()

	if m1Res, err := s.m1.TryMatch(c); err == nil {
		res.M1 = m1Res
		if c.EndOfSubnode() {
			c.Reset(cp)
			return res, newSyntaxErrorCustom(c, fmt.Errorf("unexpected end of subnode, expected 3 more elements"))
		}
	} else {
		c.Reset(cp)
		return res, err
	}

	if m2Res, err := s.m2.TryMatch(c); err == nil {
		res.M2 = m2Res
		if c.EndOfSubnode() {
			c.Reset(cp)
			return res, newSyntaxErrorCustom(c, fmt.Errorf("unexpected end of subnode, expected 2 more elements"))
		}
	} else {
		c.Reset(cp)
		return res, err
	}

	if m3Res, err := s.m3.TryMatch(c); err == nil {
		res.M3 = m3Res
		if c.EndOfSubnode() {
			c.Reset(cp)
			return res, newSyntaxErrorCustom(c, fmt.Errorf("unexpected end of subnode, expected 1 more element"))
		}
	} else {
		c.Reset(cp)
		return res, err
	}

	if m4Res, err := s.m4.TryMatch(c); err == nil {
		res.M4 = m4Res
	} else {
		c.Reset(cp)
		return res, err
	}

	return res, nil
}

func (s *sequence4[T1, T2, T3, T4]) expects() string {
	return fmt.Sprintf(`sequence of %s, %s, %s and %s`, s.m1.expects(), s.m2.expects(), s.m3.expects(), s.m4.expects())
}

func (o *optional[T]) TryMatch(c *Cursor) (T, error) {
	if c.EndOfSubnode() {
		var zero T
		return zero, newSyntaxError(c, o)
	}

	if res, err := o.submatcher.TryMatch(c); err == nil {
		return res, nil
	}

	var zeroT T
	return zeroT, nil
}

func (o *optional[T]) expects() string {
	return fmt.Sprintf(`optional %s`, o.submatcher.expects())
}

func (a *oneOfSame[T]) TryMatch(c *Cursor) (res tuple.Of2[T, int], err error) {
	if c.EndOfSubnode() {
		return res, newSyntaxError(c, a)
	}

	cp := c.Mark()

	for i, v := range a.variants {
		if mRes, err := v.TryMatch(c); err == nil {
			return tuple.New2(mRes, i), nil
		}
	}

	c.Reset(cp)

	return res, newSyntaxError(c, a)
}

func (a oneOfSame[T]) expects() string {
	if len(a.variants) == 0 {
		return `nothing`
	} else if len(a.variants) == 1 {
		return a.variants[0].expects()
	} else if len(a.variants) == 2 {
		return a.variants[0].expects() + " or " + a.variants[1].expects()
	}

	var sb strings.Builder
	sb.WriteString("one of: ")

	for i, v := range a.variants {
		sb.WriteString(v.expects())

		if i < len(a.variants)-1 {
			sb.WriteString(", ")
		} else {
			sb.WriteString(" or ")
		}
	}

	return sb.String()
}

func (a *oneOf2[T1, T2]) TryMatch(c *Cursor) (res tuple.Of3[T1, T2, int], err error) {
	if c.EndOfSubnode() {
		return res, newSyntaxError(c, a)
	}

	cp := c.Mark()

	if mRes, err := a.v1.TryMatch(c); err == nil {
		return tuple.New3(mRes, *new(T2), 1), nil
	}

	if mRes, err := a.v2.TryMatch(c); err == nil {
		return tuple.New3(*new(T1), mRes, 2), nil
	}

	c.Reset(cp)

	return res, newSyntaxError(c, a)
}

func (a oneOf2[T1, T2]) expects() string {
	return a.v1.expects() + " or " + a.v2.expects()
}

func (a *oneOf3[T1, T2, T3]) TryMatch(c *Cursor) (res tuple.Of4[T1, T2, T3, int], err error) {
	if c.EndOfSubnode() {
		return res, newSyntaxError(c, a)
	}

	cp := c.Mark()

	if mRes, err := a.v1.TryMatch(c); err == nil {
		return tuple.New4(mRes, *new(T2), *new(T3), 1), nil
	}

	if mRes, err := a.v2.TryMatch(c); err == nil {
		return tuple.New4(*new(T1), mRes, *new(T3), 2), nil
	}

	if mRes, err := a.v3.TryMatch(c); err == nil {
		return tuple.New4(*new(T1), *new(T2), mRes, 3), nil
	}

	c.Reset(cp)

	return res, newSyntaxError(c, a)
}

func (a oneOf3[T1, T2, T3]) expects() string {
	return a.v1.expects() + ", " + a.v2.expects() + " or " + a.v3.expects()
}

func (a *oneOf4[T1, T2, T3, T4]) TryMatch(c *Cursor) (res tuple.Of5[T1, T2, T3, T4, int], err error) {
	if c.EndOfSubnode() {
		return res, newSyntaxError(c, a)
	}

	cp := c.Mark()

	if mRes, err := a.v1.TryMatch(c); err == nil {
		return tuple.New5(mRes, *new(T2), *new(T3), *new(T4), 1), nil
	}

	if mRes, err := a.v2.TryMatch(c); err == nil {
		return tuple.New5(*new(T1), mRes, *new(T3), *new(T4), 2), nil
	}

	if mRes, err := a.v3.TryMatch(c); err == nil {
		return tuple.New5(*new(T1), *new(T2), mRes, *new(T4), 3), nil
	}

	if mRes, err := a.v4.TryMatch(c); err == nil {
		return tuple.New5(*new(T1), *new(T2), *new(T3), mRes, 4), nil
	}

	c.Reset(cp)

	return res, newSyntaxError(c, a)
}

func (a oneOf4[T1, T2, T3, T4]) expects() string {
	return "one of: " +
		a.v1.expects() + ", " +
		a.v2.expects() + ", " +
		a.v3.expects() + " or " +
		a.v4.expects()
}

func (a *matchAll2[T1, T2]) TryMatch(c *Cursor) (res tuple.Of2[T1, T2], err error) {
	if c.EndOfSubnode() {
		return res, newSyntaxError(c, a)
	}

	cp := c.Mark()

	for i := range res.Size() {
		isLast := i == res.Size()-1

		if mRes, err := a.m1.TryMatch(c); err == nil {
			res.M1 = mRes
			if c.EndOfSubnode() && !isLast {
				c.Reset(cp)
				return res, newSyntaxErrorCustom(c, fmt.Errorf("unexpected end of subnode, expected %s", a.expects()))
			}
			continue
		}

		if mRes, err := a.m2.TryMatch(c); err == nil {
			res.M2 = mRes
			if c.EndOfSubnode() && !isLast {
				c.Reset(cp)
				return res, newSyntaxErrorCustom(c, fmt.Errorf("unexpected end of subnode, expected %s", a.expects()))
			}
			continue
		}

		c.Reset(cp)
		return res, newSyntaxError(c, a)
	}

	return res, nil
}

func (a *matchAll2[T1, T2]) expects() string {
	return a.m1.expects() + " or " + a.m2.expects()
}

func (a *matchAll3[T1, T2, T3]) TryMatch(c *Cursor) (res tuple.Of3[T1, T2, T3], err error) {
	if c.EndOfSubnode() {
		return res, newSyntaxError(c, a)
	}

	cp := c.Mark()

	for i := range res.Size() {
		isLast := i == res.Size()-1

		if mRes, err := a.m1.TryMatch(c); err == nil {
			res.M1 = mRes
			if c.EndOfSubnode() && !isLast {
				c.Reset(cp)
				return res, newSyntaxErrorCustom(c, fmt.Errorf("unexpected end of subnode, expected %s", a.expects()))
			}
			continue
		}

		if mRes, err := a.m2.TryMatch(c); err == nil {
			res.M2 = mRes
			if c.EndOfSubnode() && !isLast {
				c.Reset(cp)
				return res, newSyntaxErrorCustom(c, fmt.Errorf("unexpected end of subnode, expected %s", a.expects()))
			}
			continue
		}

		if mRes, err := a.m3.TryMatch(c); err == nil {
			res.M3 = mRes
			if c.EndOfSubnode() && !isLast {
				c.Reset(cp)
				return res, newSyntaxErrorCustom(c, fmt.Errorf("unexpected end of subnode, expected %s", a.expects()))
			}
			continue
		}

		c.Reset(cp)
		return res, newSyntaxError(c, a)
	}

	return res, nil
}

func (a *matchAll3[T1, T2, T3]) expects() string {
	return a.m1.expects() + ", " + a.m2.expects() + " or " + a.m3.expects()
}

func (a *matchAll4[T1, T2, T3, T4]) TryMatch(c *Cursor) (res tuple.Of4[T1, T2, T3, T4], err error) {
	if c.EndOfSubnode() {
		return res, newSyntaxError(c, a)
	}

	cp := c.Mark()

	for i := range res.Size() {
		isLast := i == res.Size()-1

		if mRes, err := a.m1.TryMatch(c); err == nil {
			res.M1 = mRes
			if c.EndOfSubnode() && !isLast {
				c.Reset(cp)
				return res, newSyntaxErrorCustom(c, fmt.Errorf("unexpected end of subnode, expected %s", a.expects()))
			}
			continue
		}

		if mRes, err := a.m2.TryMatch(c); err == nil {
			res.M2 = mRes
			if c.EndOfSubnode() && !isLast {
				c.Reset(cp)
				return res, newSyntaxErrorCustom(c, fmt.Errorf("unexpected end of subnode, expected %s", a.expects()))
			}
			continue
		}

		if mRes, err := a.m3.TryMatch(c); err == nil {
			res.M3 = mRes
			if c.EndOfSubnode() && !isLast {
				c.Reset(cp)
				return res, newSyntaxErrorCustom(c, fmt.Errorf("unexpected end of subnode, expected %s", a.expects()))
			}
			continue
		}

		if mRes, err := a.m4.TryMatch(c); err == nil {
			res.M4 = mRes
			if c.EndOfSubnode() && !isLast {
				c.Reset(cp)
				return res, newSyntaxErrorCustom(c, fmt.Errorf("unexpected end of subnode, expected %s", a.expects()))
			}
			continue
		}

		c.Reset(cp)
		return res, newSyntaxError(c, a)
	}

	return res, nil
}

func (a *matchAll4[T1, T2, T3, T4]) expects() string {
	return a.m1.expects() + ", " + a.m2.expects() + ", " + a.m3.expects() + " or " + a.m4.expects()
}

func (r *repeat[T]) TryMatch(c *Cursor) (res []T, err error) {
	if c.EndOfSubnode() {
		return res, newSyntaxError(c, r)
	}

	cp := c.Mark()

	for range r.maxRepeats {
		mRes, mErr := r.matcher.TryMatch(c)

		if mErr != nil {
			break
		}

		res = append(res, mRes)

		if len(res) > r.maxRepeats {
			break
		}
	}

	if len(res) < r.minRepeats {
		c.Reset(cp)
		return nil, newSyntaxError(c, r)
	}

	return res, nil
}

func (r *repeat[T]) expects() string {
	return r.matcher.expects()
}

func (t *transform[In, Out]) TryMatch(c *Cursor) (res Out, err error) {
	if mRes, err := t.m.TryMatch(c); err != nil {
		return res, err
	} else {
		return t.f(mRes), nil
	}
}

func (t *transform[In, Out]) expects() string {
	return t.m.expects()
}

type SyntaxError struct {
	Node *ast.Node
	Err  error
}

func (e *SyntaxError) Error() string {
	return e.Err.Error()
}

func newSyntaxError[T any](c *Cursor, matcher Matcher[T]) *SyntaxError {
	return &SyntaxError{
		Node: c.lastNonNilNode(),
		Err:  fmt.Errorf("expected %s, got %s", matcher.expects(), c.lastNonNilNode()),
	}
}

func newSyntaxErrorCustom(c *Cursor, err error) *SyntaxError {
	return &SyntaxError{
		Node: c.lastNonNilNode(),
		Err:  err,
	}
}
