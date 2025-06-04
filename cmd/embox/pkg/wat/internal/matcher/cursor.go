package matcher

import (
	"sdl3/cmd/embox/pkg/wat/internal/ast"
	"slices"
)

type Cursor struct {
	pos     *ast.Node
	parents []*ast.Node
}

type checkpoint struct {
	pos     *ast.Node
	parents []*ast.Node
}

func NewCursor(root *ast.Node) *Cursor {
	return &Cursor{
		pos: root,
	}
}

func (c *Cursor) Node() *ast.Node {
	return c.pos
}

func (c *Cursor) GotoNextSibling() bool {
	if c != nil && c.pos != nil {
		c.pos = c.pos.NextSibling
		return true
	}

	return false
}

func (c *Cursor) EnterChildren() bool {
	if c != nil && c.pos != nil && c.pos.FirstChild != nil {
		c.parents = append(c.parents, c.pos)
		c.pos = c.pos.FirstChild
		return true
	}

	return false
}

func (c *Cursor) ExitChildren() bool {
	if c == nil || len(c.parents) == 0 {
		return false
	}

	n := len(c.parents)
	c.pos = c.parents[n-1]
	c.parents = c.parents[:n-1]

	return true
}

func (c *Cursor) Mark() checkpoint {
	return checkpoint{
		pos:     c.pos,
		parents: slices.Clone(c.parents),
	}
}

func (c *Cursor) Reset(cp checkpoint) {
	c.pos = cp.pos
	c.parents = c.parents[:0]
	c.parents = append(c.parents, cp.parents...)
}

func (c *Cursor) EndOfSubnode() bool {
	return c.pos == nil
}
