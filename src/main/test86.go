package main

import . "MyLikou/src/MyStruct"

func main() {

}

// 链表中等题 分隔链表 将链表分隔为 小于x 和 大于 x 的前后两个区间，不要求有序
// 方法1 使用两个链表 分别链接大于x和小于x的节点
func partition(head *ListNode, x int) *ListNode {
	//原始链表的哑节点
	dumbNode := head
	//小链表的哑节点
	smallDumbNode := &ListNode{}
	smallTemp := smallDumbNode
	//大链表的哑节点
	bigDumbNode := &ListNode{}
	bigTemp := bigDumbNode
	//遍历
	for dumbNode != nil {
		if dumbNode.Val < x {
			smallTemp.Next = dumbNode
			smallTemp = smallTemp.Next
		} else {
			bigTemp.Next = dumbNode
			bigTemp = bigTemp.Next
		}
		dumbNode = dumbNode.Next
	}
	//链接两个链表
	bigTemp.Next = nil
	smallTemp.Next = bigDumbNode.Next
	//拼接
	return smallDumbNode.Next
}
