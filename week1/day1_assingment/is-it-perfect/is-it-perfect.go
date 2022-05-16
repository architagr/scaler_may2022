package isitperfect

import (
	"math"
)

func IsItPerfect(n int64) bool {
	sqrt := int64(math.Sqrt(float64(n)))
	var i, sum int64 = 2, 1

	for ; i <= sqrt; i++ {
		if i*i == n {
			sum += i
		} else if n%i == 0 {
			sum += i
			sum += n / i
		}
	}
	return sum == n
}
