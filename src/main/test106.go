package main

import . "MyLikou/src/MyStruct"

func main() {

}

// 二叉树 中等题 从中序 后序 构建二叉树 同理105
// 方法1 递归
// 先根据 后序遍历倒数确定根节点
func buildTree106(inorder []int, postorder []int) *TreeNode {
	// 后序遍历序列确认根节点
	// 哈希表存中序遍历序列
	indexMap := make(map[int]int, len(inorder))
	for i, v := range inorder {
		indexMap[v] = i
	}
	// 递归匿名函数
	var buildTreeHelper func(int, int) *TreeNode
	buildTreeHelper = func(start int, end int) *TreeNode {
		if start > end {
			return nil
		}
		// 获取根节点
		rootVal := postorder[len(postorder)-1]
		// 更新后序遍历序列
		postorder = postorder[:len(postorder)-1]
		//构建根节点
		root := &TreeNode{Val: rootVal}
		// 获取 哈希表中的 左右子树
		i := indexMap[rootVal]
		root.Right = buildTreeHelper(i+1, end)
		root.Left = buildTreeHelper(start, i-1)
		return root
	}
	return buildTreeHelper(0, len(inorder)-1)
}
