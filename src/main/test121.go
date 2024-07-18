package main

func main() {

}

// 力扣121简单题 买卖股票最佳时机1
// 方法1 动态规划
func maxProfit(prices []int) int {
	//每个细分阶段 要想利润最大，则一定是最小点买入的
	min := prices[0]
	//利润
	profit := 0
	for _, price := range prices {
		if price < min {
			min = price
		}
		if price-min > profit {
			profit = price - min
		}
	}
	return profit
}
