package dijsktra

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA, inputC int
	inputB         [][]int
	expected       []int
}

func TestPosibilityOfFinishing(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   6,
		inputB:   [][]int{{0, 4, 9}, {3, 4, 6}, {1, 2, 1}, {2, 5, 1}, {2, 4, 5}, {0, 3, 7}, {0, 1, 1}, {4, 5, 7}, {0, 5, 1}},
		inputC:   4,
		expected: []int{7, 6, 5, 6, 0, 6},
	})

	testCases = append(testCases, TestCase{
		inputA:   7,
		inputB:   [][]int{{2, 4, 10}, {3, 4, 1}, {3, 6, 1}, {1, 2, 4}, {4, 5, 6}},
		inputC:   2,
		expected: []int{-1, 4, 0, 11, 10, 16, 12},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := solve(testCase.inputA, testCase.inputB, testCase.inputC)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
