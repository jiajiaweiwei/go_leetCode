package main

import . "MyLikou/src/MyStruct"

func main() {

}

// 二叉树 简单题 中序遍历

// 方法1 递归
func inorderTraversal(root *TreeNode) []int {
	// 返回值
	resultNum := make([]int, 0)
	// 匿名函数
	var midFS func(node *TreeNode)
	midFS = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 中序遍历 : 左根右
		midFS(node.Left)
		resultNum = append(resultNum, node.Val)
		midFS(node.Right)
	}
	midFS(root)
	return resultNum
}

// 方法2 迭代
// 方法3 Morris中序遍历
