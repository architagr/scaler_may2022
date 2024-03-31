package pickfrombothsides

import (
	"fmt"
	"testing"
)

type testCase struct {
	inputA   []int
	inputB   int
	expected int
}

func TestLogic(t *testing.T) {

	testCases := []testCase{
		{
			inputA:   []int{5, -2, 3, 1, 2},
			inputB:   3,
			expected: 8,
		},
		{
			inputA:   []int{2, 3, -1, 4, 2, 1},
			inputB:   4,
			expected: 9,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i), func(tb *testing.T) {
			tb.Logf("run %d", i)
			if got := solve(tc.inputA, tc.inputB); got != tc.expected {
				tb.Fatalf("expected %d, got %d, inputA %+v, inputB %d", tc.expected, got, tc.inputA, tc.inputB)
			}
		})
	}
}
