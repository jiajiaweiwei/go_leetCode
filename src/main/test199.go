package main

import . "MyLikou/src/MyStruct"

func main() {

}

// 二叉树 中等题 二叉树的右视图
// 方法1 层序遍历
func rightSideView(root *TreeNode) []int {
	//特殊情况处理
	if root == nil {
		return []int{}
	}
	//返回值
	resultNum := make([]int, 0)
	//层序遍历队列
	queue := make([]*TreeNode, 0)
	//头节点入队
	queue = append(queue, root)
	//开始遍历
	for len(queue) > 0 {
		//每层的节点数
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)

			}
			if i == l-1 {
				resultNum = append(resultNum, node.Val)
			}
		}
	}
	return resultNum
}
