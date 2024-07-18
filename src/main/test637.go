package main

import . "MyLikou/src/MyStruct"

func main() {

}

// 二叉树 简单题 求二叉树的层平均值
// 方法1 层序遍历
func averageOfLevels(root *TreeNode) []float64 {
	//特殊情况处理
	if root == nil {
		return []float64{}
	}
	//函数返回值
	resultSlice := make([]float64, 0)
	//层序遍历队列
	queue := []*TreeNode{}
	queue = append(queue, root)
	//每一层的节点值的和
	tempSum := 0
	//开始层序遍历
	for len(queue) > 0 {
		//每一层的节点数
		cengLen := len(queue)
		tempCengLen := cengLen
		//每一层的节点值的和
		tempSum = 0
		//对每一层的节点值求和
		for tempCengLen > 0 {
			node := queue[0]
			tempSum += node.Val
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			tempCengLen--
		}
		//将平均数存入返回切片
		resultSlice = append(resultSlice, float64(tempSum)/float64(cengLen))
	}
	return resultSlice
}
