package main

import . "MyLikou/src/MyStruct"

func main() {

}

// 中等题 实现插入排序
// head = [4,2,1,3]
// 输出: [1,2,3,4]
func insertionSortList(head *ListNode) *ListNode {
	// 异常处理情况
	if head == nil {
		return nil
	}
	// 哑节点
	dummyHead := &ListNode{Next: head}
	// 最后一个节点和临时节点
	lastSorted, curr := head, head.Next
	for curr != nil {
		// 找到第一个需要移动的节点位置
		if lastSorted.Val <= curr.Val {
			lastSorted = lastSorted.Next
		} else {
			// 找到第一个需要移动的节点位置的前一个位置
			prev := dummyHead
			for prev.Next.Val <= curr.Val {
				prev = prev.Next
			}
			// 开始移动元素
			lastSorted.Next = curr.Next
			curr.Next = prev.Next
			prev.Next = curr
		}
		// 更新临时节点的位置
		curr = lastSorted.Next
	}
	return dummyHead.Next
}
