package isitprime

import (
	"fmt"
	"testing"
)

func TestIsPrime(t *testing.T) {
	testCases := make(map[int64]bool)
	testCases[3] = true
	testCases[4] = false
	testCases[113] = true

	for key, testcase := range testCases {
		t.Run(fmt.Sprintf("testing %d", key), func(tb *testing.T) {
			got := IsPrime(key)
			if got != testcase {
				tb.Errorf("Tested %d expected %t but got %t", key, testcase, got)
			}
		})
	}
}
