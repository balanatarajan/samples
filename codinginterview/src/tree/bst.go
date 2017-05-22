package tree

type TreeNode struct {
	Left, Right *TreeNode
	Val         int
}

// Create BST from an array of ints. root is the node around which the
// subtree will be formed
func CreateBT(ints []int, root *TreeNode) *TreeNode {

	if len(ints) == 0 {
		return nil
	}

	ctr := len(ints) / 2

	root = &TreeNode{Val: ints[ctr]}

	root.Left = CreateBT(ints[:ctr], root)
	root.Right = CreateBT(ints[ctr+1:], root)

	return root
}
