package search_in_bitonic_array

import (
	"fmt"
	"testing"
)

func TestBitonicArraySearch(t *testing.T) {
	type TestCase struct {
		input    []int
		b        int
		expected int
	}
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    []int{5, 6, 7, 8, 9, 10, 3, 2, 1},
		b:        30,
		expected: -1,
	})

	testCases = append(testCases, TestCase{
		input:    []int{3, 9, 10, 20, 17, 5, 1},
		b:        20,
		expected: 3,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := BitonicArraySearch(testCase.input, testCase.b)
			if testCase.expected != got {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}

func TestFindPivotIndex(t *testing.T) {
	type TestCase struct {
		input    []int
		expected int
	}
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    []int{5, 6, 7, 8, 9, 10, 3, 2, 1},
		expected: 5,
	})

	testCases = append(testCases, TestCase{
		input:    []int{3, 9, 10, 20, 17, 5, 1},
		expected: 3,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := findPivotIndex(testCase.input)
			if testCase.expected != got {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
