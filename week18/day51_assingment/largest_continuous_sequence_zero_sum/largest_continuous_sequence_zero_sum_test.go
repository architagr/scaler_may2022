package largest_continuous_sequence_zero_sum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLszero(t *testing.T) {
	type TestCase struct {
		input    []int
		expected []int
	}
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    []int{-1, 1, 1},
		expected: []int{-1, 1},
	})

	testCases = append(testCases, TestCase{
		input:    []int{2, -1, 1, 1},
		expected: []int{-1, 1},
	})
	testCases = append(testCases, TestCase{
		input:    []int{1, 2, -2, 4, -4},
		expected: []int{2, -2, 4, -4},
	})

	testCases = append(testCases, TestCase{
		input:    []int{1, 2, -2, 1, 4, -4, 2, -2, 4, -4},
		expected: []int{4, -4, 2, -2, 4, -4},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := Lszero(testCase.input)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
