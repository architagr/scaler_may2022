package maximum_unsorted_subarray

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	a, expected []int
}

func TestSubUnsort(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		a:        []int{1, 3, 2, 4, 5},
		expected: []int{1, 2},
	})
	testCases = append(testCases, TestCase{
		a:        []int{1, 2, 3, 4, 5},
		expected: []int{-1},
	})
	testCases = append(testCases, TestCase{
		a:        []int{1, 1, 3, 4, 4, 11, 11, 15, 17, 19, 11, 15, 11, 16, 15, 18, 18, 19, 17, 16},
		expected: []int{7, 19},
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := SubUnsort(testCase.a)

			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
