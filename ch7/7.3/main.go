package main

import (
	"bytes"
	"fmt"
	"strconv"
)

/*
	Write a String method for the *tree type in gopl.io/ch4/treesort that
	reveals the sequence of values in the tree.
*/

func main() {

	t := &tree{}
	for _, val := range []int{1, 2, 3, -4, -5, 20, 10} {
		add(t, val)
	}
	fmt.Println(t)
}

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	var buf bytes.Buffer

	vals := walk(t)
	for i, val := range vals {
		buf.Write([]byte(strconv.FormatInt(int64(val), 10)))
		if i != len(vals) {
			buf.Write([]byte(" "))
		}
	}

	return buf.String()
}

func walk(t *tree) []int {
	s := []int{t.value}

	if t.left != nil {
		left := walk(t.left)
		s = append(left, s...)
	}

	if t.right != nil {
		s = append(s, walk(t.right)...)
	}
	return s
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
