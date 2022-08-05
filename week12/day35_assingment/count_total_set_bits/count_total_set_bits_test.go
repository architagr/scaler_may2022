package count_total_set_bits

import (
	"fmt"
	"testing"
)

type TestCase struct {
	a, expected int
}

func TestCountSetBits(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		a:        4,
		expected: 5,
	})
	testCases = append(testCases, TestCase{
		a:        3,
		expected: 4,
	})
	testCases = append(testCases, TestCase{
		a:        5,
		expected: 7,
	})
	testCases = append(testCases, TestCase{
		a:        8,
		expected: 13,
	})
	testCases = append(testCases, TestCase{
		a:        7,
		expected: 12,
	})
	testCases = append(testCases, TestCase{
		a:        100000000,
		expected: 314447109,
	})
	testCases = append(testCases, TestCase{
		a:        1000000000,
		expected: 846928043,
	})

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", testCase.a), func(tb *testing.T) {
			got := CountSetBits(testCase.a)
			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", testCase.a, testCase.expected, got)
			}
		})
	}
}
