package main

import "fmt"

func removeDuplicates(nums []int) int {
	k := 1 // number of unique elements (the first is already unique)

	last_num := nums[0]
	for _, num := range nums {
		if num != last_num {
			nums[k] = num
			k++
		}
		last_num = num
	}

	return k
}

func main() {
	input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	var k int = removeDuplicates(input)

	output := input[:k]

	fmt.Println(output) // Expected [0,1,2,3,4]
}
