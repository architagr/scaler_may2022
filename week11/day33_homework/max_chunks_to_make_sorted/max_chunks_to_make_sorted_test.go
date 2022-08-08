package max_chunks_to_make_sorted

import (
	"fmt"
	"testing"
)

type TestCase struct {
	a        []int
	expected int
}

func TestMaxChunk(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		a:        []int{1, 2, 3, 4, 0},
		expected: 1,
	})

	testCases = append(testCases, TestCase{
		a:        []int{2, 0, 1, 3},
		expected: 2,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d ", (i+1)), func(tb *testing.T) {
			got := MaxChunk(testCase.a)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
