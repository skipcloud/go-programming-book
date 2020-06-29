// Package display provides a means to display structured data.
package display

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

// formatAtom formats a value without inspecting its internal structure.
// It is a copy of the the function in gopl.io/ch11/format.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr,
		reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

const maxRecusionDepth = 3

var recursionDepth = 0

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			kind := v.Field(i).Kind()
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)

			// this isn't the most graceful solution but it works
			if kind == reflect.Ptr {
				if recursionDepth < maxRecusionDepth {
					recursionDepth++
				} else {
					recursionDepth = 0
					fmt.Printf("%s... = %s\n", path, formatAtom(v))
					continue
				}
			}
			display(fieldPath, v.Field(i))
		}
		recursionDepth = 0
	case reflect.Map:
		for _, key := range v.MapKeys() {
			var k string
			switch key.Kind() {
			case reflect.Struct:
				// let's print out the struct, as in "Name{field: value, field: value}"
				var buf bytes.Buffer

				buf.WriteString(key.Type().String())
				buf.WriteByte('{')

				for i := 0; i < key.NumField(); i++ {
					if i > 0 {
						buf.WriteString(", ")
					}
					buf.WriteString(key.Type().Field(i).Name + ": ")
					// to do this properly we would probably want to drill down
					// into the field values but I'm going to be lazy and just
					// format the value as an atom.
					buf.WriteString(formatAtom(key.Field(i)))
				}
				buf.WriteByte('}')
				k = buf.String()
			case reflect.Array:
				var buf bytes.Buffer

				buf.WriteString(key.Type().String())
				buf.WriteByte('[')
				for i := 0; i < key.Len(); i++ {
					if i > 0 {
						buf.WriteString(", ")
					}
					// like with structs above let's just formatAtom
					// the element
					buf.WriteString(formatAtom(key.Index(i)))
				}
				buf.WriteByte(']')
				k = buf.String()
			default:
				k = formatAtom(key)
			}
			display(fmt.Sprintf("%s[%s]", path, k), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}
