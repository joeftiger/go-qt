package tree

import (
	"fmt"
	"testing"
)

func TestRangeSearch_DIM2(t *testing.T) {

	const dim = 2
	tree := NewQTree[[]int](dim, CompareOrdered[int])

	//insert 5 items
	var items [5][]int
	items[0] = []int{-9, -9}
	items[1] = []int{-9, -11}
	items[2] = []int{-9, 2}
	items[3] = []int{-10, -9}
	items[4] = []int{12, -1}

	for i := 0; i < 5; i++ {
		tree.NaiveInsert(items[i])
	}

	tl := []int{-10, 10}
	tr := []int{10, 10}

	bl := []int{-10, -10}
	br := []int{10, -10}

	bound := [][]int{tr, tl, bl, br}

	for _, f := range tree.RangeSearch(bound) {
		fmt.Printf("%v\n", f.item)
	}

}

func TestRangeSearch_DIM3(t *testing.T) {

	const dim = 3
	tree := NewQTree[[]int](dim, CompareOrdered[int])

	//insert 5 items
	var items [5][]int
	items[0] = []int{-9, -9, -9}
	items[1] = []int{-9, -11, -9}
	items[2] = []int{-9, 2, -9}
	items[3] = []int{-10, -9, -9}
	items[4] = []int{12, -1, -9}

	for i := 0; i < 5; i++ {
		tree.NaiveInsert(items[i])
	}

	bound := [][]int{
		[]int{-10, -10, -10},
		[]int{10, 10, 10},
		[]int{10, 10, -10},
		[]int{-10, 10, 10},
		[]int{10, -10, 10},
		[]int{-10, 10, -10},
		[]int{-10, -10, 10},
		[]int{10, -10, -10},
	}

	fmt.Println("Found items in range search")
	for _, f := range tree.RangeSearch(bound) {
		fmt.Printf("%v\n", f.item)
	}

}
