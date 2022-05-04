package tree

import (
	"math"
	"testing"
)

func TestQuadrantFrom(t *testing.T) {
	for i := byte(0); i < math.MaxUint8; i++ {
		success := testQuadrantFromSuccess(i)

		if IsValidQuadrant(i) {
			if !success {
				t.Errorf("%d: Actual quadrant = panic, Expected == success", i)
			}
		} else if success {
			t.Errorf("%d: Actual quadrant = success, Expected panic", i)
		}
	}
}

// testQuadrantFromSuccess tests QuadrantFrom for success or a panic.
func testQuadrantFromSuccess(dir byte) bool {
	defer func() bool {
		if r := recover(); r != nil {
			return false
		} else {
			return true
		}
	}()

	QuadrantFrom(dir)

	return true
}

func TestQuadrantTryFrom(t *testing.T) {
	for i := byte(0); i < math.MaxUint8; i++ {
		q, err := QuadrantTryFrom(i)

		if IsValidQuadrant(i) {
			if q == nil {
				t.Errorf("Actual quadrant = %d, Expected == %v", q, nil)
			}
			if err != nil {
				t.Errorf("Actual error = %v, Expected == %v", err, nil)
			}
		} else {
			if q != nil {
				t.Errorf("Actual quadrant = %d, Expected == %v", q, "not nil")
			}
			if err == nil {
				t.Errorf("Actual error = %v, Expected == %v", nil, "invalid direction")
			}
		}
	}
}
