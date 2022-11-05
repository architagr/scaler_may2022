package maximum_xor_subarray

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	head     []int
	expected []int
}

func TestMaxXorSubarray(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		head:     []int{1, 4, 3},
		expected: []int{2, 3},
	})
	testCases = append(testCases, TestCase{
		head:     []int{8},
		expected: []int{1, 1},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := MaxXorSubarray(testCase.head)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
