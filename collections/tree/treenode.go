package tree

type TreeNode[T any] struct {
	Value    T
	Children []*TreeNode[T]
}
