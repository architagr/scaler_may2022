package construct_roads

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA   int
	inputB   [][]int
	expected int
}

func TestCheckRoads(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   3,
		inputB:   [][]int{{1, 2}, {1, 3}},
		expected: 0,
	})
	testCases = append(testCases, TestCase{
		inputA:   5,
		inputB:   [][]int{{1, 3}, {1, 4}, {3, 2}, {3, 5}},
		expected: 2,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := CheckRoads(testCase.inputA, testCase.inputB)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
