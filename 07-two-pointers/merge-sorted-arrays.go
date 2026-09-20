package twopointers

import "slices"

func SlicesMergeSortedArrays(nums1 []int, nums2 []int) []int {
	slices.Concat(nums1, nums2)
	slices.Sort(nums1)
	return nums1
}

// LeetCode 88. Merge Sorted Array
func MergeSortedArrays(nums1 []int, nums2 []int, n int, m int) []int {
	last := n + m - 1

	for m > 0 && n > 0 {
		if nums1[n-1] < nums2[m-1] {
			nums1[last] = nums2[m-1]
			m--
		} else {
			nums1[last] = nums1[n-1]
			n--
		}
		last--
	}

	for m > 0 {
		nums1[last] = nums2[m-1]
		m--
		last--
	}

	return nums1
}
