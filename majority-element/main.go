package main

import "fmt"

func majorityElement(nums []int) int {
	counts := make(map[int]int)
	majorityElement := nums[0]
	maxCount := 0

	for _, num := range nums {
		counts[num]++
		if counts[num] > maxCount {
			maxCount = counts[num]
			majorityElement = num
		}
	}

	return majorityElement
}

func main() {
	var input []int = []int{2, 2, 1, 1, 1, 2, 2}

	fmt.Println(majorityElement(input)) // Expeted output: 2
}
