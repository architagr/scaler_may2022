package sudoku

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input, expected [][]byte
}

func TestNQueens(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input: [][]byte{
			[]byte("53..7...."),
			[]byte("6..195..."),
			[]byte(".98....6."),
			[]byte("8...6...3"),
			[]byte("4..8.3..1"),
			[]byte("7...2...6"),
			[]byte(".6....28."),
			[]byte("...419..5"),
			[]byte("....8..79"),
		},
		expected: [][]byte{
			[]byte("534678912"),
			[]byte("672195348"),
			[]byte("198342567"),
			[]byte("859761423"),
			[]byte("426853791"),
			[]byte("713924856"),
			[]byte("961537284"),
			[]byte("287419635"),
			[]byte("345286179"),
		},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			SolveSudoku(&testCase.input)

			if !reflect.DeepEqual(testCase.input, testCase.expected) {
				test.Fatalf("tested %d expected \n%+v but got \n%+v", (i + 1), testCase.expected, testCase.input)
			}
		})
	}
}
