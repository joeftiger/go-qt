package tree

type QTree[T any] struct {
	compDim int
	compare func(a, b T) (bool, int)
	root    *QNode[T]
}

func NewQTree[T any](compDim int, compare func(a, b T) (bool, int)) QTree[T] {
	return QTree[T]{compDim, compare, nil}
}

func (tree *QTree[T]) CreateNode(item T) QNode[T] {
	return NewQNode(item, tree.compDim)
}

func (tree *QTree[T]) NaiveInsert(item T) {
	tree.NaiveInsertNode(tree.CreateNode(item))
}

func (tree *QTree[T]) NaiveInsertNode(node QNode[T]) {
	if tree.root == nil {
		tree.root = &node
	} else {
		tree.root.NaiveInsert(node, tree.compare)
	}
}

func (tree *QTree[T]) PointSearch(item T) *QNode[T] {
	if tree.root == nil {
		return nil
	} else {
		return tree.root.PointSearch(item, tree.compare)
	}
}

func (tree *QTree[T]) Traverse(f func(*T)) {
	if tree.root != nil {
		tree.root.Traverse(f)
	}
}

func (tree *QTree[T]) RangeSearch(bound []T) []*QNode[T] {
	if tree.root == nil {
		return nil
	} else {
		return tree.root.RangeSearch(bound, tree.compare, tree.compDim)
	}
}
