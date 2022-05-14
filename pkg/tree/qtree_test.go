package tree

import (
	"fmt"
	"testing"
)

func compare(a, b []int) int {
	quad := 0

	for i := 0; i < len(a); i++ {
		if b[i] >= a[i] {
			quad += 0b1 << i
		}
	}

	return quad
}

func TestNewQTree(t *testing.T) {
	const dim = 3
	tree := NewQTree(dim, compare)

	if tree.root != nil {
		t.Errorf("Actual root node = %v, Expected == nil", tree.root)
	}
	if tree.dim != dim {
		t.Errorf("Actual tree dim = %d, Expected == %d", tree.dim, dim)
	}
}

func TestQTree_NaiveInsertOnce(t *testing.T) {
	const dim = 3
	tree := NewQTree[[]int](dim, compare)

	root := NewQNode(make([]int, dim), dim)
	tree.NaiveInsert(root)

	if &tree.root.item == &root.item {
		t.Errorf("Actual root node = %v, Expected == %v", tree.root, root)
	}
	if tree.dim != dim {
		t.Errorf("Actual tree dim = %d, Expected == %d", tree.dim, dim)
	}
}

func TestQTree_NaiveInsert(t *testing.T) {
	const dim = 3
	tree := NewQTree[[]int](dim, compare)

	root := NewQNode(make([]int, dim), dim)
	tree.NaiveInsert(root)

	for x := -dim; x <= dim; x++ {
		for y := -dim; y <= dim; y++ {
			for z := -dim; z <= dim; z++ {
				if x == 0 && y == 0 && z == 0 {
					continue
				}

				item := []int{x, y, z}
				node := NewQNode(item, dim)

				tree.NaiveInsert(node)
			}
		}
	}

	fmt.Printf("%v", tree.root)

	if &tree.root.item == &root.item {
		t.Errorf("Actual root node = %v, Expected == %v", tree.root, root)
	}
	if tree.dim != dim {
		t.Errorf("Actual tree dim = %d, Expected == %d", tree.dim, dim)
	}
}
