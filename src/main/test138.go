package main

import . "MyLikou/src/MyStruct"

func main() {

}

// 链表中等题 随机链表的复制 既要保存链表 也要保存随机指针指向的节点下标
// 方法1 两次遍历 head 不带投头节点
func copyRandomList(head *Node) *Node {
	//用户第二次循环的哈希
	m := make(map[*Node]*Node)
	//第一次 构建普通的链表
	//返回值
	resultNode := new(Node)
	resultTemp := resultNode
	tempNode := head
	for tempNode != nil {
		resultTemp.Next = &Node{
			Val:    tempNode.Val,
			Next:   nil,
			Random: nil,
		}
		m[tempNode] = resultTemp.Next
		resultTemp = resultTemp.Next
		tempNode = tempNode.Next
	}
	// 第二次遍历 根据哈希构建random链
	tempNode = head
	resultTemp = resultNode.Next
	for tempNode != nil {
		if tempNode.Random != nil {
			resultTemp.Random = m[tempNode.Random]
		}
		tempNode = tempNode.Next
		resultTemp = resultTemp.Next
	}
	return resultNode.Next
}

//方法2 迭代加节点差拆分 拆分并插入一个节点，random就是原来random的后面那个新插入的节点
