package allocate_books

import (
	"fmt"
	"testing"
)

func TestBooks(t *testing.T) {
	type TestCase struct {
		input    []int
		b        int
		expected int
	}
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    []int{12, 34, 67, 90},
		b:        2,
		expected: 113,
	})

	testCases = append(testCases, TestCase{
		input:    []int{97, 26, 12, 67, 10, 33, 79, 49, 79, 21, 67, 72, 93, 36, 85, 45, 28, 91, 94, 57, 1, 53, 8, 44, 68, 90, 24},
		b:        26,
		expected: 97,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := Books(testCase.input, testCase.b)
			if testCase.expected != got {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
