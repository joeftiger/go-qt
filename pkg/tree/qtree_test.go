package tree

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestNewQTree(t *testing.T) {
	const dim = 3
	tree := NewQTree(dim, CompareOrdered[int])

	if tree.root != nil {
		t.Errorf("Actual root node = %v, Expected == nil", tree.root)
	}
}

func TestQTree_NaiveInsertOnce(t *testing.T) {
	const dim = 3
	tree := NewQTree[[]int](dim, CompareOrdered[int])

	item := make([]int, dim)

	root := NewQNode(item, dim)
	tree.NaiveInsertNode(root)

	if tree.root == &root {
		t.Errorf("Actual root node = %v, Expected == %v", tree.root, root)
	}

	found := tree.PointSearch(item)
	if !reflect.DeepEqual(found, &root) {
		t.Errorf("Actual found = %v, Expected == %v", found, root)
	}
}

func TestQTree_NaiveInsertMany(t *testing.T) {
	const dim = 3
	tree := NewQTree[[]int](dim, CompareOrdered[int])

	root := NewQNode(make([]int, dim), dim)
	tree.NaiveInsertNode(root)

	for x := -dim; x <= dim; x++ {
		for y := -dim; y <= dim; y++ {
			for z := -dim; z <= dim; z++ {
				if x == 0 && y == 0 && z == 0 {
					continue
				}

				item := []int{x, y, z}
				tree.NaiveInsert(item)
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

func FuzzQTree_NaiveInsert(f *testing.F) {
	const dim = 16
	tree := NewQTree[[]int](dim, CompareOrdered[int])

	root := NewQNode(make([]int, dim), dim)
	tree.NaiveInsertNode(root)

	f.Fuzz(func(t *testing.T, unusedButNeeded int) {
		item := make([]int, dim)
		for i := 0; i < dim; i++ {
			item[i] = rand.Int()
		}

		tree.NaiveInsert(item)

		found := tree.PointSearch(item)
		if found == nil {
			t.Errorf("%v: Actual found = nil, Expected == not nil", item)
		}
	})
}
