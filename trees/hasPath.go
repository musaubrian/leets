package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	stack := []*TreeNode{root}
	sums := []int{targetSum - root.Val}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curSum := sums[len(sums)-1]
		sums = sums[:len(sums)-1]

		if node.Left == nil && node.Right == nil && curSum == 0 {
			return true
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
			sums = append(sums, curSum-node.Right.Val)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
			sums = append(sums, curSum-node.Left.Val)
		}
	}

	return false
}
