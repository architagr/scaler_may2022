package special_median

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA   []int
	expected int
}

func TestMishaAndCandies(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   []int{4, 6, 8, 4},
		expected: 1,
	})

	testCases = append(testCases, TestCase{
		inputA:   []int{2147483647, -2147483648, 0},
		expected: 0,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := SpecialMedian(testCase.inputA)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
