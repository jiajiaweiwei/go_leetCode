package main

import (
	. "MyLikou/src/MyStruct"
	"fmt"
)

func main() {
	// 创建一棵三层的二叉树
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	// 调用 maxDepth 函数求解二叉树的最大深度
	depth := maxDepth(root)
	fmt.Println("二叉树的最大深度为:", depth)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 简单题 求二叉树的最大深度

// 方法1 递归 深度优先搜索
func maxDepth(root *TreeNode) int {
	//异常情况处理
	if root == nil {
		return 0
	}
	//深度优先搜索
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 方法1.1 实现DFS的第二种实现（DFS模板的方式）
func maxDepth104Temp(root *TreeNode) int {
	// 存储最终结果
	resultNum := 0
	// 初始化当前结果
	startNum := 0
	resultNum = dfs104Temp(root, startNum)
	return resultNum
}

// DFS 模板递归函数 返回值是最大深度
func dfs104Temp(node *TreeNode, deep int) int {
	// 构造递归函数
	if node == nil {
		return deep
	}
	// 更新当前结果
	deep++
	// 到达叶子节点 直接返回
	if node.Left == nil && node.Right == nil {
		return deep
	}

	// 递归遍历左右子树
	leftDepth := dfs104Temp(node.Left, deep)
	rightDepth := dfs104Temp(node.Right, deep)

	// 返回左右子树深度的最大值
	return max104(leftDepth, rightDepth)
}

func max104(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 方法2 层次遍历 需要使用到队列 返回层数 既是二叉树最大深度
func maxDepth104(root *TreeNode) int {
	//异常情况处理
	if root == nil {
		return 0
	}
	//层次遍历 获取层数
	//队列
	queue := []*TreeNode{}
	queue = append(queue, root)
	//循环遍历 实现 层次遍历
	//层数
	deepNum := 0
	for len(queue) != 0 {
		//每一层的节点个数
		cengLen := len(queue)
		//遍历每一层
		for cengLen > 0 {
			root := queue[0]
			queue = queue[1:]
			if root.Left != nil {
				queue = append(queue, root.Left)
			}
			if root.Left != nil {
				queue = append(queue, root.Right)
			}
			cengLen--
		}
		deepNum++

	}
	return deepNum
}
