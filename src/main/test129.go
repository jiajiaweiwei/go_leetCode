package main

import . "MyLikou/src/MyStruct"

func main() {

}

// 二叉树 中等题 求根节点到叶子节点数字之和
// 方法1 DFS 递归求解
func sumNumbers(root *TreeNode) int {
	var DFS func(node *TreeNode, temp int) int
	DFS = func(node *TreeNode, temp int) int {
		if node == nil {
			// 没有叶子节点就是0 （说明到叶子节点没有路径，直接返回0）
			return 0
		}
		resultNum := temp*10 + node.Val
		if node.Left == nil && node.Right == nil {
			return resultNum
		}
		return DFS(node.Left, resultNum) + DFS(node.Right, resultNum)
	}
	return DFS(root, 0)
}
