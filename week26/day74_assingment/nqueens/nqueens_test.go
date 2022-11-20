package nqueens

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    int
	expected [][]string
}

func TestNQueens(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input: 4,
		expected: [][]string{
			{"..Q.", "Q...", "...Q", ".Q.."},
			{".Q..", "...Q", "Q...", "..Q."},
		},
	})

	testCases = append(testCases, TestCase{
		input: 1,
		expected: [][]string{
			{"Q"},
		},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := NQueens(testCase.input)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
