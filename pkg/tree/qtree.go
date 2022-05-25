package tree

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
