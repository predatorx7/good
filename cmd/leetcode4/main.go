package main

import (
	"fmt"
	"good/pkg/leetcode4"
)

func display(nums1 []int, nums2 []int, expected float64) {
	r := leetcode4.Solve(nums1, nums2)
	fmt.Printf("expected = %f, result = %f\n", expected, r)
}

func main() {
	display([]int{1, 2}, []int{3, 4}, 2.5)
	display([]int{1, 3}, []int{2, 7}, 2.5)
	display([]int{}, []int{1}, 1)
	display([]int{1, 3}, []int{2}, 2)
	display([]int{2, 3, 4, 5}, []int{1}, 3)
	display([]int{2, 3, 4, 5, 6, 7}, []int{1}, 4)
	display([]int{2, 3, 4, 5, 6, 7, 8, 9}, []int{1}, 5)
}
