package distinct_numbers_in_window

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDistinctNumberInWindow(t *testing.T) {
	type TestCase struct {
		input    []int
		B        int
		expected []int
	}
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    []int{1, 2, 1, 3, 4, 3},
		B:        3,
		expected: []int{2, 3, 3, 2},
	})

	testCases = append(testCases, TestCase{
		input:    []int{1, 1, 1, 1, 1, 1},
		B:        3,
		expected: []int{1, 1, 1, 1},
	})

	testCases = append(testCases, TestCase{
		input:    []int{1, 2, 1, 1, 1, 1},
		B:        3,
		expected: []int{2, 2, 1, 1},
	})

	testCases = append(testCases, TestCase{
		input:    []int{1, 1, 2, 2},
		B:        1,
		expected: []int{1, 1, 1, 1},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := DistinctNumberInWindow(testCase.input, testCase.B)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
