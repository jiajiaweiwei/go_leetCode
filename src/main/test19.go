package main

import . "MyLikou/src/MyStruct"

func main() {

}

// 链表中等题 用一趟扫描实现 删除链表倒数第n个节点 并返回链表的头节点
// 方法1 可以使用双指针 实现一次遍历 就找到 倒数第n个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//哑节点
	muteNode := &ListNode{Next: head}
	//提前指针
	firstNode := muteNode
	//滞后指针
	afterNode := muteNode
	//第一次遍历 让提前指针遍历n+1次
	for i := 0; firstNode != nil && i < n+1; i++ {
		firstNode = firstNode.Next
	}
	//第二次遍历让 滞后指针遍历到提前指针指向空即可
	for firstNode != nil {
		firstNode = firstNode.Next
		afterNode = afterNode.Next
	}
	//删除afterNode 后面的节点
	if afterNode.Next.Next != nil {
		afterNode.Next = afterNode.Next.Next
	} else {
		afterNode.Next = nil
	}
	return muteNode.Next
}
