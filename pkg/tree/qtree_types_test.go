package tree

import (
	"testing"
)

// Test float
func TestFloat(t *testing.T) {
	const dim = 3
	tree := NewQTree[[]float32](dim, CompareOrdered[float32])

	item := []float32{1.434, 21.222, 332.23432}
	tree.NaiveInsert(item)

	// Search
	found := tree.PointSearch(item)
	AssertEqualDeep(t, found.item, item)

	notFound := tree.PointSearch([]float32{1.4, 21.2, 332.2}) // doesn't exist
	var nilQnode *QNode[[]float32]
	AssertEqual(t, notFound, nilQnode)
}
