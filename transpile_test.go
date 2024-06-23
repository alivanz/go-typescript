package typescript

import (
	"strings"
	"testing"
)

func TestTranspile(t *testing.T) {
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
	if s := buf.String(); s != `type Account = {
	ID: string
	Email: string
	Active: boolean
}
` {
		t.Fatal(s)
	}
	type Group struct {
		ID        string     `gorm:"primaryKey"`
		Name      string     `gorm:"not null"`
		Member    []Account  `gorm:"foreignKey:AccountID"`
		MemberPtr []*Account `gorm:"foreignKey:AccountID"`
	}
	buf.Reset()
	Transpile(&buf, Interpret(Group{}))
	if s := buf.String(); s != `type Group = {
	ID: string
	Name: string
	Member: Account[]
	MemberPtr: (null | Account)[]
}
` {
		t.Fatal(s)
	}
}
