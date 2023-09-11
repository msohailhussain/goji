package tree

type TreeNode[T any] struct {
	Value    T
	Children []*TreeNode[T]
}

func (t TreeNode[T]) String() string {
	return "" // TODO
}
