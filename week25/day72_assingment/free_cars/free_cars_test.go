package free_cars

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA, inputB []int
	expected       int
}

func TestMishaAndCandies(t *testing.T) {
	testCases := make([]TestCase, 0)

	// testCases = append(testCases, TestCase{
	// 	inputA:   []int{1, 3, 2, 3, 3},
	// 	inputB:   []int{5, 6, 1, 3, 9},
	// 	expected: 20,
	// })

	// testCases = append(testCases, TestCase{
	// 	inputA:   []int{3, 8, 7, 5},
	// 	inputB:   []int{3, 1, 7, 19},
	// 	expected: 30,
	// })

	testCases = append(testCases, TestCase{
		inputA:   []int{1, 7, 6, 2, 8, 4, 4, 6, 8, 2},
		inputB:   []int{8, 11, 7, 7, 10, 8, 7, 5, 4, 9},
		expected: 65,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := solve(testCase.inputA, testCase.inputB)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
