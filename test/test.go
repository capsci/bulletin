package test

import (
	"fmt"
	"reflect"
	"testing"
)

// TypeMismatchErrorMsg message when comparing variables of differnt kind or type
const TypeMismatchErrorMsg = "KindMismatch: Cannot compare type '%T' to type '%T'"

// ValueNotEqualMsg message when the values of variables are not equal
const ValueNotEqualMsg = "\nActual : %v\nExpected : %v\n%s"

// Expected method for testing
func Expected(t *testing.T, calculated interface{}, expected interface{}, message string) {
	// TODO: Print caller line
	if reflect.TypeOf(calculated).String() != reflect.TypeOf(expected).String() {
		t.Error(fmt.Sprintf(TypeMismatchErrorMsg, calculated, expected))
		return
	}
	if calculated != expected {
		t.Error(fmt.Sprintf(ValueNotEqualMsg, calculated, expected, message))
	}
}
