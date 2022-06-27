package largest_sequence_zero_sum

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input, expected []int
}

func TestFindLsZero(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{1, 2, -2, 4, -4},
		expected: []int{2, -2, 4, -4},
	})
	testCases = append(testCases, TestCase{
		input:    []int{1, 2, 2, 4, 4},
		expected: []int{},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := FindLsZero(testCase.input)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}

}
