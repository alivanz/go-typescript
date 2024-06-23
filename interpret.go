package typescript

import (
	"reflect"
	"strings"
	"time"
)

var (
	date = reflect.TypeOf(time.Time{})
)

func Interpret(v interface{}) *Type {
	return interpret(reflect.TypeOf(v))
}

func interpret(rt reflect.Type) *Type {
	if rt == nil {
		return TypeNull
	}
	if rt.Kind() == reflect.Pointer {
		rt = rt.Elem()
		for rt.Kind() == reflect.Pointer {
			rt = rt.Elem()
		}
		return &Type{
			Kind:  KindComposite,
			Types: []*Type{TypeNull, interpret(rt)},
		}
	}
	switch kind := rt.Kind(); kind {
	default:
		return nil
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		fallthrough
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		fallthrough
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		return TypeNumber
	case reflect.Bool:
		return TypeBoolean
	case reflect.String:
		return TypeString
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		return &Type{
			Kind:  KindArray,
			Inner: interpret(rt.Elem()),
		}
	case reflect.Struct:
		if rt == date {
			return TypeDate
		}
		n := rt.NumField()
		e := make([]Element, 0, n)
		for i := 0; i < n; i++ {
			f := rt.Field(i)
			name := f.Name
			switch v := strings.Split(f.Tag.Get("json"), ",")[0]; v {
			default:
				name = v
			case "-":
				continue
			case "":
				// nothing
			}
			e = append(e, Element{
				Name: name,
				Type: interpret(f.Type),
			})
		}
		return &Type{
			Name:    rt.Name(),
			Kind:    KindStruct,
			Element: e,
		}
	}
}
