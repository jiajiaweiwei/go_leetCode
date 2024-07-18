package main

import . "MyLikou/src/MyStruct"

func main() {

}

// 链表 中等题  链表数据的两数相加 返回值 也是链表结构
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 创建一个哑结点作为结果链表的头结点
	//头节点的写一个节点为返回值的最低位
	dummy := &ListNode{}
	// 用于遍历结果链表
	current := dummy
	carry := 0                                // 进位
	for l1 != nil || l2 != nil || carry > 0 { //进位大于0 也要进行下一轮循环
		//某一位的总和
		sum := carry
		//先加l1 指针后移
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		//再加l2 指针后移
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		// 计算当前位的值和进位
		current.Next = &ListNode{Val: sum % 10} //初始化一个节点 并添加到结果节点上
		carry = sum / 10                        //进位 给下一级使用
		// 移动到下一个节点
		current = current.Next
	}
	return dummy.Next // 返回结果链表的头结点的下一个节点
}

// 方法2 分三种情况 l1为空l2不为空   l2为空l1不为空  l1l2都不为空
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	//返回值
	reusltHead := new(ListNode)
	current := reusltHead //指针遍历
	carry := 0            //进位
	//分三种情况
	//1.两者都不为空
	for l1 != nil && l2 != nil {
		sum := (l1.Val + l2.Val + carry) % 10
		carry = (l1.Val + l2.Val + carry) / 10 //进位
		current.Next = &ListNode{Val: sum}
		current = current.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	//2.只有l1为空
	for l1 != nil {
		sum := (l1.Val + carry) % 10
		carry = (l1.Val + carry) / 10 //进位
		current.Next = &ListNode{Val: sum}
		current = current.Next
		l1 = l1.Next
	}
	//3.只有l2为空
	for l2 != nil {
		sum := (l2.Val + carry) % 10
		carry = (l2.Val + carry) / 10 //进位
		current.Next = &ListNode{Val: sum}
		current = current.Next
		l2 = l2.Next
	}
	//3.都为空但 还剩个进位
	if carry != 0 {
		current.Next = &ListNode{Val: carry}
		current = current.Next
	}
	return reusltHead.Next
}
