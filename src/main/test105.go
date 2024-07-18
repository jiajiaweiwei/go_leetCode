package main

import . "MyLikou/src/MyStruct"

func main() {

}

// 根据前序中序构造二叉树
// 方法1 使用递归 + 切片传递左右子树
func buildTree(preorder []int, inorder []int) *TreeNode {
	indexMap := make(map[int]int, len(inorder))
	for i, v := range inorder {
		indexMap[v] = i
	}
	// 递归匿名函数
	var buildTreeHelper func(int, int) *TreeNode // inStart 和 inEnd，表示当前子树在中序遍历数组中的起始位置和结束位置。
	buildTreeHelper = func(inStart, inEnd int) *TreeNode {
		if inStart > inEnd {
			return nil
		}
		// 根据前序遍历数组中的第一个元素构建根节点
		rootVal := preorder[0]
		root := &TreeNode{Val: rootVal}
		// 在中序遍历数组中查找根节点的位置(使用哈希表)
		rootIndex := indexMap[rootVal]
		// 递归构建左子树和右子树
		preorder = preorder[1:] // 新的前序遍历
		root.Left = buildTreeHelper(inStart, rootIndex-1)
		root.Right = buildTreeHelper(rootIndex+1, inEnd)
		return root
	}
	return buildTreeHelper(0, len(inorder)-1)
}

// 复习 前序+中序->构建二叉树
func buildTree105(preorder []int, inorder []int) *TreeNode {
	// 根据前序遍历确认根节点
	// 根据中序遍历 确认左右子树
	// 递归构建二叉树
	// 使用哈希表 将中序遍历序列放入哈希表中，降低 确认左右子树的复杂度
	indexMap := make(map[int]int, len(inorder))
	// 将中序遍历节点放入哈希表
	for i, v := range inorder {
		indexMap[v] = i
	}
	// 递归匿名函数
	var buildTreeHelper func(int, int) *TreeNode
	// 返回节点 构建二叉树
	buildTreeHelper = func(start int, end int) *TreeNode {
		if start > end {
			return nil
		}
		// 依次获取 前序遍历的头节点 确认根节点
		rootVal := preorder[0]
		root := &TreeNode{Val: rootVal}
		preorder = preorder[1:] //  新的前序遍历序列
		//根据中序遍历哈希表 确认 左右子树序列 以来构建二叉树
		rootIndex := indexMap[rootVal]
		root.Left = buildTreeHelper(start, rootIndex-1)
		root.Right = buildTreeHelper(rootIndex+1, end)
		return root
	}
	return buildTreeHelper(0, len(inorder)-1)
}
