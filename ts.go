package typescript

import "fmt"

type Type struct {
	Name    string    `json:""`
	Kind    Kind      `json:""`
	Inner   *Type     `json:",omitempty"`
	Element []Element `json:",omitempty"`
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
	KindPointer
)

func (t *Type) RName() string {
	switch t.Kind {
	default:
		return t.Name
	case KindArray:
		return t.Inner.RName() + "[]"
	case KindPointer:
		return fmt.Sprintf("(undefined | %s)", t.Inner.RName())
	}
}
