package main

import . "MyLikou/src/MyStruct"

func main() {

}

// 分而治之，简单题 升序数组 转化成平衡二叉搜索树
// 二叉搜索树的中序遍历就是 该升序数组
func sortedArrayToBST(nums []int) *TreeNode {
	// 执行递归函数
	return sortToBSTTemp(nums, 0, len(nums)-1)
}

// 建立 二叉搜索树递归函数
func sortToBSTTemp(num []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	//先从根节点开始建立
	// 选取中间位置左边作为中间节点
	mid := (left + right) / 2
	root := &TreeNode{
		Val: num[mid],
	}
	root.Left = sortToBSTTemp(num, left, mid-1)
	root.Right = sortToBSTTemp(num, mid+1, right)
	return root
}
