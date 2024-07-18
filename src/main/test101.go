package main

import . "MyLikou/src/MyStruct"

func main() {

}

// 简单题 对称二叉树
// 方法1 使用DFS
func isSymmetric(root *TreeNode) bool {
	return DFSTemp(root, root)
}

// DFSTemp 递归函数
func DFSTemp(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && DFSTemp(p.Left, q.Right) && DFSTemp(p.Right, q.Left)
}
