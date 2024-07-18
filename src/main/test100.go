package main

import . "MyLikou/src/MyStruct"

func main() {

}

// 二叉树 简单题 相同的树
// 方法1 使用DFS
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	// 递归实现DFS
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) //判断左子树和右子树
}
