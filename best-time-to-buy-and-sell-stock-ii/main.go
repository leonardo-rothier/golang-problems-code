package main

import "fmt"

func maxProfit(prices []int) int {
	profit := 0

	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1]-prices[i] > 0 {
			profit += prices[i+1] - prices[i]
		}
	}
	return profit
}

func main() {
	input := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(input))
}
