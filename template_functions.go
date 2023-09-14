package view

import (
	"reflect"
)

// InGroupsOf splits a slice into groups of size n.
func InGroupsOf(s any, n int) [][]any {
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Slice {
		panic("unabled to split non-slice")
	}

	final := make([][]any, val.Len()/n+1)

	for i := 0; i < val.Len(); i += n {
		for j := 0; j < n; j++ {
			if i+j < val.Len() {
				final[i/n] = append(final[i/n], val.Index(i+j).Interface())
			}
		}
	}

	return final
}
