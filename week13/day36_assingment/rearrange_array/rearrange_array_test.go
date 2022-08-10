package rearrange_array

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	a, expected []int
}

func TestRearrangeArray(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		a:        []int{4, 0, 2, 1, 3},
		expected: []int{3, 4, 2, 0, 1},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d ", (i+1)), func(tb *testing.T) {
			RearrangeArray(testCase.a)

			if !reflect.DeepEqual(testCase.a, testCase.expected) {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, testCase.a)
			}
		})
	}
}
