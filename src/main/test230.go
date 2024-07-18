package main

import . "MyLikou/src/MyStruct"

func main() {

}

// 二叉树 中等题 二叉搜索树中第K小的元素
// TODO 方法1 利用二叉搜索树的中序遍历的特点
func kthSmallest(root *TreeNode, k int) int {
	// 中序遍历
	var result int // 返回值
	var count int  // 计数
	var midFs func(node *TreeNode)
	midFs = func(node *TreeNode) {
		if node == nil {
			return
		}
		midFs(node.Left)
		count++
		if count == k {
			result = node.Val
			return
		}
		midFs(node.Right)
	}
	midFs(root)
	return result
}
