package main

import . "MyLikou/src/MyStruct"

func main() {

}

/*



啥玩意、题目没看懂



*/

// BSTIterator 二叉树 中等题 实现二叉搜索树的升序迭代器
// 迭代器结构体
type BSTIterator struct {
}

// Constructor173 构造器
func Constructor173(root *TreeNode) BSTIterator {

	return BSTIterator{}
}

func (this *BSTIterator) Next() int {

	return 0
}

func (this *BSTIterator) HasNext() bool {
	return true
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
