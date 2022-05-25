package tree

import (
	"testing"
)

type KVEntryString struct {
	key   *[]string
	value *string
}

// // The compare function which basically relies on compare_ordered
func compare_kv_str(a, b KVEntryString) (equal bool, quad int) {
	return compare_ordered[string](*a.key, *b.key)
}

// An example usage for the KV Store.
func TestKVStringStore(t *testing.T) {
	const dim = 3
	tree := NewQTree[KVEntryString](compare_kv_str)

	// First entry
	key_fr := []string{"switzerland", "fribourg"} // key composite
	val_fr := "hello fribourg"
	entry_fr := KVEntryString{&key_fr, &val_fr}
	node_fr := NewQNode(entry_fr, dim)
	tree.NaiveInsert(node_fr)

	// Second entry
	key_be := []string{"switzerland", "bern"} // key composite
	val_be := "hello bern"
	entry_be := KVEntryString{&key_be, &val_be}
	node_be := NewQNode(entry_be, dim)
	tree.NaiveInsert(node_be)

	// Search first for bern
	empty_entry := KVEntryString{&([]string{"switzerland", "bern"}), nil}
	found := tree.PointSearch(empty_entry)
	if found != nil {
		found_val := *found.item.value
		AssertEqualDeep(t, found_val, val_be)
	}

}
