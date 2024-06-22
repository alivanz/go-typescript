package typescript

import (
	"fmt"
	"io"
)

func Transpile(w io.Writer, ts ...*Type) {
	for _, t := range ts {
		transpile(w, t)
	}
}

func transpile(w io.Writer, t *Type) {
	switch t.Kind {
	case KindNative:
		// nothing todo
	case KindArray:
	case KindStruct:
		fmt.Fprintf(w, "type %s = {\n", t.Name)
		for _, e := range t.Element {
			fmt.Fprintf(w, "\t%s: %s\n", e.Name, e.Type.RName())
		}
		fmt.Fprintf(w, "}\n")
	}
}
