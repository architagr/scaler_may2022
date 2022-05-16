package isitprime

import (
	"math"
)

func IsPrime(n int64) bool {
	sqrt := int64(math.Sqrt(float64(n)))
	var i int64 = 1
	count := 0
	for ; i <= sqrt; i++ {
		if i*i == n {
			count++
		} else if n%i == 0 {
			count += 2
		}
	}
	return count == 2
}
