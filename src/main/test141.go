package main

import (
	. "MyLikou/src/MyStruct"
	"fmt"
)

func main() {
	fmt.Println("----")
}

// 链表 简单题 回环链表 判断某个单链表中是否有环
// 方法1 快慢指针 如果存在换，快指针总会追上慢指针 (注意索引不表示节点的位置)
func hasCycle(head *ListNode) bool {
	slow, fast := head, head.Next
	//特殊情况处理
	if slow == nil || fast == nil {
		return false
	}
	for fast != slow {
		//先判断 防止空指针异常
		if fast == nil || fast.Next == nil || fast.Next.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}

//方法2 哈希表 直接用哈希表存储节点 判断后继节点是否在前面已经出现 （不推荐）
