package main

import . "MyLikou/src/MyStruct"

func main() {

}

// 二叉树 中等题 锯齿形层序遍历
// 方法1 层序遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	//特殊情况处理
	if root == nil {
		return [][]int{}
	}
	resultNum := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	direction := 1
	tempNum := make([]int, 0)
	for len(queue) > 0 {
		l := len(queue)
		//结果
		if direction%2 == 0 {
			for i := l; i > 0; i-- {
				node := queue[i-1]
				tempNum = append(tempNum, node.Val)
			}
		} else {
			for i := 0; i < l; i++ {
				node := queue[i]
				tempNum = append(tempNum, node.Val)
			}
		}
		//入队
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
		resultNum = append(resultNum, tempNum)
		tempNum = []int{}
		direction++
	}
	return resultNum
}
