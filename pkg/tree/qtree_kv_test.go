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

// The compare function which basically relies on CompareOrdered
func compareKvInt(a, b KVEntry) (equal bool, quad int) {
	return CompareOrdered[int](a.key, b.key)
}

// An example usage for the KV Store.
func TestKVStore(t *testing.T) {
	const dim = 3
	tree := NewQTree[KVEntry](dim, compareKvInt)

	// First entry
	val := "hello world"
	key := []int{1, 2, 3}
	entry := KVEntry{key, &val}
	tree.NaiveInsert(entry)

	// Search
	emptyEntry := KVEntry{[]int{1, 2, 3}, nil}
	found := tree.PointSearch(emptyEntry)
	AssertEqual(t, *found.item.value, val)
}
