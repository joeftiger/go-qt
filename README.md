# go-qt
A Quadtree implemented in Go.


## Basic Usage

Before initializing the tree, one needs to decide a type for the item they want to store.

Say you want to only store vectors of dimension 3 with no associated values. Also, let the entries of the vectors be of type `float32`.

Then, you pass the type and a comparison function, which compares your item. Here, we're going to use the provided `compare_ordered()` function, which can compare `Integers` and `Floats` and `~string`. (See <https://golang.org/x/exp/constraints> for more information)

```go
tree := NewQTree[KVEntry](compare_ordered)
```

To insert an item:

```go
item := []float32{1.434, 21.222, 332.23432}
node := NewQNode(item, dim)
tree.NaiveInsert(node)
```

And search for an item:

```go
found := tree.PointSearch([]float32{1.434, 21.222, 332.23432})
fmt.Println(found.item)
```

Please be aware that `PointSearch()` returns a `*QNode[T]` where `T` corresponds to `[]float32` in our case. The returned node has an property called `.item` to retrieve your stored item.


## Custom Structs

One can also easily store any kind of struct. In this exmaple, we will show you how to use the tree as a KV Store.

First, let's define a simple Key-Value pair structure, where we will use a list of strings as keys and the corresponding value is a string, too.

```go
type KVEntryString struct {
	key   *[]string
	value *string
}
```

Next, we need to specify a comparison function for our struct. On what properties do we want to compare it? Since strings are ordered, we can again rely on the provided `compare_ordered()`. 

```go
func compare_kv_str(a, b KVEntryString) (equal bool, quad int) {
	return compare_ordered[string](*a.key, *b.key)
}
```

This is all we need. Now, we can insert elements as follows:

```go
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
```

If we want to fetch now the associated value, then we do the following:

```go
search_for := []string{"switzerland", "bern"}
empty_entry := KVEntryString{&search_for, nil} // We pass nil as the value parameter
found := tree.PointSearch(empty_entry)
if found != nil {
	fmt.Println(*found.item.value) // always access the item, provided it exists!
}
```


## The comparison function

The comparison function is basically the heart of this data structure. We provide a `compare_ordered()` function which compares `Integers`, `Floats` and `~string`. 

This is the built in comparison function. Your comparison function should behave, in regards to the return value, the same. This means, you should return a pair of `(equal bool, quad int)` where `equal` is `true` iff `a` and `b` are the same. It's good practice to return for `quad` a value like `-1` $\notin [0,2^n] \subset \mathbb{N}$, in this case.
However, if `a` and `b` are not the same, then `equal` should be `false` and your `quad` should be the integer from 0 to $2^n-1$ in which the item you're lookging for (`b`) lies in.

```go
func compare_ordered[T constraints.Ordered](a, b []T) (equal bool, quad int) {

	if reflect.DeepEqual(a, b) {
		return true, -1
	}

	quad = 0

	for i := 0; i < len(a); i++ {
		if b[i] >= a[i] {
			quad += 0b1 << i
		}
	}

	return false, quad
}
```


