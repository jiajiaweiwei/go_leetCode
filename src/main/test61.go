package main

import . "MyLikou/src/MyStruct"

func main() {

}

// 链表中等题旋转链表
// 方法 1 使链表闭合为环，移动头指针即可，然后再断开 断开倒数第k个位置 即可
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	temp := head
	last := head
	//链表长度
	l := 0
	//先找到最后一个节点闭合为环
	for temp != nil {
		l++
		last = temp
		temp = temp.Next
	}
	//特殊情况处理 旋转前后结果不变，直接返回
	if k%l == 0 {
		return head
	}
	k = k % l
	//最后一个节点是last
	//闭合成环
	last.Next = head
	//找到要断开的位置 就是遍历 l-k-1次的位置
	temp2 := head
	for i := 0; i < l-k-1; i++ {
		temp2 = temp2.Next
	}
	resultNode := temp2.Next
	temp2.Next = nil
	return resultNode
}
