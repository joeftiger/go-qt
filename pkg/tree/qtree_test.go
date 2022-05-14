package tree

import (
	"reflect"
	"testing"
)

func compare(a, b []int) int {
	quad := 0

	if reflect.DeepEqual(a, b) {
		return quad
	}

	for i := 0; i < len(a); i++ {
		if b[i] >= a[i] {
			quad += 0b1 << i
		}
	}

	return quad
}

func TestNewQTree(t *testing.T) {
	const dim = 3
	tree := NewQTree(compare)

	if tree.root != nil {
		t.Errorf("Actual root node = %v, Expected == nil", tree.root)
	}
}

func TestQTree_NaiveInsertOnce(t *testing.T) {
	const dim = 3
	tree := NewQTree[[]int](compare)

	item := make([]int, dim)

	root := NewQNode(item, dim)
	tree.NaiveInsert(root)

	if tree.root == &root {
		t.Errorf("Actual root node = %v, Expected == %v", tree.root, root)
	}

	found := tree.PointSearch(item)
	if !reflect.DeepEqual(found, &root) {
		t.Errorf("Actual found = %v, Expected == %v", found, root)
	}
}

func TestQTree_NaiveInsert(t *testing.T) {
	const dim = 3
	tree := NewQTree[[]int](compare)

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

	for x := -dim; x <= dim; x++ {
		for y := -dim; y <= dim; y++ {
			for z := -dim; z <= dim; z++ {
				if x == 0 && y == 0 && z == 0 {
					continue
				}

				item := []int{x, y, z}
				found := tree.PointSearch(item)

				if found == nil {
					t.Errorf("%v: Actual found = nil, Expected == not nil", item)
				}
			}
		}
	}

	if &tree.root.item == &root.item {
		t.Errorf("Actual root node = %v, Expected == %v", tree.root, root)
	}
}
