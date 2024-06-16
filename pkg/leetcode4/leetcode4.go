package leetcode4

import "slices"

func Solve(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)

	r := make([]int, n)

	filled := copy(r, nums1)

	for i, e := range nums2 {
		r[filled+i] = e
	}

	slices.Sort(r)

	if n%2 == 0 {
		center := (n / 2) + 1
		return float64(r[center-1]+r[center-2]) / 2
	} else {
		center := (n + 1) / 2
		return float64(r[center-1])
	}
}
