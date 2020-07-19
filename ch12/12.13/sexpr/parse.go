package sexpr

import (
	"reflect"
	"strings"
)

func parseTags(st reflect.StructField) (name string, tags []string) {
	vv := strings.Split(st.Tag.Get("sexpr"), ",")
	name = vv[0]
	if name == "" {
		name = st.Name
	}
	tags = vv[1:]
	return
}

func tagsContains(tag string, tags []string) bool {
	for _, v := range tags {
		if v == tag {
			return true
		}
	}
	return false
}
