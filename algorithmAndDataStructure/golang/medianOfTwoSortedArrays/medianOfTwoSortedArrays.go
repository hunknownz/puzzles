package medianOfTwoSortedArrays

import (
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	if total&1 == 0 {
		return (findKthValue(nums1, nums2, total/2) + findKthValue(nums1, nums2, total/2+1)) / 2
	} else {
		return findKthValue(nums1, nums2, total/2+1)
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func findKthValue(nums1 []int, nums2 []int, k int) float64 {
	if len(nums1) == 0 {
		return float64(nums2[k-1])
	}
	if len(nums2) == 0 {
		return float64(nums1[k-1])
	}
	if k == 1 {
		return math.Min(float64(nums1[0]), float64(nums2[0]))
	}

	mid1, mid2 := min(len(nums1), k/2), min(len(nums2), k/2)
	if nums1[mid1-1] < nums2[mid2-1] {
		return findKthValue(nums1[mid1:], nums2, k-mid1)
	} else {
		return findKthValue(nums1, nums2[mid2:], k-mid2)
	}
}
