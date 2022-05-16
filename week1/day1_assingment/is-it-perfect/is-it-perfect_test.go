package isitperfect

import (
	"fmt"
	"testing"
)

func TestIsPerfect(t *testing.T) {
	testCases := make(map[int64]bool)
	testCases[3] = false
	testCases[4] = false
	testCases[6] = true

	for key, testcase := range testCases {
		t.Run(fmt.Sprintf("testing %d", key), func(tb *testing.T) {
			got := IsItPerfect(key)
			if got != testcase {
				tb.Errorf("Tested %d expected %t but got %t", key, testcase, got)
			}
		})
	}
}
