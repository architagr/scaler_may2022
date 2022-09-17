package closest_pair_from_sorted_arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindClosestPair(t *testing.T) {
	type TestCase struct {
		inputA, inputB []int
		c              int
		expected       []int
	}
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   []int{1, 2, 3, 4, 5},
		inputB:   []int{2, 4, 6, 8},
		c:        9,
		expected: []int{1, 8},
	})

	testCases = append(testCases, TestCase{
		inputA:   []int{5, 10, 20},
		inputB:   []int{1, 2, 30},
		c:        13,
		expected: []int{10, 2},
	})
	testCases = append(testCases, TestCase{
		inputA:   []int{1},
		inputB:   []int{2, 4},
		c:        4,
		expected: []int{1, 2},
	})

	testCases = append(testCases, TestCase{
		inputA:   []int{1, 3, 5, 7, 9},
		inputB:   []int{2, 4, 6, 8, 10},
		c:        10,
		expected: []int{1, 8},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := FindClosestPair(testCase.inputA, testCase.inputB, testCase.c)
			if !reflect.DeepEqual(testCase.expected, got) {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
