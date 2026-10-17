package leetcode

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(n^2)
// Space Complexity: O(n)
func recursion(node *TreeNode, s string, ans *[]string) {
	if node == nil {
		return
	}

	if s == "" {
		s = strconv.Itoa(node.Val)
	} else {
		s += "->" + strconv.Itoa(node.Val)
	}

	if node.Left == nil && node.Right == nil {
		*ans = append(*ans, s)
		return
	}

	if node.Left != nil {
		recursion(node.Left, s, ans)
	}
	if node.Right != nil {
		recursion(node.Right, s, ans)
	}
}

func binaryTreePaths(root *TreeNode) []string {
	ans := []string{}

	if root == nil {
		return ans
	}

	recursion(root, "", &ans)

	return ans
}
