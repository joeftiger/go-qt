package tree

import (
	"testing"
)

// An example struct for a KV Store.
// We will use key for the comparison and value to store our value,
// in this case a string.
type KVEntry struct {
	key   []int
	value *string
}

// The compare function which basically relies on compare_int
func compare_kv_int(a, b KVEntry) (equal bool, quad int) {
	return compare_int(a.key, b.key)
}

// An example usage for the KV Store.
func TestKVStore(t *testing.T) {

	const dim = 3

	val := "hello world"
	key := []int{1, 2, 3}
	entry := KVEntry{key, &val}

	tree := NewQTree[KVEntry](compare_kv_int)
	fst := NewQNode(entry, dim)
	tree.NaiveInsert(fst)

	// Search
	empty_entry := KVEntry{[]int{1, 2, 3}, nil}
	found := tree.PointSearch(empty_entry)
	found_v := *found.item.value

	AssertEqual(t, found_v, val)
}
