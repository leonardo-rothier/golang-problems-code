package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	k := 2

	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[k-2] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

func main() {
	input := []int{1, 1, 1, 2, 2, 3}

	k := removeDuplicates(input)
	output := input[:k]

	fmt.Println(output) // Expected: [1 1 2 2 3]
}
