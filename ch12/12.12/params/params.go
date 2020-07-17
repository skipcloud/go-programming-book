// Package params provides a reflection-based parser for URL parameters.
package params

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type field struct {
	val reflect.Value
	con string
}

// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	fields := make(map[string]*field)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		condition := tag.Get("validate")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = &field{
			val: v.Field(i),
			con: condition,
		}
	}

	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		f := fields[name]
		if !f.val.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		for _, value := range values {
			if err := validate(f, value); err != nil {
				return err
			}

			if f.val.Kind() == reflect.Slice {
				elem := reflect.New(f.val.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.val.Set(reflect.Append(f.val, elem))
			} else {
				if err := populate(f.val, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}

func validate(f *field, value string) error {
	// no condition set then nothing needs validated
	if f.con == "" {
		return nil
	}

	// this switch deals with two cases but obviously it could be
	// extended to work with a whole load more validation conditions
	switch f.con {
	case "email":
		// contains an @? Let's pretend that's valid
		if !strings.ContainsRune(value, '@') {
			return errors.New("not a valid email")
		}
	case "zip":
		// assume normal five-digit ZIP code
		_, err := strconv.ParseInt(value, 10, 0)
		if len(value) != 5 || err != nil {
			return errors.New("not a valid zip code")
		}
	default:
		return errors.New("unknown validation condition")
	}
	return nil
}
