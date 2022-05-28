package tree

import (
	"reflect"
	"testing"
)

// For testing purposes

// AssertEqual checks if values are equal.
// (https://gist.github.com/samalba/6059502)
func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
	}
}

// AssertNotEqual checks if values are not equal.
func AssertNotEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		t.Errorf("Received %v (type %v) for both, expected different values", a, reflect.TypeOf(a))
	}
}

// AssertEqualDeep checks if values are deeply equal.
// See reflect.DeepEqual for more information.
func AssertEqualDeep(t *testing.T, a interface{}, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
	}
}
