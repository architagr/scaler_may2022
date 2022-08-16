package prime_sum

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	a        int
	expected []int
}

func TestCountPrimeSubsequence(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		a:        4,
		expected: []int{2, 2},
	})
	testCases = append(testCases, TestCase{
		a:        8,
		expected: []int{3, 5},
	})

	testCases = append(testCases, TestCase{
		a:        10,
		expected: []int{3, 7},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d ", (i+1)), func(tb *testing.T) {
			got := PrimeSum(testCase.a)

			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
