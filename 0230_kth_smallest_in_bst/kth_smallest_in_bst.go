package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/* RECURSIVE VERSION
func kthSmallest(root *TreeNode, k int) int {
	var inorderTraversal func(r *TreeNode)
	buf := []int{}

	inorderTraversal = func(r *TreeNode) {
		if r == nil || len(buf) >= k {
			return
		}
		inorderTraversal(r.Left)
		buf = append(buf, r.Val)
		inorderTraversal(r.Right)
	}

	inorderTraversal(root)
	return buf[k-1]
}
*/

// Iterative version
func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	buf := []int{}
	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		buf = append(buf, curr.Val)
		curr = curr.Right
	}
	return buf[k-1]
}
