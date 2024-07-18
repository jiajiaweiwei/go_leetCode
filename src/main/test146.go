package main

func main() {

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
/*
type ListNodeTwo struct {
	Val  int
	Pre  *ListNodeTwo
	Next *ListNodeTwo
}

// 链表中等题 设计 页面置换算法： 最近最少使用LRU算法 约束的数据结构 使用双向链表+map实现
type LRUCache struct {
	Size     int                  //已经包含的节点数
	Capacity int                  //总容量
	Cache    map[int]*ListNodeTwo //元素（cache页缓存） key:数字 value :节点地址
	DumpNode ListNodeTwo          //链表头哑节点
}

// Constructor 初始化容量 并返回一个 LRUCache
func Constructor(capacity int) LRUCache {
	return LRUCache{
		Size:     0,
		Capacity: capacity,
		DumpNode: ListNodeTwo{},
	}
}

// Get 获取到某元素之后，如果获取到 将该元素移动到链表头
// 获取元素 是直接通过key value 的形式，故此处要用到map
func (this *LRUCache) Get(key int) int {
	//拿到LRUCache的哑节点（指向第一个节点）
	dumpNode := this.DumpNode
	//如果有该节点 将它移动到双向链表头 并返回对应节点的内容
	if getNode, ok := this.Cache[key]; !ok {
		return -1
	} else {
		//处理 getNode 临近节点
		// 1.该节点不是末尾节点的情况
		if getNode.Next != nil {
			getNode.Pre.Next = getNode.Next
			getNode.Next.Pre = getNode.Pre
		} else {
			// 该节点是末尾的情况
			getNode.Pre.Next = nil
		}
		//移动到链头
		getNode.Next = dumpNode.Next
		dumpNode.Next.Pre = getNode
		//返回
		return getNode.Val
	}
}

// Put 添加某元素之后，将该元素移动到表头
// 如果容量不够 不仅要将改元素移动到链表头，还要删除链表尾节点
func (this *LRUCache) Put(key int, value int) {
	//key不存在时
	if _, ok := this.Cache[key]; !ok {
		firstNode := this.DumpNode.Next
		putNode := &ListNodeTwo{
			Val:  value,
			Pre:  &this.DumpNode,
			Next: firstNode,
		}
		// 直接插入头
		firstNode.Pre = putNode
		putNode.Next = firstNode
		this.DumpNode.Next = putNode
		//添加到map
		this.Cache[key] = putNode
		//判断容量是否够
		if this.Size == this.Capacity {
			// 不够 还要删除链表尾节点
			// TODO
			tempNode := this.DumpNode
			for tempNode.Next != nil {
				tempNode = *tempNode.Next
			}
			tempNode.Pre.Next = nil
			tempNode.Pre = nil
			this.Size--
		}
		this.Size++
	} else { //key存在时 更新value 并移动到链表头即可
		tempNode := this.Cache[key].Next
		this.Cache[key].Val = value
		//移动到链表头
		firstNode := this.DumpNode.Next
		firstNode.Pre = this.Cache[key]
		this.DumpNode.Next = this.Cache[key]
		this.Cache[key].Pre = &this.DumpNode
		this.Cache[key].Next = firstNode
		//不是最后一个节点时
		if tempNode != nil {
			tempNode.Pre.Next = tempNode.Next
			tempNode.Next.Pre = tempNode.Pre
		} else {
			//是最后一个节点时
			tempNode.Pre.Next = nil
		}

	}

}*/
