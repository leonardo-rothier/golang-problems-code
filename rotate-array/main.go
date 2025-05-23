package main

import "fmt"

func rotate(nums []int, k int) {
	times := len(nums) - (k % len(nums))
	copy(nums, append(nums[times:], nums[:times]...))
}

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	k := 7
	rotate(input, k)
	fmt.Println(input)
}
