package main

import . "MyLikou/src/MyStruct"

func main() {

}

// 二叉树 简单题 翻转二叉树
// 方法1 递归实现翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	// 递归实现深度优先遍历
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}
