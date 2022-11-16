package minimum_largest_element

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA   []int
	inputB   int
	expected int
}

func TestMishaAndCandies(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   []int{-2, -4, -8, -16},
		inputB:   10,
		expected: -2,
	})
	testCases = append(testCases, TestCase{
		inputA:   []int{8, 6, 4, 2},
		inputB:   8,
		expected: 12,
	})

	testCases = append(testCases, TestCase{
		inputA:   []int{1, 2, 3, 4},
		inputB:   3,
		expected: 4,
	})

	testCases = append(testCases, TestCase{
		inputA:   []int{5, 1, 4, 2},
		inputB:   5,
		expected: 5,
	})

	testCases = append(testCases, TestCase{
		inputA:   []int{5, 7, 8},
		inputB:   9,
		expected: 28,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := MinLargestElement(testCase.inputA, testCase.inputB)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
