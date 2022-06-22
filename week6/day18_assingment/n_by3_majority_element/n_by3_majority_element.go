package n_by3_majority_element

func MajorityElementNby3(nums []int) int {
	me1, f1, me2, f2, n := -1, 0, -1, 0, len(nums)
	//sort.Ints(nums)

	me1 = nums[0]
	f1 = 1

	for i := 1; i < n; i++ {
		if nums[i] == me1 {
			f1++
		} else if nums[i] == me2 {
			f2++
		} else if f1 == 0 {
			me1 = nums[i]
			f1 = 1
		} else if f2 == 0 {
			me2 = nums[i]
			f2 = 1
		} else {
			f2--
			f1--
		}
	}
	if f1 > 0 {
		count := getCount(nums, me1)
		if count > n/3 {
			return me1
		}
	}
	if f2 > 0 {
		count := getCount(nums, me2)
		if count > n/3 {
			return me2
		}
	}
	return -1
}

func getCount(A []int, nm int) int {
	count := 0
	for i := 0; i < len(A); i++ {
		if A[i] == nm {
			count++
		}
	}
	return count
}
