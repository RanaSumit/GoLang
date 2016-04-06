package bstrees

const (
	maxInt = int(^uint(0) >> 1)
	minInt = -maxInt - 1
)

type BSTree struct {
	Data  interface{}
	left  *BSTree
	right *BSTree
}

// areKeysInRange checks if tree t satisfies the BST property.
func areKeysInRange(t *BSTree, min, max int) bool {
	if t == nil {
		return true
	}
	e := t.Data.(int)
	if e < min || e > max {
		return false
	}
	return areKeysInRange(t.left, min, e) && areKeysInRange(t.right, e, max)
}

// IsBinaryTreeBST returns true if given tree doesn't violate the BST property.
// The time complexity is O(n), and O(h) additional space is needed, where h is
// the binary tree height.
// Note: Only int keys are accepted.
func IsBinaryTreeBST(tree *BSTree) bool {
	return areKeysInRange(tree, minInt, maxInt)
}