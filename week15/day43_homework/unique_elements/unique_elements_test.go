package unique_elements

import (
	"fmt"
	"testing"
)

type TestCase struct {
	a        []int
	expected int
}

func TestUniqueElements(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		a:        []int{1, 3, 1},
		expected: 1,
	})

	testCases = append(testCases, TestCase{
		a:        []int{2, 4, 5},
		expected: 0,
	})

	testCases = append(testCases, TestCase{
		a:        []int{2, 2, 2},
		expected: 3,
	})

	testCases = append(testCases, TestCase{
		a:        []int{1, 1, 2, 3, 3, 4, 5},
		expected: 9,
	})

	testCases = append(testCases, TestCase{
		a:        []int{1, 1, 2, 3, 3, 5},
		expected: 6,
	})

	testCases = append(testCases, TestCase{
		a:        []int{1, 1, 2, 3, 3, 5, 1},
		expected: 12,
	})

	testCases = append(testCases, TestCase{
		a:        []int{51, 6, 10, 8, 22, 61, 56, 48, 88, 85, 21, 98, 81, 76, 71, 68, 18, 6, 14, 23, 72, 18, 56, 30, 97, 100, 81, 5, 99, 2, 85, 67, 46, 32, 66, 51, 76, 53, 36, 31, 81, 56, 26, 75, 69, 54, 54, 54, 83, 41, 86, 48, 7, 32, 85, 23, 47, 23, 18, 45, 79, 95, 73, 15, 55, 16, 66, 73, 13, 85, 14, 80, 39, 92, 66, 20, 22, 25, 34, 14, 51, 14, 17, 10, 100, 35, 9, 83, 31, 60, 24, 37, 69, 62},
		expected: 239,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := UniqueElements(testCase.a)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
