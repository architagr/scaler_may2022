package squarerootofanumber

import (
	"fmt"
	"testing"
)

func TestSquareRoot(t *testing.T) {
	testCases := make(map[int]int)
	// testCases[4] = 2
	// testCases[1001] = -1
	// testCases[121] = 11
	// testCases[121] = 11
	// testCases[214358881] = 14641
	// testCases[214358896] = -1
	// testCases[43046721] = 6561
	// testCases[3481] = 59
	testCases[2304] = 48

	for key, expected := range testCases {
		t.Run(fmt.Sprintf("Testing %d", key), func(tb *testing.T) {
			got := SquareRoot(key)
			if got != expected {
				tb.Errorf("Tested %d, expected %d, got %d", key, expected, got)
			}
		})
	}
}
