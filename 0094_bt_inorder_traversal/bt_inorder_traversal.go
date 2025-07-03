package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var traverse func(r *TreeNode)
	ans := []int{}

	traverse = func(r *TreeNode) {
		if r == nil {
			return
		}
		traverse(r.Left)
		ans = append(ans, r.Val)
		traverse(r.Right)
	}
	traverse(root)
	return ans
}
