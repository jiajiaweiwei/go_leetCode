package main

import . "MyLikou/src/MyStruct"

func main() {

}

// 二叉树 简单题 路径总和
// 方法1 使用递归
func hasPathSum(root *TreeNode, targetSum int) bool {
	// 异常情况处理
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	//开始递归，判断 子节点的值是否等于目标值减去根节点的值
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum2(root.Right, targetSum-root.Val)
}

// 方法2 使用BFS
func hasPathSum2(root *TreeNode, targetSum int) bool {
	// 异常情况处理
	if root == nil {
		return false
	}
	// BFS队列
	queue := make([]*TreeNode, 0)
	// 存储路径的队列
	tempQueue := make([]int, 0)
	//初始化队列
	queue = append(queue, root)
	tempQueue = append(tempQueue, root.Val)
	// 开始BFS
	for len(queue) > 0 {
		l := len(queue)
		for l > 0 {
			node := queue[0]
			queue = queue[1:]
			temp := tempQueue[0]
			tempQueue = tempQueue[1:]
			if node.Left == nil && node.Right == nil {
				if temp == targetSum {
					return true
				}
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
				tempQueue = append(tempQueue, node.Left.Val+temp)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				tempQueue = append(tempQueue, node.Right.Val+temp)
			}
			l--
		}

	}
	return false
}
