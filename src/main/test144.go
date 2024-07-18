package main

import . "MyLikou/src/MyStruct"

func main() {

}

// 二叉树 简单题 前序遍历 求前序遍历序列 根左右
func preorderTraversal(root *TreeNode) []int {
	// 返回值
	resultNum := make([]int, 0)
	// 递归匿名函数
	var preFS func(node *TreeNode)
	preFS = func(node *TreeNode) {
		if node == nil {
			return
		}
		resultNum = append(resultNum, node.Val)
		preFS(node.Left)
		preFS(node.Right)
	}
	preFS(root)
	return resultNum
}
