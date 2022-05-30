package tree

import (
	"golang.org/x/exp/constraints"
	"reflect"
)

// CompareOrdered compares vectors with values that are part of the constraints.Ordered type, i.e. int, float, string.
func CompareOrdered[T constraints.Ordered](a, b []T) (equal bool, quad int) {
	if reflect.DeepEqual(a, b) {
		return true, -1
	}

	quad = 0
	for i := 0; i < len(a); i++ {
		if b[i] >= a[i] {
			quad += 0b1 << i
		}
	}

	return false, quad
}
