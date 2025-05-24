package main

import "fmt"

func maxProfit(prices []int) int {
	buyProfitDay := 0
	sellProfitDay := 0
	profit := 0

	for i := range prices {
		if prices[i] < prices[buyProfitDay] {
			buyProfitDay = i
			sellProfitDay = i
		} else if prices[i] > prices[sellProfitDay] {
			sellProfitDay = i
		}
		if prices[sellProfitDay]-prices[buyProfitDay] > profit {
			profit = prices[sellProfitDay] - prices[buyProfitDay]
		}
	}

	return profit
}

func main() {
	input := []int{7, 1, 5, 3, 6, 4}
	output := maxProfit(input)

	fmt.Println(output) // Expected: 5
}
