package main

import . "MyLikou/src/MyStruct"

func main() {

}

// 链表中等题 删除有序链表中的重复元素2
func deleteDuplicates(head *ListNode) *ListNode {
	//哑节点
	muteNode := &ListNode{Next: head}
	temp := muteNode
	//去除重复元素
	for temp.Next != nil && temp.Next.Next != nil {
		if temp.Next.Val == temp.Next.Next.Val {
			x := temp.Next.Val
			//遍历删除相同的元素
			for temp.Next != nil && temp.Next.Val == x {
				temp.Next = temp.Next.Next
			}
		} else {
			temp = temp.Next
		}
	}
	return muteNode.Next
}
