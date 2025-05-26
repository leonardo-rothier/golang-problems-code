package main

import "fmt"

func jump(nums []int) int {
	n := len(nums)
	currentEnd := 0
	farthest := 0

	if n <= 1 {
		return 0
	}
	jumps := 0

	for i := range nums {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}

		if currentEnd == i {
			jumps++

			currentEnd = farthest

			if currentEnd >= n-1 {
				break
			}
		}
	}

	return jumps
}

func main() {
	input := []int{2, 3, 1, 1, 4}

	fmt.Println(jump(input)) // Expected: 2
}
