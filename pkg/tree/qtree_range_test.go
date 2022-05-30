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

	for i, f := range tree.RangeSearch(bound) {
		fmt.Printf("[%d]: %v\n", i, f.item)
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
	for i, f := range tree.RangeSearch(bound) {
		fmt.Printf("[%d]: %v\n", i, f.item)
	}
}

func TestRangeSearch_DIM3_KV(t *testing.T) {
	const dim = 2
	tree := NewQTree[KVEntryString](dim, compareKvStr)

	// First entry
	key1 := []string{"Switzerland", "Fribourg"}
	val1 := "Hello Fribourg"
	entry1 := KVEntryString{&key1, &val1}
	tree.NaiveInsert(entry1)

	// Second entry
	key2 := []string{"Switzerland", "Bern"}
	val2 := "Hello Bern"
	entry2 := KVEntryString{&key2, &val2}
	tree.NaiveInsert(entry2)

	// Third entry
	key3 := []string{"Swaziland", "Mbabane"}
	val3 := "Hello Mbabane "
	entry3 := KVEntryString{&key3, &val3}
	tree.NaiveInsert(entry3)

	bound := [][]string{
		[]string{"Swa", "C"},
		[]string{"Saw", "Z"},
		[]string{"Z", "A"},
		[]string{"Z", "Z"},
	}

	bound_kv := []KVEntryString{
		KVEntryString{&bound[0], nil},
		KVEntryString{&bound[1], nil},
		KVEntryString{&bound[2], nil},
		KVEntryString{&bound[3], nil},
	}

	fmt.Println("Found items in range search for KV String Store")

	//Should return
	// [0]: key: [Switzerland Fribourg] val: [Switzerland Fribourg]
	// [1]: key: [Swaziland Mbabane] val: [Swaziland Mbabane]
	for i, f := range tree.RangeSearch(bound_kv) {
		fmt.Printf("[%d]: key: %v ", i, *f.item.key)
		fmt.Printf("val: %v\n", *f.item.key)
	}

}
