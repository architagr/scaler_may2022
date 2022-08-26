package inversion_count_in_an_array

import (
	"fmt"
	"testing"
)

type TestCase struct {
	a        []int
	expected int
}

func TestGetInversionCount(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		a:        []int{45, 10, 15, 25, 50},
		expected: 3,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := GetInversionCount(testCase.a)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
