package typescript

import (
	"strings"
	"testing"
)

func TestTranspileStruct(t *testing.T) {
	var (
		buf strings.Builder
	)
	type Account struct {
		ID       string `gorm:"primaryKey"`
		Email    string `gorm:"not null;unique"`
		Active   bool   `gorm:"not null"`
		Password string `gorm:"not null" json:"-"`
	}
	Transpile(&buf, Interpret(Account{}))
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
	type Account struct {
		ID       string `gorm:"primaryKey"`
		Email    string `gorm:"not null;unique"`
		Active   bool   `gorm:"not null"`
		Password string `gorm:"not null" json:"-"`
	}
	type Group struct {
		ID        string     `gorm:"primaryKey"`
		Name      string     `gorm:"not null"`
		Member    []Account  `gorm:"foreignKey:AccountID"`
		MemberPtr []*Account `gorm:"foreignKey:AccountID"`
	}
	Transpile(&buf, Interpret(Group{}))
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
