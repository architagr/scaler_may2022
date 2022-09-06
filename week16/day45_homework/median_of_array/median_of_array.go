package median_of_array

import "math"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	n, m := len(nums1), len(nums2)
	if n > m {
		return FindMedianSortedArrays(nums2, nums1)
	}
	low, high := 0, n
	for low <= high {
		cut1 := (low + high) / 2
		cut2 := ((n + m + 1) / 2) - cut1

		left1, left2 := math.MinInt32, math.MinInt32

		if cut1 > 0 {
			left1 = nums1[cut1-1]
		}

		if cut2 > 0 {
			left2 = nums2[cut2-1]
		}

		right1, right2 := math.MaxInt32, math.MaxInt32

		if cut1 != n {
			right1 = nums1[cut1]
		}
		if cut2 != m {
			right2 = nums2[cut2]
		}

		if left1 <= right2 && left2 <= right1 {
			if (n+m)%2 == 0 {
				return float64((maxValue(left1, left2))+minValue(right1, right2)) / 2.0
			} else {
				return float64((maxValue(left1, left2)))
			}
		}
		if left1 > right2 {
			high = cut1 - 1
		}
		if left2 > right1 {
			low = cut1 + 1
		}
	}
	return 0.0
}
func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
