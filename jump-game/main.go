package main

import "fmt"

func canJump(nums []int) bool {
	lastPos := len(nums) - 1

	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}

	return lastPos == 0
}

func main() {
	input := []int{2, 5, 0, 0}
	fmt.Println(canJump(input)) // Expected: true
}
