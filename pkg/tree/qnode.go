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
