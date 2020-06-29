package mapset

import (
	"bytes"
	"fmt"
)

// MapIntSet is my built-in map implementation for IntSet, this will be
// used to test against the IntSet implementation.
type MapIntSet map[uint64]struct{}

func (m MapIntSet) Has(x int) bool {
	_, ok := m[uint64(x)]
	return ok
}

func (m MapIntSet) Add(x int) {
	m[uint64(x)] = struct{}{}
}

func (m MapIntSet) UnionWith(mi MapIntSet) {
	for k := range mi {
		if _, ok := m[k]; !ok {
			m[k] = struct{}{}
		}
	}
}

func (m MapIntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for k := range m {
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", k)
	}
	buf.WriteByte('}')
	return buf.String()
}
