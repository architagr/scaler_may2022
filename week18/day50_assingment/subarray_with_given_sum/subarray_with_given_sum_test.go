package subarray_with_given_sum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAggressiveCows(t *testing.T) {
	type TestCase struct {
		input    []int
		b        int
		expected []int
	}
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    []int{1, 2, 3, 4, 5},
		b:        5,
		expected: []int{2, 3},
	})

	testCases = append(testCases, TestCase{
		input:    []int{5, 10, 20, 100, 105},
		b:        110,
		expected: []int{-1},
	})

	testCases = append(testCases, TestCase{
		input:    []int{12, 1, 50, 39, 32, 23, 22, 44, 17, 5, 9, 12, 10, 50, 26, 5, 23, 38, 31, 5, 34, 8, 21, 11, 24, 44, 18, 19, 6, 31, 3, 47, 5, 2, 33, 44, 14, 9},
		b:        58,
		expected: []int{44, 14},
	})

	testCases = append(testCases, TestCase{
		input:    []int{1, 1000000000},
		b:        1000000000,
		expected: []int{1000000000},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := SubarrayWithGivenSum(testCase.input, testCase.b)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("Tested # %d, where we expected %+v but we got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
