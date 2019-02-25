package test

import (
	"reflect"
	"testing"
)

func AssertEquals(t *testing.T, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Equals assertion failed. Expected: %v. Actual: %v", expected, actual)
	}
}

func AssertNil(t *testing.T, actual interface{}) {
	if actual != nil {
		t.Errorf("Nil assertion failed. Actual: %v", actual)
	}
}

func AssertNotNil(t *testing.T, actual interface{}) {
	if actual == nil {
		t.Errorf("Not nil assertion failed. Actual: %v", actual)
	}
}
