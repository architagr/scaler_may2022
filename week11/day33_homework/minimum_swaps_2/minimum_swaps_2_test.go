package minimum_swaps_2

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input       []int
	b, expected int
}

func TestMinSwaps(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{1, 12, 10, 3, 14, 10, 5},
		b:        8,
		expected: 2,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := MinSwaps(testCase.input, testCase.b)

			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
