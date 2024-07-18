package main

import . "MyLikou/src/MyStruct"

// 简单题 反转链表1
func reverseList(head *ListNode) *ListNode {
	// 处理特殊情况
	if head == nil || head.Next == nil {
		return head
	}

	//改变链表指向方向键即可
	//临时节点
	tempLeft := head
	tempRight := head.Next.Next
	for tempRight != nil {
		tempRight = tempLeft.Next.Next
		tempLeft.Next.Next = tempLeft
		tempLeft.Next = tempLeft
		tempRight.Next = tempLeft
	}
	head = tempRight
	return head
}
