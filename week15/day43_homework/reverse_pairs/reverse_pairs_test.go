package reverse_pairs

import (
	"fmt"
	"testing"
)

type TestCase struct {
	a        []int
	expected int
}

func TestReversePairs(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		a:        []int{1, 3, 2, 3, 1},
		expected: 2,
	})

	testCases = append(testCases, TestCase{
		a:        []int{4, 1, 2},
		expected: 1,
	})

	testCases = append(testCases, TestCase{
		a:        []int{14046, 57239, 78362, 99387, 27609, 55100, 65536, 62099, 40820, 33056, 88380, 78549, 57512, 33137, 81212, 32365, 42276, 65368, 52459, 74924, 25355, 76044, 78056, 45190, 94365, 58869, 20611},
		expected: 51,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := ReversePairs(testCase.a)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
