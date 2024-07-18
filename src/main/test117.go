package main

import . "MyLikou/src/MyStruct"

func main() {

}

// 二叉树 中等题  给定一个二叉树根节点 将每一层构建成链表
// 方法1 BFS
func connect(root *TreeNodeList) *TreeNodeList {
	// 特殊情况处理
	if root == nil {
		return nil
	}
	// BFS队列
	queue := make([]*TreeNodeList, 0)
	queue = append(queue, root)
	//BFS
	for len(queue) > 0 {
		l := len(queue)
		for l > 0 {
			// 头节点出队
			node := queue[0]
			// 构建链表
			if l > 1 {
				node.Next = queue[1]
			}
			queue = queue[1:]
			//入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			l--
		}
	}
	return root
}

// 方法2 空间复杂度O1的解法 利用了next指针
func connect2(root *TreeNodeList) *TreeNodeList {
	if root == nil {
		return nil
	}
	// 每层的第一个节点
	firstNode := root
	for firstNode != nil {
		// 下一层的起始节点
		nextLevelStart := &TreeNodeList{}
		tempNode := nextLevelStart
		for firstNode != nil {
			if firstNode.Left != nil {
				tempNode.Next = firstNode.Left
				tempNode = tempNode.Next
			}
			if firstNode.Right != nil {
				tempNode.Next = firstNode.Right
				tempNode = tempNode.Next
			}
			firstNode = firstNode.Next
		}
		// 进入下一层
		firstNode = nextLevelStart.Next
	}
	return root
}
