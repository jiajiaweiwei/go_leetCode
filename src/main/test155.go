package main

import "math"

// MinStack
// Your MinStack object will be instantiated and called as such:
// obj := Constructor();
// obj.Push(val);
// obj.Pop();
// param_3 := obj.Top();
// param_4 := obj.GetMin();
// 栈 中等题 最小栈
// 方法 使用辅助栈
type MinStack struct {
	// 栈数据
	stack []int
	// 辅助栈
	tempStack []int
}

// Constructor155 初始化栈
func Constructor155() MinStack {
	return MinStack{
		[]int{},
		[]int{math.MaxInt64},
	}
}

// Push 将元素val入栈
func (this *MinStack) Push(val int) {
	// 将元素入栈
	this.stack = append(this.stack, val)
	// 更新辅助栈
	if val < this.tempStack[len(this.tempStack)-1] {
		this.tempStack = append(this.tempStack, val)
	} else {
		this.tempStack = append(this.tempStack, this.tempStack[len(this.tempStack)-1])
	}
}

// Pop 删除栈顶元素
func (this *MinStack) Pop() {
	// 删除栈顶元素
	this.stack = this.stack[:len(this.stack)-1]
	// 删除辅助栈栈顶
	this.tempStack = this.tempStack[:len(this.tempStack)-1]
}

// Top 获取栈顶元素 但不出栈
func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

// GetMin 获取栈的最小值
func (this *MinStack) GetMin() int {
	return this.tempStack[len(this.tempStack)-1]
}
