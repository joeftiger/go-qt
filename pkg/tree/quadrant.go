package tree

import "fmt"

/*
Quadrant is a directional quadrant from a point in a 2D universe.
Each quadrant is either north-east (Quadrant1), north-west (Quadrant2), south-west (Quadrant3) or south-east (Quadrant4).

The quadrants divide the space into open and closed spaces with Quadrant1 and Quadrant3 being closed.
*/
type Quadrant struct {
	// The direction of the quadrant. Should be inside [1, 4] to be IsValid.
	dir byte
}

// IsValid returns whether the quadrant is valid (direction in the interval [1, 4]).
func (q Quadrant) IsValid() bool {
	return IsValidQuadrant(q.dir)
}

// IsOpen returns whether the quadrant is open. (Quadrant2 or Quadrant4)
func (q Quadrant) IsOpen() bool {
	return q == Quadrant2() || q == Quadrant4() // more efficient than modulo
}

// IsClosed returns whether the quadrant is closed. (Quadrant1 or Quadrant3)
func (q Quadrant) IsClosed() bool {
	return q == Quadrant1() || q == Quadrant3() // more efficient than modulo
}

// Quadrant1 is the first / north-east quadrant.
func Quadrant1() Quadrant {
	return Quadrant{1}
}

// Quadrant2 is the second / north-west quadrant.
func Quadrant2() Quadrant {
	return Quadrant{2}
}

// Quadrant3 is the third / south-west quadrant.
func Quadrant3() Quadrant {
	return Quadrant{3}
}

// Quadrant4 is the fourth / south-west quadrant.
func Quadrant4() Quadrant {
	return Quadrant{4}
}

// IsValidQuadrant returns whether the given quadrant direction is valid (in the interval [1, 4]).
func IsValidQuadrant(dir byte) bool {
	return dir >= 1 && dir <= 4
}

// QuadrantFrom creates a Quadrant from an integer direction.
// This function panics when the given direction is invalid and outside the interval [1, 4].
func QuadrantFrom(dir byte) Quadrant {
	switch dir {
	case 1:
		return Quadrant1()
	case 2:
		return Quadrant2()
	case 3:
		return Quadrant3()
	case 4:
		return Quadrant4()
	default:
		err := fmt.Errorf("invalid quadrant direction: %d", dir)
		panic(err)
	}
}

// QuadrantTryFrom tries to create a Quadrant from an integer direction.
// This function returns an error when the given direction is invalid and outside the interval [1, 4].
func QuadrantTryFrom(dir byte) (*Quadrant, error) {
	if IsValidQuadrant(dir) {
		quad := QuadrantFrom(dir)
		return &quad, nil
	} else {
		return nil, fmt.Errorf("invalid quadrant direction: %d", dir)
	}
}
