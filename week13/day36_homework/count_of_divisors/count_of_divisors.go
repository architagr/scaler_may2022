package count_of_divisors

func CountDivisors(A []int) []int {
	max := A[0]

	//Calculate max to initialize divisor array of max length
	for i := 1; i < len(A); i++ {
		if A[i] > max {
			max = A[i]
		}
	}
	max++
	divisor_array := make([]int, max)
	divisor_array[1] = 1 //Setting number of divisor of 1 as 1

	//Setting initial divisors as 2 as all numbers will have at least 2 divisors i.e. 1 and itself
	for j := 2; j < max; j++ {
		divisor_array[j] = 2
	}

	//incrementing the number of divisors for each multiples of numbers
	for k := 2; k < max; k++ {

		for l := k + k; l < max; l += k {
			divisor_array[l] += 1
		}
	}

	//Fetching the values of numbers in array directly from divisor_array into the output array
	for c := 0; c < len(A); c++ {
		A[c] = divisor_array[A[c]]
	}

	return A

}
