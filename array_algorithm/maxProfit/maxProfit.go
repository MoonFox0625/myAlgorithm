// Package maxProfit :
// File:  maxProfit.go
// Date:  2021/1/19 11:12
// Desc:
package maxProfit

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	amount := make([][2]int, len(prices)) // 当天是否决定卖出、买入后的能获得的最大总额
	amount[0][0] = 0
	amount[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		amount[i][0] = max(amount[i-1][1]+prices[i], amount[i-1][0])
		amount[i][1] = max(amount[i-1][0]-prices[i], amount[i-1][1])
	}

	return amount[len(amount)-1][0]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
