package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
	Write a function that reports whether its argument is a
	cyclic data structure.
*/

type mystruct struct {
	num  int
	name string
	st   *mystruct
}

func main() {
	x := struct{}{}
	fmt.Println(isCyclic(&x)) // should be false

	y := mystruct{
		num:  1,
		name: "Sally",
	}
	fmt.Println(isCyclic(&y)) // should be false
	fmt.Println(isCyclic(y))  // should be false

	y.st = &y
	fmt.Println(isCyclic(&y)) // should be true
}

func isCyclic(arg interface{}) bool {
	v := reflect.ValueOf(arg)

	// if pointer then dereference to get the element
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// if we can't get an address or it isn't a struct then
	// it cannot be cyclic
	if !v.CanAddr() || v.Kind() != reflect.Struct {
		return false
	}
	objPtr := unsafe.Pointer(v.UnsafeAddr())

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Kind() != reflect.Ptr || f.IsZero() {
			continue
		}
		f = f.Elem()
		fPtr := unsafe.Pointer(f.UnsafeAddr())
		if objPtr == fPtr {
			return true
		}
	}
	return false
}
