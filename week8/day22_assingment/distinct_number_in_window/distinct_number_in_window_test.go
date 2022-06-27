package distinct_number_in_window

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA   []int
	inputB   int
	expected []int
}

func TestDistinctNumbers(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   []int{1, 2, 1, 3, 4, 3},
		inputB:   3,
		expected: []int{2, 3, 3, 2},
	})
	testCases = append(testCases, TestCase{
		inputA:   []int{1, 1, 2, 2},
		inputB:   1,
		expected: []int{1, 1, 1, 1},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := DistinctNumbers(testCase.inputA, testCase.inputB)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}

}
