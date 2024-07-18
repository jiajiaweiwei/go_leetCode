package main

import (
	. "MyLikou/src/MyStruct"
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.MinInt32)
}

// 二叉树 困难题 二叉树中的最大路径和
// 方法1 递归(回溯的思想) 求每个节点的最大贡献值
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	// 节点的最大贡献值函数
	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		//递归计算左右节点 的最大贡献值 负数不要，直接置0
		leftGain := max124(maxGain(node.Left), 0)
		rightGain := max124(maxGain(node.Right), 0)
		// 求最大路径
		priceNewPath := node.Val + leftGain + rightGain
		// 更新最大路径
		maxSum = max124(maxSum, priceNewPath)
		// 返回节点的最大贡献值
		return node.Val + max124(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}

func max124(x, y int) int {
	if x > y {
		return x
	}
	return y
}
