package main

import "fmt"

func removeElement(nums []int, val int) int {
	k := 0 // are not equal to val

	for i := range nums {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	removed_elements := removeElement(nums, val)

	fmt.Println(removed_elements)
	fmt.Println(nums)
}
