package main

func main() {

}

// 力扣169 简单题 多数元素
// 方法1 哈希表法
func majorityElement(nums []int) int {
	//特殊情况处理 只有一个元素时
	if len(nums) == 1 {
		return nums[0]
	}

	l := len(nums) / 2
	m := make(map[int]int)
	for _, value := range nums {
		//如果键已经存在，说明重复
		if _, ok := m[value]; ok {
			m[value]++
			//当发现第一个元素的个数就已经够了时，直接返回
			if m[value] > l {
				return value
			}
		} else {
			//不存在就放入map中
			m[value] = 1
		}
	}
	//异常
	return 0
}

// 方法2 摩尔——投票法 维护众数
//摩尔投票算法的核心
/*我们维护一个候选众数 candidate 和它出现的次数 count。初始时 candidate 可以为任意值，count 为 0；

我们遍历数组 nums 中的所有元素，对于每个元素 x，在判断 x 之前，如果 count 的值为 0，我们先将 x 的值赋予 candidate，随后我们判断 x：

如果 x 与 candidate 相等，那么计数器 count 的值增加 1；

如果 x 与 candidate 不等，那么计数器 count 的值减少 1。*/

func majorityElement1(nums []int) int {
	//特殊情况处理 只有一个元素时
	if len(nums) == 1 {
		return nums[0]
	}
	//实现摩尔——投票算法
	var candidate, count int
	for _, num := range nums {
		if count == 0 {
			candidate = num
			count++
		} else if num != candidate {
			count--
		} else {
			count++
		}

	}
	return candidate
}
