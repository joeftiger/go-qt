package tree

import "fmt"

type QTree[T any] struct {
	compare func(a, b T) (bool, int)
	root    *QNode[T]
}

func NewQTree[T any](compare func(a, b T) (bool, int)) QTree[T] {
	return QTree[T]{compare, nil}
}

func (tree *QTree[T]) NaiveInsert(node QNode[T]) {
	if tree.root == nil {
		tree.root = &node
	} else {
		tree.root.NaiveInsert(node, tree.compare)
	}
}

func (tree *QTree[T]) PointSearch(key T) *QNode[T] {
	if tree.root == nil {
		return nil
	} else {
		return tree.root.PointSearch(key, tree.compare)
	}
}

func (tree *QTree[T]) Traverse(f func(*T)) {
	if tree.root != nil {
		tree.root.Traverse(f)
	}
}

type QNode[T any] struct {
	parent   *QNode[T]
	key      T
	children []*QNode[T]
}

func NewQNode[T any](key T, dim int) QNode[T] {
	return QNode[T]{
		parent:   nil,
		key:      key,
		children: make([]*QNode[T], 0b1<<dim),
	}
}

func (n *QNode[T]) NaiveInsert(node QNode[T], compare func(a, b T) (bool, int)) {
	equal, quad := compare(n.key, node.key)

	if equal {
		fmt.Println("equal, cannot insert")
	} else {
		if n.children[quad] == nil {
			n.children[quad] = &node
		} else {
			node.parent = n
			n.children[quad].NaiveInsert(node, compare)
		}
	}

}

func (n *QNode[T]) PointSearch(key T, compare func(a, b T) (bool, int)) *QNode[T] {
	equal, quad := compare(n.key, key)

	if equal {
		return n
	} else if n.children[quad] != nil {
		return n.children[quad].PointSearch(key, compare)
	} else {
		return nil
	}
}

func (n *QNode[T]) Traverse(f func(*T)) {
	f(&n.key)

	for _, c := range n.children {
		if c != nil {
			c.Traverse(f)
		}
	}
}
