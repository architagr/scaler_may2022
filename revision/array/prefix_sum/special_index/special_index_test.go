package specialindex

import (
	"fmt"
	"testing"
)

type testCase struct {
	input    []int
	expected int
}

func TestLogic(t *testing.T) {
	testCases := []testCase{
		{
			input:    []int{2, 1, 6, 4},
			expected: 1,
		},
		{
			input:    []int{1, 1, 1},
			expected: 3,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i), func(tb *testing.T) {
			tb.Logf("run %d", i)
			if got := solve(tc.input); got != tc.expected {
				tb.Fatalf("expected %d, got %d, input %+v", tc.expected, got, tc.input)
			}
		})
	}
}
