package main

import (
	"fmt"
	"math"
	"reflect"
)

/*
	Define a deep comparison function that considers numbers (of any type) equal
	if they differ less than one part in a billion

	edit: not sure if they mean a function just to compare numbers or to create
	      a comparison function that checks all types but treats number specially.
	      I'll just write a function to compare numbers which could be used in a
	      lager comparison function.
*/

func main() {
	x := -1
	y := uint(100)
	fmt.Println(deepCompare(x, y))

	a := -10.11111111121
	b := -10.11111111131
	fmt.Println(deepCompare(a, b))
}

func deepCompare(a, b interface{}) bool {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)
	if !av.IsValid() || !bv.IsValid() {
		return av.IsValid() == bv.IsValid()
	}
	if valueSign(av) != valueSign(bv) {
		return false
	}
	return valueToString(av) == valueToString(bv)
}

// valueSign returns a boolean indicating the value is positive
// or negative
func valueSign(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() >= 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	case reflect.Float32, reflect.Float64:
		// math.Signbit returns a boolean for the actual sign bit, which is
		// true for negative and false for positive.
		return !math.Signbit(v.Float())
	default:
		panic("unsupported type")
	}
}

// valueToString takes a reflect.Value and returns a string up to the 9th decimal place.
// Not sure if 9 decimal places is what is considered "one part in a billion" but this
// is what I'm running with.
//
// Also, I'm sure there is a way for me to compare values of different types by multiplying
// the number by a billion then converting to an int to drop the remainder of the decimal
// places, or something like that. I just opted for a string because it was simple.
func valueToString(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%.9f", float64(v.Int()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%.9f", float64(v.Uint()))
	case reflect.Float64, reflect.Float32:
		return fmt.Sprintf("%.9f", v.Float())
	default:
		panic("unsupported type")
	}
}
