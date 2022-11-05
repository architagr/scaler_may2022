package maximum_xor

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	head     []int
	expected int
}

func TestPrefix(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		head:     []int{1, 2, 3, 4, 5},
		expected: 7,
	})

	testCases = append(testCases, TestCase{
		head:     []int{5, 17, 100, 11},
		expected: 117,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := MaxXor(testCase.head)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
