package leetcode4

import (
	"slices"
)

func Solve(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)

	r := make([]int, n)
	i := 0
	for _, e := range nums1 {
		r[i] = e
		i++
	}
	for _, e := range nums2 {
		r[i] = e
		i++
	}

	// TODO: Check if partial sort is possible
	slices.Sort(r)

	if n%2 == 0 {
		center := (n / 2) + 1
		return (float64(r[center-1]) + float64(r[center-2])) / 2
	} else {
		center := (n + 1) / 2
		return float64(r[center-1])
	}
}
