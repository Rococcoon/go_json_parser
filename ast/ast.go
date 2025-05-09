package ast

import (
	"fmt"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Value interface {
	Node
	valueNode()
}

// root node
type RootNode struct {
	Value Value
}

func (r *RootNode) TokenLiteral() string {
	if r.Value != nil {
		return r.Value.TokenLiteral()
	}
	return ""
}

func (r *RootNode) String() string {
	if r.Value != nil {
		return r.Value.TokenLiteral()
	}
	return ""
}

// Property
type Property struct {
	Key   string
	Value Value
}

func (p *Property) TokenLiteral() string {
	return p.Key
}

func (p *Property) String() string {
	return fmt.Sprintf(`"%s": %s`, p.Key, p.Value.String())
}

// Object
type Object struct {
	Pairs []Property
}

func (o *Object) TokenLiteral() string {
	if len(o.Pairs) > 0 {
		return o.Pairs[0].TokenLiteral()
	}
	return "{}"
}

func (o *Object) String() string {
	var pairs []string

	for _, pair := range o.Pairs {
		pairs = append(pairs, pair.String())
	}

	return "{" + strings.Join(pairs, ", ") + "}"
}

func (o *Object) valueNode() {}

// Array
type Array struct {
	Elements []Value
}

func (a *Array) TokenLiteral() string {
	if len(a.Elements) > 0 {
		return a.Elements[0].TokenLiteral()
	}
	return "[]"
}

func (a *Array) String() string {
	var elements []string

	for _, elem := range a.Elements {
		elements = append(elements, elem.String())
	}

	return "[" + strings.Join(elements, ", ") + "]"
}

func (a *Array) valueNode() {}

// String literal
type StringLiteral struct {
	Value string
}

func (s *StringLiteral) TokenLiteral() string {
	return s.Value
}

func (s *StringLiteral) String() string {
	return `"` + s.Value + `"`
}

func (s *StringLiteral) valueNode() {}

// Number literal
type NumberLiteral struct {
	Value float64
}

func (n *NumberLiteral) TokenLiteral() string {
	return fmt.Sprintf("%g", n.Value)
}

func (n *NumberLiteral) String() string {
	return fmt.Sprintf("%f", n.Value)
}

func (n *NumberLiteral) valueNode() {}

// boolean literal
type BooleanLiteral struct {
	Value bool
}

func (b *BooleanLiteral) TokenLiteral() string {
	return fmt.Sprintf("%t", b.Value)
}

func (b *BooleanLiteral) String() string {
	return fmt.Sprintf("%t", b.Value)
}

func (b *BooleanLiteral) valueNode() {}

// Null Literal
type NullLiteral struct{}

func (n *NullLiteral) TokenLiteral() string {
	return "null"
}

func (n *NullLiteral) String() string {
	return "null"
}

func (n *NullLiteral) valueNode() {}

// Illegal Literal
type IllegalLiteral struct {
	Message string
}

func (i *IllegalLiteral) TokenLiteral() string {
	return "Illegal"
}

func (i *IllegalLiteral) String() string {
	if i.Message != "" {
		return "Illegal" + i.Message
	}
	return "Illegal"
}

func (i *IllegalLiteral) valueNode() {}
