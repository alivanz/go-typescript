package typescript

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestInterpretPrimitive(t *testing.T) {
	if v := Interpret(true); !reflect.DeepEqual(v, TypeBoolean) {
		t.Fatal(v)
	}
	if v := Interpret(123); !reflect.DeepEqual(v, TypeNumber) {
		t.Fatal(v)
	}
	if v := Interpret(float32(1)); !reflect.DeepEqual(v, TypeNumber) {
		t.Fatal(v)
	}
	if v := Interpret(float64(3)); !reflect.DeepEqual(v, TypeNumber) {
		t.Fatal(v)
	}
	if v := Interpret(nil); !reflect.DeepEqual(v, TypeNull) {
		t.Fatal(v)
	}
	if v := Interpret(time.Now()); !reflect.DeepEqual(v, TypeDate) {
		t.Fatal(v)
	}
}

func TestInterpretStruct(t *testing.T) {
	type TestStruct struct {
		Name    string
		Alias   int `json:"alias"`
		Ignored int `json:"-"`
		Ptr1    *string
		Ptr2    **string
		Ptr3    ***string
		Array   [4]string
		Slice   []string
	}
	ref := &Type{
		Name: "TestStruct",
		Kind: KindStruct,
		Element: []Element{
			{Name: "Name", Type: TypeString},
			{Name: "alias", Type: TypeNumber},
			{Name: "Ptr1", Type: &Type{Kind: KindComposite, Types: []*Type{TypeNull, TypeString}}},
			{Name: "Ptr2", Type: &Type{Kind: KindComposite, Types: []*Type{TypeNull, TypeString}}},
			{Name: "Ptr3", Type: &Type{Kind: KindComposite, Types: []*Type{TypeNull, TypeString}}},
			{Name: "Array", Type: &Type{Kind: KindArray, Inner: TypeString}},
			{Name: "Slice", Type: &Type{Kind: KindArray, Inner: TypeString}},
		},
	}
	if v := Interpret(TestStruct{}); !reflect.DeepEqual(v, ref) {
		r, _ := json.MarshalIndent(ref, "", "\t")
		j, _ := json.MarshalIndent(v, "", "\t")
		t.Logf("%s", r)
		t.Fatalf("%s", j)
	}
}

func TestInterpretNested(t *testing.T) {
	type Child struct {
		Inner string
	}
	type Parent struct {
		Int    int
		Nested Child
		Array  []Child
	}
	ref := &Type{
		Name: "Parent",
		Kind: KindStruct,
		Element: []Element{
			{Name: "Int", Type: TypeNumber},
			{Name: "Nested", Type: &Type{
				Name: "Child",
				Kind: KindStruct,
				Element: []Element{
					{Name: "Inner", Type: TypeString},
				},
			}},
			{Name: "Array", Type: &Type{
				Kind: KindArray,
				Inner: &Type{
					Name: "Child",
					Kind: KindStruct,
					Element: []Element{
						{Name: "Inner", Type: TypeString},
					},
				},
			}},
		},
	}
	if v := Interpret(Parent{}); !reflect.DeepEqual(v, ref) {
		r, _ := json.MarshalIndent(ref, "", "\t")
		j, _ := json.MarshalIndent(v, "", "\t")
		t.Logf("%s", r)
		t.Fatalf("%s", j)
	}
}

func TestInterpretAnonymous(t *testing.T) {
	type Parent struct {
		Int    int
		Nested struct {
			Inner string
		}
	}
	ref := &Type{
		Name: "Parent",
		Kind: KindStruct,
		Element: []Element{
			{Name: "Int", Type: TypeNumber},
			{Name: "Nested", Type: &Type{
				Kind: KindStruct,
				Element: []Element{
					{Name: "Inner", Type: TypeString},
				},
			}},
		},
	}
	if v := Interpret(Parent{}); !reflect.DeepEqual(v, ref) {
		r, _ := json.MarshalIndent(ref, "", "\t")
		j, _ := json.MarshalIndent(v, "", "\t")
		t.Logf("%s", r)
		t.Fatalf("%s", j)
	}
}
