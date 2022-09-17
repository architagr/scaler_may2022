package three_sum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	type TestCase struct {
		inputA      []int
		b, expected int
	}
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   []int{-1, 2, 1, -4},
		b:        1,
		expected: 2,
	})

	testCases = append(testCases, TestCase{
		inputA:   []int{-5, 1, 4, -7, 10, -7, 0, 7, 3, 0, -2, -5, -3, -6, 4, -7, -8, 0, 4, 9, 4, 1, -8, -6, -6, 0, -9, 5, 3, -9, -5, -9, 6, 3, 8, -10, 1, -2, 2, 1, -9, 2, -3, 9, 9, -10, 0, -9, -2, 7, 0, -4, -3, 1, 6, -3},
		b:        -1,
		expected: -1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := ThreeSumClosest(testCase.inputA, testCase.b)
			if !reflect.DeepEqual(testCase.expected, got) {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
