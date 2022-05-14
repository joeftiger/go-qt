package tree

type QTree[T any] struct {
	dim     int
	compare func(a, b T) int
	root    *QNode[T]
}

func NewQTree[T any](dim int, compare func(a, b T) int) QTree[T] {
	return QTree[T]{dim, compare, nil}
}

func (tree *QTree[T]) Conjugate(q int) int {
	return ((q + 1) % tree.dim) + 1
}

func (tree *QTree[T]) NaiveInsert(node QNode[T]) {
	if tree.root == nil {
		tree.root = &node
	} else {
		tree.root.NaiveInsert(node, tree.compare)
	}
}

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

func (n *QNode[T]) NaiveInsert(node QNode[T], compare func(a, b T) int) {
	quad := compare(n.item, node.item)

	if n.children[quad] == nil {
		n.children[quad] = &node
	} else {
		node.parent = n
		n.children[quad].NaiveInsert(node, compare)
	}
}