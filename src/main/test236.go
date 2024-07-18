package main

import (
	. "MyLikou/src/MyStruct"
	"fmt"
)

func main() {
	// 构建测试用例二叉树
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}

	// 构建 p 和 q 节点
	p := root.Left
	q := root.Right

	// 调用 lowestCommonAncestor 函数
	result := lowestCommonAncestor(root, p, q)

	// 输出结果
	if result != nil {
		fmt.Printf("Lowest Common Ancestor of %d and %d is %d\n", p.Val, q.Val, result.Val)
	} else {
		fmt.Println("Lowest Common Ancestor not found")
	}
}

// 二叉树 中等题 二叉树最近的公共祖先
// 方法1 中等题 使用map存储节点的父节点，相当于可以从下网上找父节点
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// p节点的父节点路径map
	pPathMap := make(map[int]bool)
	// 记录所有节点 父节点的map
	allHeadMap := make(map[int]*TreeNode)
	// 构建allHeadMap
	var dfs func(node, parent *TreeNode)
	dfs = func(node, parent *TreeNode) {
		if node == nil {
			return
		}
		allHeadMap[node.Val] = parent
		dfs(node.Left, node)
		dfs(node.Right, node)
	}
	dfs(root, nil)
	for p != nil {
		pPathMap[p.Val] = true
		p = allHeadMap[p.Val]
	}
	for q != nil {
		if pPathMap[q.Val] {
			return q
		}
		q = allHeadMap[q.Val]
	}
	return nil
}
