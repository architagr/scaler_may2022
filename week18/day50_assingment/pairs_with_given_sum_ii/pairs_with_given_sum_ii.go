package pairs_with_given_sum_ii

func Sum2II(A []int, B int) int {
	MOD := 1000000007
	count := 0
	i, j := 0, len(A)-1
	for i < j {
		sum := A[i] + A[j]
		if sum < B {
			i++
		} else if sum > B {
			j--
		} else {
			if A[i] == A[j] {
				n := j - i + 1
				count = (count + ((n*(n-1))/2)%MOD) % MOD
				break
			} else {
				x, y := 1, 1

				for i+x < len(A) && A[i] == A[i+x] {
					x++
				}
				for j-y >= 0 && A[j] == A[j-y] {
					y++
				}
				i += x
				j -= y

				count = (count + (x*y)%MOD) % MOD
			}
		}
	}
	return count % MOD
}
