package typescript

import (
	"strings"
)

type Type struct {
	Name    string    `json:""`
	Kind    Kind      `json:""`
	Inner   *Type     `json:",omitempty"`
	Element []Element `json:",omitempty"`
	Types   []*Type   `json:",omitempty"`
}

type Element struct {
	Name string
	Type *Type
}

type Kind int

const (
	KindNative Kind = iota
	KindArray
	KindStruct
	KindComposite
)

func (t *Type) RName() string {
	switch t.Kind {
	default:
		return t.Name
	case KindArray:
		return t.Inner.RName() + "[]"
	case KindComposite:
		var buf strings.Builder
		buf.WriteByte('(')
		for i, t := range t.Types {
			if i != 0 {
				buf.WriteString(" | ")
			}
			buf.WriteString(t.RName())
		}
		buf.WriteByte(')')
		return buf.String()
	}
}
