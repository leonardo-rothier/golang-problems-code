package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	idx1, idx2 := 0, 0

	merged := make([]int, 0, m+n)

	for idx1 < m && idx2 < n {
		if nums1[idx1] <= nums2[idx2] {
			merged = append(merged, nums1[idx1])
			idx1++
		} else {
			merged = append(merged, nums2[idx2])
			idx2++
		}
	}

	// append remaning elements
	merged = append(merged, nums1[idx1:m]...)
	merged = append(merged, nums2[idx2:]...)

	copy(nums1, merged)

}

func main() {
	input_a := []int{1, 2, 3, 0, 0, 0}
	input_b := []int{2, 5, 6}

	merge(input_a, 3, input_b, 3)

	fmt.Println(input_a)
}
