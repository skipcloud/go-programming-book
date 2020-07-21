package main

import "reflect"

/*
	Define a deep comparison function that considers numbers (of any type) equal
	if they differ less than one part in a billion

	edit: not sure if they mean a function just to compare numbers or compare
	      any value. I'll just write a function to compare numbers which could
		  be used in a lager comparison function.
*/

func main() {

}

func deepCompare(a, b interface{}) bool {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)
	if !av.IsValid() || !bv.IsValid() {
		return av.IsValid() == bv.IsValid()
	}
	return true
}

// let's use the short-scale billion
const billion = 1_000_000_000

func valueToInt(v reflect.Value) int64 {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u := v.Uint()
	default:
		panic("unknown type")
	}
}
