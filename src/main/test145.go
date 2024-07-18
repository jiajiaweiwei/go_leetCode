package main

import . "MyLikou/src/MyStruct"

func main() {

}

// 二叉树 简单题  求后序遍历序列 前后中
func postorderTraversal(root *TreeNode) []int {
	// 返回值
	resultNum := make([]int, 0)
	// 匿名函数
	var postFS func(node *TreeNode)
	postFS = func(node *TreeNode) {
		if node == nil {
			return
		}
		postFS(node.Left)
		postFS(node.Right)
		resultNum = append(resultNum, node.Val)
	}
	postFS(root)
	return resultNum
}
