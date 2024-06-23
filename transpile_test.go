package typescript

import (
	"strings"
	"testing"
)

func TestTranspileStruct(t *testing.T) {
	var (
		buf strings.Builder
	)
	Transpile(&buf, &Type{
		Name: "Account",
		Kind: KindStruct,
		Element: []Element{
			{Name: "ID", Type: TypeString},
			{Name: "Email", Type: TypeString},
			{Name: "Active", Type: TypeBoolean},
		},
	})
	if s := buf.String(); s != `
type Account = {
	ID: string
	Email: string
	Active: boolean
}
` {
		t.Fatal(s)
	}
}

func TestTranspileArray(t *testing.T) {
	var (
		buf strings.Builder
	)
	account := &Type{
		Name: "Account",
		Kind: KindStruct,
		Element: []Element{
			{Name: "ID", Type: TypeString},
			{Name: "Email", Type: TypeString},
			{Name: "Active", Type: TypeBoolean},
		},
	}
	Transpile(&buf, &Type{
		Name: "Group",
		Kind: KindStruct,
		Element: []Element{
			{Name: "ID", Type: TypeString},
			{Name: "Name", Type: TypeString},
			{Name: "Member", Type: &Type{
				Kind:  KindArray,
				Inner: account,
			}},
			{Name: "MemberPtr", Type: &Type{
				Kind: KindArray,
				Inner: &Type{
					Kind:  KindComposite,
					Types: []*Type{TypeNull, account},
				},
			}},
		},
	})
	if s := buf.String(); s != `
type Group = {
	ID: string
	Name: string
	Member: Account[]
	MemberPtr: (null | Account)[]
}
` {
		t.Fatal(s)
	}
}

func TestTranspileAnonymous(t *testing.T) {
	var (
		buf strings.Builder
	)
	ref := &Type{
		Name: "Parent",
		Kind: KindStruct,
		Element: []Element{
			{Name: "Int", Type: TypeNumber},
			{Name: "Nested", Type: &Type{
				Kind: KindStruct,
				Element: []Element{
					{Name: "Index", Type: TypeNumber},
					{Name: "Name", Type: TypeString},
				},
			}},
		},
	}
	Transpile(&buf, ref)
	if s := buf.String(); s != `
type Parent = {
	Int: number
	Nested: {
		Index: number
		Name: string
	}
}
` {
		t.Fatal(s)
	}
}
