package main

import . "MyLikou/src/MyStruct"

func main() {

}

// 中等题 返回二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	//特殊情况处理
	if root == nil {
		return [][]int{}
	}
	//返回值
	resultNum := make([][]int, 0)
	//队列
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		//每一层的节点数
		layerLen := len(queue)
		intNum2 := make([]int, 0) //二级切片
		for j := 0; j < layerLen; j++ {
			//队头出队
			node := queue[0]
			intNum2 = append(intNum2, node.Val)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			//一级切片
		}
		resultNum = append(resultNum, intNum2)
	}
	return resultNum
}
