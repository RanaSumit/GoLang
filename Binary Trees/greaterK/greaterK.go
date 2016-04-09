package bstrees

// FindFirstGreaterK returns the first node in tree greater then k
// (when encountering in-order) or nil if such a node isn't present.
// The time complexity is O(log(n)). The O(1) additional space is needed.
type BSTree struct {
	
	Data  interface{}
	left  *BSTree
	right *BSTree
}
func FindFirstGreaterK(tree *BSTree, k int) *BSTree {
	var g *BSTree
	node := tree
	for node != nil {
		if k < node.Data.(int) {
			g, node = node, node.left
		} else {
			node = node.right
		}
	}
	return g
}