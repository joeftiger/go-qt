package tree

import (
	"testing"
)

// Test float
func TestFloat(t *testing.T) {

	const dim = 3
	tree := NewQTree[[]float32](compare_ordered[float32])

	item := []float32{1.434, 21.222, 332.23432}
	fst := NewQNode(item, dim)
	tree.NaiveInsert(fst)

	// Search
	found := tree.PointSearch(item)
	AssertEqualDeep(t, found.item, item)

	not_found := tree.PointSearch([]float32{1.4, 21.2, 332.2}) // doesn't exist
	var nil_QNode *QNode[[]float32]
	AssertEqual(t, not_found, nil_QNode)
}
