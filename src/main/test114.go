package main

import (
	. "MyLikou/src/MyStruct"
)

func main() {

}

// 二叉树 中等题 二叉树原地展开为链表
// 方法1 递归实现的前序遍历 + 根据前序遍历序列特性 原地调整树
func flatten(root *TreeNode) {
	// 构建前序遍历递归
	var preFS func(node *TreeNode)
	preFS = func(node *TreeNode) {
		if node == nil {
			return
		}
		//前序遍历
		//调整树
		if node.Left != nil {
			if node.Right != nil {
				tempNode := node.Right
				node.Right = nil
				tempRight := getRightEnd(node.Left)
				node.Right = node.Left
				node.Left = nil
				tempRight.Right = tempNode
			} else {
				node.Right = node.Left
				node.Left = nil
			}
		}
		preFS(node.Left)
		preFS(node.Right)
	}
	preFS(root)
}

// 递归找到最右节点
func getRightEnd(node *TreeNode) *TreeNode {
	for node.Right != nil {
		node = node.Right
	}
	return node
}
