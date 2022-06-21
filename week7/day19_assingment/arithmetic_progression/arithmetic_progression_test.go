package arithmetic_progression

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	expected int
}

func TestIfAp(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{3, 5, 1},
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		input:    []int{2, 4, 1},
		expected: 0,
	})

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", testCase.input), func(tb *testing.T) {
			got := IfAp(testCase.input)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %d but got %d", testCase.input, testCase.expected, got)
			}
		})
	}

}
