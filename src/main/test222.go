package main

import (
	. "MyLikou/src/MyStruct"
	"sort"
)

func main() {

}

// 二叉树 简单题 完全二叉树的节点个数
// 方法1 BFS记录节点个数
func countNodes(root *TreeNode) int {
	// 异常情况处理
	if root == nil {
		return 0
	}
	resultNum := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		resultNum += l
		for l > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			l--
		}
	}
	return resultNum
}

// 方法2 (重点) 二分查找 + 位运算   （利用了题目所给的 完全二叉树的特性）
func countNodes2(root *TreeNode) int {
	// 异常情况处理
	if root == nil {
		return 0
	}
	tall := 0
	//先使用递归求出树高
	for node := root; node.Left != nil; node = node.Left {
		tall++
	}
	// 使用位运算+二分查找猜出最后一层元素个数
	return sort.Search(1<<(tall+1), func(k int) bool {
		if k <= 1<<tall {
			return false
		}
		bits := 1 << (tall - 1)
		node := root
		for node != nil && bits > 0 {
			if bits&k == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
			bits >>= 1
		}
		return node == nil
	}) //最后一个节点的下标就是完全二叉树节点的个数
}
