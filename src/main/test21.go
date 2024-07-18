package main

import . "MyLikou/src/MyStruct"

func main() {

}

// 简单题  合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	//返回链表头节点
	resultNode := new(ListNode)
	tempNode := resultNode
	for list1 != nil || list2 != nil {
		list1, list2, tempNode.Next = minNode21(list1, list2)
		tempNode = tempNode.Next
	}
	return resultNode.Next
}
func minNode21(list1, list2 *ListNode) (*ListNode, *ListNode, *ListNode) {
	temps := new(ListNode)
	if list1 == nil {
		temps = list2
		list2 = list2.Next
		return list1, list2, temps
	}
	if list2 == nil {
		temps = list1
		list1 = list1.Next
		return list1, list2, temps
	}
	if list1.Val > list2.Val {
		temps = list2
		list2 = list2.Next
		return list1, list2, temps
	}
	temps = list1
	list1 = list1.Next
	return list1, list2, temps
}
