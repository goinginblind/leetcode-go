package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// If a node has 0 or 1 children its an easy case
// If it has 2, then you just find minimum val node in the right subtree
// or maximum in the left
// And then simply replace val, delete the min/max node

func findMinimumNode(root *TreeNode) int {
	for root.Left != nil {
		root = root.Left
	}
	return root.Val
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			minNodeKey := findMinimumNode(root.Right)
			root.Val = minNodeKey
			root.Right = deleteNode(root.Right, minNodeKey)
		}
	}
	return root
}
