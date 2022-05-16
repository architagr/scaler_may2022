package summationgame

import (
	"fmt"
	"testing"
)

func TestSummationGame(t *testing.T) {
	testCases := make(map[int]int)
	testCases[5] = 15
	testCases[15] = 120

	for key, expected := range testCases {
		t.Run(fmt.Sprintf("testing for %d", key), func(tb *testing.T) {
			got := SummationGame(key)
			if got != expected {
				tb.Errorf("tested for %d expected %d and got %d", key, expected, got)
			}
		})
	}
}
