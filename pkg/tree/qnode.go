package tree

import (
	"fmt"
)

type QNode[T any] struct {
	parent   *QNode[T]
	item     T
	children []*QNode[T]
}

func NewQNode[T any](item T, dim int) QNode[T] {
	return QNode[T]{
		parent:   nil,
		item:     item,
		children: make([]*QNode[T], 0b1<<dim),
	}
}

func (n *QNode[T]) IsRoot() bool {
	return n.parent == nil
}

func (n *QNode[T]) IsLeaf() bool {
	for _, c := range n.children {
		if c != nil {
			return false
		}
	}

	return true
}

func (n *QNode[T]) NaiveInsert(node QNode[T], compare func(a, b T) (bool, int)) {
	equal, quad := compare(n.item, node.item)

	if equal {
		fmt.Println("Updating not implemented yet.")
	} else {
		if n.children[quad] == nil {
			node.parent = n
			n.children[quad] = &node
		} else {
			n.children[quad].NaiveInsert(node, compare)
		}
	}

}

func (n *QNode[T]) PointSearch(item T, compare func(a, b T) (bool, int)) *QNode[T] {
	equal, quad := compare(n.item, item)

	if equal {
		return n
	} else if n.children[quad] != nil {
		return n.children[quad].PointSearch(item, compare)
	} else {
		return nil
	}
}

func (n *QNode[T]) Traverse(f func(*T)) {
	f(&n.item)

	for _, c := range n.children {
		if c != nil {
			c.Traverse(f)
		}
	}
}

// Performs range search
// from_item, to_item are both items that should be handled by the compare function
// n := the node we are in right now
func (n *QNode[T]) RangeSearch(bound []T, compare func(a, b T) (bool, int), dim int) []*QNode[T] {

	stack := make([]*QNode[T], 0) // can be improved, e.g. how about log(n)?

	n.RangeSearch_Recursion(bound, compare, &stack, dim)

	return stack
}

func (n *QNode[T]) RangeSearch_Recursion(bound []T, compare func(a, b T) (bool, int), stack *[]*QNode[T], dim int) {

	if IsInHypercube(n, bound, compare, dim) {
		(*stack) = append(*stack, n)

		// Continue search inside all quads until it's no longer in hypercube,
		for _, child := range n.children {
			if child != nil {
				child.RangeSearch_Recursion(bound, compare, stack, dim)
			}
		}

	}
}

// a point P lies inside a Hypercube iff compare() returns n different values
// for n comparisons (with n boundary points) with P

func IsInHypercube[T any](pnode *QNode[T], hypercube []T, compare func(a, b T) (bool, int), dim int) bool {

	p := pnode.item // The point to be tested

	table := map[int]bool{} // Store directionals here with a true, then count

	for _, corner := range hypercube {
		equal, quad := compare(corner, p)

		if equal == true {
			return true
		} else {
			if !table[quad] {
				table[quad] = true
			}
		}
	}

	// Returns true, if the filled table is of length 2^n, namely all quadrants appeared
	return len(table) == 0b1<<dim
}
