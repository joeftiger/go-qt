package tree

import (
	"testing"
)

type KVEntryString struct {
	key   *[]string
	value *string
}

// // The compare function which basically relies on CompareOrdered
func compareKvStr(a, b KVEntryString) (equal bool, quad int) {
	return CompareOrdered[string](*a.key, *b.key)
}

// An example usage for the KV Store.
func TestKVStringStore(t *testing.T) {
	const dim = 3
	tree := NewQTree[KVEntryString](dim, compareKvStr)

	// First entry
	keyFr := []string{"Switzerland", "Fribourg"} // key composite
	valFr := "Hello Fribourg"
	entryFr := KVEntryString{&keyFr, &valFr}
	tree.NaiveInsert(entryFr)

	// Second entry
	keyBe := []string{"Switzerland", "Bern"} // key composite
	valBe := "Hello Bern"
	entryBe := KVEntryString{&keyBe, &valBe}
	tree.NaiveInsert(entryBe)

	// Search first for bern
	emptyEntry := KVEntryString{&([]string{"Switzerland", "Bern"}), nil}
	found := tree.PointSearch(emptyEntry)
	if found != nil {
		AssertEqual(t, *found.item.value, valBe)
		t.Log(*found.item.value) // will print "Hello Bern"
	} else {
		t.Errorf("Expected found = %T, Actual == %T", found, nil)
	}

}
