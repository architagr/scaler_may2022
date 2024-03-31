package productarraypuzzle

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	input, expected []int
}

func TestLogic(t *testing.T) {
	testCases := []testCase{
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{120, 60, 40, 30, 24},
		},
		{
			input:    []int{5, 1, 10, 1},
			expected: []int{10, 50, 5, 50},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint(i), func(tb *testing.T) {
			tb.Logf("run %d", i)
			if got := solve(tc.input); !reflect.DeepEqual(got, tc.expected) {
				tb.Fatalf("expected %+v, got %+v, input %+v", tc.expected, got, tc.input)
			}
		})
	}
}
