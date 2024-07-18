package main

import (
	. "MyLikou/src/MyStruct"
	"math"
)

func main() {

}

// 二叉树 中等题 验证二叉搜索树 既判断其是否是一个有效的二叉搜索树
// 方法1 递归
func isValidBST(root *TreeNode) bool {
	var verifySearchTree func(node *TreeNode, pre int, end int) bool
	verifySearchTree = func(node *TreeNode, pre int, end int) bool {
		if node == nil {
			return true
		}
		if node.Val <= pre || node.Val >= end {
			return false
		}
		return verifySearchTree(node.Left, pre, node.Val) && verifySearchTree(node.Right, node.Val, end)
	}
	return verifySearchTree(root, math.MinInt64, math.MaxInt64)
}

// 方法2 中序遍历
func isValidBST2(root *TreeNode) bool {
	//先写中序遍历
	tempBool, pre := true, &TreeNode{}
	var midFs func(node *TreeNode)
	midFs = func(node *TreeNode) {
		if node == nil {
			return
		}
		midFs(node.Left)
		if pre != nil && pre.Val >= node.Val {
			tempBool = false
			return
		}
		pre = node
		midFs(node.Right)
	}
	// 中序遍历
	midFs(root)
	return tempBool
}
