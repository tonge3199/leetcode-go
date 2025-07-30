package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric_01(root *TreeNode) bool {

	// put nodes in arr to compare
	//left_sub2arr := make([]*TreeNode, 0)
	//right_sub2arr := make([]*TreeNode, 0)
	//
	//left_subTree := root.Left
	//right_subTree := root.Right

	// a Iteration way
	return false

}

func isMirror(left, right *TreeNode) bool {
	// If both nodes are nil, they are symmetric
	if left == nil && right == nil {
		return true
	}
	// If one node is nil, they are not symmetric
	if left == nil || right == nil {
		return false
	}
	// Check if:
	// 1. Current nodes have same value
	// 2. Left's left subtree mirrors right's right subtree
	// 3. Left's right subtree mirrors right's left subtree
	return left.Val == right.Val &&
		isMirror(left.Left, right.Right) &&
		isMirror(left.Right, right.Left)
}

func isSymmetric_02(root *TreeNode) bool {
	return isMirror(root, root)
}
