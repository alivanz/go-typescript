package typescript

import (
	"fmt"
	"io"
)

func Transpile(w io.Writer, ts ...*Type) {
	for _, t := range ts {
		transpile(w, t, "")
	}
}

func transpile(w io.Writer, t *Type, indent string) {
	switch t.Kind {
	case KindNative:
		// nothing todo
	case KindArray:
	case KindStruct:
		fmt.Fprintf(w, "\ntype %s = ", t.Name)
		transpileStruct(w, t, indent)
	}
}

func transpileStruct(w io.Writer, t *Type, indent string) {
	fmt.Fprintf(w, "{\n")
	for _, e := range t.Element {
		tname := e.Type.RName()
		if tname == "" {
			fmt.Fprintf(w, "%s\t%s: ", indent, e.Name)
			transpileStruct(w, e.Type, indent+"\t")
		} else {
			fmt.Fprintf(w, "%s\t%s: %s\n", indent, e.Name, tname)
		}
	}
	fmt.Fprintf(w, "%s}\n", indent)
}
