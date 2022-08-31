package rotated_sorted_array_search

import (
	"fmt"
	"testing"
)

type TestCasePivotIndex struct {
	a        []int
	expected int
}

func TestFindPivotIndex(t *testing.T) {
	testCases := make([]TestCasePivotIndex, 0)

	testCases = append(testCases, TestCasePivotIndex{
		a:        []int{4, 5, 6, 7, 0, 1, 2, 3},
		expected: 3,
	})

	testCases = append(testCases, TestCasePivotIndex{
		a:        []int{3, 4, 5, 6, 7, 0, 1, 2},
		expected: 4,
	})

	testCases = append(testCases, TestCasePivotIndex{
		a:        []int{1},
		expected: 0,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := findPivotIndex(testCase.a)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}

type TestCase struct {
	a []int

	expected, b int
}

func TestRotatedSortedArray(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		a:        []int{4, 5, 6, 7, 0, 1, 2, 3},
		b:        4,
		expected: 0,
	})

	testCases = append(testCases, TestCase{
		a:        []int{1},
		b:        4,
		expected: -1,
	})

	testCases = append(testCases, TestCase{
		a:        []int{1},
		b:        1,
		expected: 0,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := RotatedSortedArray(testCase.a, testCase.b)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
