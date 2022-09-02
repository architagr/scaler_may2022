package median_of_array

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n, m := len(nums1), len(nums2)
	isOdd := true
	if m > n {
		n, m = m, n
		nums1, nums2 = nums2, nums1
	}
	if (n+m)&1 == 0 {
		isOdd = false
	}
	length := (n + m) / 2

	arr := make([]int, 0, length+1)
	k := 0
	arr = append(arr, nums1...)

	for k < m && ((len(arr) > length && nums2[k] < arr[length]) || (len(arr) <= length)) {
		l, r := 0, len(arr)-1
		for l <= r {
			mid := (l + r) / 2
			if nums2[k] < arr[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		arr = append(arr[:l], append([]int{nums2[k]}, arr[l:]...)...)
		k++
	}
	if isOdd {
		return float64(arr[length])
	}
	return (float64(arr[length]) + float64(arr[length-1])) / 2.0
}
