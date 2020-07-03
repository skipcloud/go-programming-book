package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// encode writes to buf an S-expression representation of v.
func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem())

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				writeWhiteSpace(buf)
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(')')

	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			var str string
			if i > 0 {
				writeWhiteSpace(buf)
			}
			name := v.Type().Field(i).Name
			fmt.Fprintf(buf, "(%s ", name)
			if needsIndent(v.Field(i).Kind()) {
				str = generateIndent(name)
				updateIndentation(str, add)
			}
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			buf.WriteByte(')')
			updateIndentation(str, remove)
		}
		buf.WriteByte(')')

	case reflect.Map: // ((key value) ...)
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(' ')
			}
			buf.WriteByte('(')
			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	case reflect.Bool: // t or nil
		if v.Bool() {
			buf.WriteString("t")
		} else {
			buf.WriteString("nil")
		}

	case reflect.Complex64, reflect.Complex128: // #C(num num)
		c := v.Complex()
		buf.WriteString("#C(")
		fmt.Fprintf(buf, "%.2f %.2f)", real(c), imag(c))

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%g", v.Float())

	case reflect.Interface:
		var str string

		buf.WriteByte('(')
		name := v.Type().String()
		fmt.Fprintf(buf, "%q ", name)
		str = generateIndent(name)
		updateIndentation(str, add)

		if err := encode(buf, v.Elem()); err != nil {
			return err
		}
		buf.WriteByte(')')
		updateIndentation(str, remove)

	default: // chan, func
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

type indentAction string

const (
	add    indentAction = "add"
	remove indentAction = "remove"
)

var indentation = ""

// updateIndentation will add or remove "str" amount of spaces to/from the top level
// indentation variable
func updateIndentation(str string, action indentAction) {
	switch action {
	case remove:
		indentation = strings.TrimSuffix(indentation, str)
	case add:
		indentation += str
	}
}

func generateIndent(fieldName string) string {
	// add indent for first paren, space, and second paren
	s := "   "
	// include whitespace for struct name
	for i := 0; i < len(fieldName); i++ {
		s += " "
	}
	return s
}

func needsIndent(k reflect.Kind) bool {
	return k == reflect.Struct ||
		k == reflect.Slice ||
		k == reflect.Array ||
		k == reflect.Interface
}
func writeWhiteSpace(b *bytes.Buffer) {
	b.WriteByte('\n')
	if indentation == "" {
		// first new line so indent enough for a single parenthesis
		indentation += " "
	}
	b.WriteString(indentation)
}
