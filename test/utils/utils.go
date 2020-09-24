package utils

import (
	"fmt"
	"reflect"
)

//FindInSlice item in the slice
func FindInSlice(slice interface{}, item interface{}) (index int, found bool) {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Lookup is not a slice. %T provided", s))
	}
	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return i, true
		}
	}
	return -1, false
}
