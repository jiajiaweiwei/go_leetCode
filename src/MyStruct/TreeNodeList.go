package MyStruct

// TreeNodeList 每一层构成链表的二叉树（每个节点多了个next指针）
type TreeNodeList struct {
	Val   int
	Left  *TreeNodeList
	Right *TreeNodeList
	Next  *TreeNodeList
}
