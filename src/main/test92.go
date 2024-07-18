package main

import (
	. "MyLikou/src/MyStruct"
)

func main() {
}

// 中等题 反转链表2
// 三指针法 一次遍历 头插法 相当于旋转一次待反转的区间
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	//反转指定区间的链表
	//头节点
	node := &ListNode{
		Next: head,
	}
	tempNode := node
	//找到左指针的前一个节点（待反转区的前一个节点）
	for i := 0; i < left-1; i++ {
		tempNode = tempNode.Next
	}
	//开始反转链表
	cur := tempNode.Next //反转区间的第一个节点 cur也是不变的
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = tempNode.Next
		tempNode.Next = next
	}
	return node.Next
}
