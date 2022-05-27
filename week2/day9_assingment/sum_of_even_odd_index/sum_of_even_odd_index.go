package sum_of_even_odd_index

func SumOfEvenOddIndex(A []int) int {
	evenPrefixArray := FindEvenPrefixSum(A)
	oddPrefixArray := FindOddPrefixSum(A)
	count := 0
	var even int64 = 0
	var odd int64 = 0

	for i := 0; i < len(A); i++ {
		if i == 0 {
			even = oddPrefixArray[len(A)-1]
			odd = evenPrefixArray[len(A)-1] - evenPrefixArray[i]
		} else {
			even = evenPrefixArray[i-1] + (oddPrefixArray[len(A)-1] - oddPrefixArray[i])
			odd = oddPrefixArray[i-1] + (evenPrefixArray[len(A)-1] - evenPrefixArray[i])
		}
		if even == odd {
			count++
		}
	}
	return count
}

func FindEvenPrefixSum(A []int) []int64 {

	result := make([]int64, len(A))
	var sum int64 = 0
	for i, val := range A {
		if i%2 != 0 {
			result[i] = result[i-1]
		} else {
			sum += int64(val)
			result[i] = sum
		}
	}
	return result
}

func FindOddPrefixSum(A []int) []int64 {
	result := make([]int64, len(A))
	var sum int64 = 0
	for i, val := range A {
		if i%2 == 0 {
			if i == 0 {
				result[i] = 0
			} else {
				result[i] = result[i-1]
			}
		} else {
			sum += int64(val)
			result[i] = sum
		}
	}
	return result
}
