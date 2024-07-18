package main

import "fmt"

func main() {

	gas := []int{5, 1, 2, 3, 4}
	cost := []int{4, 4, 1, 5, 1}
	circuit := canCompleteCircuit(gas, cost)
	fmt.Println(circuit)
}

// 一半用例未通过
func canCompleteCircuit1(gas []int, cost []int) int {
	//总长度
	n := len(gas)
	//合适的起点
	start := []int{}
	//到达某点的剩余油量
	other := 0
	//找到符合条件的起点
	for i := 0; i < n; i++ {
		if gas[i] >= cost[i] {
			start = append(start, i) //起点可能不止有一个
			other = gas[i] - cost[i] //刚跑完i还未加油时的剩余油量
		}
	}
	for i := 0; i < len(start); i++ {
		//判断是否存在这样的起点，不存在则返回-1
		if start[i] == -1 {
			return start[i]
		}
		//开始遍历，查看是否能遍历回起点
		for j := 1; j < n; j++ {
			next := (start[i] + j) % n //重点，当起点不在索引0时的遍历一轮的方式 取余
			other += gas[next] - cost[next]
			if other < 0 {
				if i == len(start)-1 {
					return -1
				}
			}
			if j == n-1 {
				return start[i]
			}
		}
	}
	return -1

}

// 力扣 134 中等题 加油站 正确答案
func canCompleteCircuit(gas []int, cost []int) int {
	//遍历加油站
	for i, n := 0, len(gas); i < n; {
		// 剩余量综合，消耗总和 计数
		sumOfGas, sumOfCost, cnt := 0, 0, 0
		for cnt < n {
			j := (i + cnt) % n   //i表示起点加油站索引 cnt表示计数器
			sumOfGas += gas[j]   //累加当前加油站的汽油量
			sumOfCost += cost[j] //累加当前加油站到下一站的消耗
			//消耗大于汽油量，无法到达下一站，退出循环
			if sumOfCost > sumOfGas {
				break
			}
			cnt++
		}
		if cnt == n { //cnt==n 说明成功绕行了一圈
			return i
		} else {
			//否则更新起点索引，跳过无法到达的加油站
			i += cnt + 1
		}
	}
	//无法绕行，返回-1
	return -1
}

func test(gas []int, cost []int) int {
	//测试加油站
	for i := 0; i < len(gas); {
		//总油量，耗油量 计数器
		sumGas, sumCost, cnt := 0, 0, 0
		//循环条件，计数器小于长度
		for cnt < len(gas) {
			//索引 内层遍历的索引 起点i+内层遍历计数器cnt
			j := (i + cnt) % len(gas)
			sumGas += gas[j]
			sumCost += cost[j]
			if sumGas < sumCost { //油量不够  直接进入下层循环
				break
			}
			cnt++ //油量够，能到达下一站，继续判断下一站
		}
		//如果cnt== 长度说明已经有一轮循环了
		if cnt == len(gas) {
			return i
		} else {
			//否则判断下一轮 其中中间论可以省略
			i += cnt + 1
		}
	}
	//如果外层遍历完了也不行，说明不存在返回-1
	return -1
}
