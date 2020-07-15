package sexpr

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

type Encoder struct {
	out io.Writer
}

func (e *Encoder) Encode(v interface{}) error {
	return encode(e.out, reflect.ValueOf(v))
}

// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// encode writes to buf an S-expression representation of v.
func encode(w io.Writer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		w.Write([]byte("nil"))

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(w, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(w, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(w, "%q", v.String())

	case reflect.Ptr:
		return encode(w, v.Elem())

	case reflect.Array, reflect.Slice: // (value ...)
		w.Write([]byte{'('})
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				w.Write([]byte{' '})
			}
			if err := encode(w, v.Index(i)); err != nil {
				return err
			}
		}
		w.Write([]byte{')'})

	case reflect.Struct: // ((name value) ...)
		w.Write([]byte{'('})
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				w.Write([]byte{' '})
			}
			fmt.Fprintf(w, "(%s ", v.Type().Field(i).Name)
			if err := encode(w, v.Field(i)); err != nil {
				return err
			}
			w.Write([]byte{')'})
		}
		w.Write([]byte{')'})

	case reflect.Map: // ((key value) ...)
		w.Write([]byte{'('})
		for i, key := range v.MapKeys() {
			if i > 0 {
				w.Write([]byte{' '})
			}
			w.Write([]byte{'('})
			if err := encode(w, key); err != nil {
				return err
			}
			w.Write([]byte{' '})
			if err := encode(w, v.MapIndex(key)); err != nil {
				return err
			}
			w.Write([]byte{')'})
		}
		w.Write([]byte{')'})

	default: // float, complex, bool, chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
