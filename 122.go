package main

func maxProfit2(prices []int) int {
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if res := prices[i] - prices[i-1]; res > 0 {
			maxProfit += res
		}
	}
	return maxProfit
}
