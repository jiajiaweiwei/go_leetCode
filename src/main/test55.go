package main

// 力扣55 中等题 跳跃游戏
func canJump(nums []int) bool {
	max := 0
	for i := 0; i < len(nums); i++ {
		if i > max {
			return false
		}
		if max < nums[i]+i {
			max = i + nums[i]
		}

	}
	return true
}

/*如果某一个作为 起跳点 的格子可以跳跃的距离是 3，那么表示后面 3 个格子都可以作为 起跳点
可以对每一个能作为 起跳点 的格子都尝试跳一次，把 能跳到最远的距离 不断更新
如果可以一直跳到最后，就成功了*/
