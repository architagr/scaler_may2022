package number_of_squareful_arrays

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    []int
	expected int
}

func TestPermute(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    []int{2, 2, 2},
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		input:    []int{1, 17, 8},
		expected: 2,
	})

	testCases = append(testCases, TestCase{
		input:    []int{2226, 175, 114, 82, 39, 10, 39, 42, 102, 94, 75},
		expected: 2,
	})
	testCases = append(testCases, TestCase{
		input:    []int{4776},
		expected: 0,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := solve(testCase.input)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
