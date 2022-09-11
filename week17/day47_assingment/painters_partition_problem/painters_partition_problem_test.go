package painters_partition_problem

import (
	"fmt"
	"testing"
)

func TestPaint(t *testing.T) {
	type TestCase struct {
		c    []int
		b, a int

		expected int
	}
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		c:        []int{1, 10},
		b:        5,
		a:        2,
		expected: 50,
	})

	testCases = append(testCases, TestCase{
		c:        []int{1, 8, 11, 3},
		b:        1,
		a:        10,
		expected: 11,
	})

	testCases = append(testCases, TestCase{
		c:        []int{185, 186, 938, 558, 655, 461, 441, 234, 902, 681},
		b:        10,
		a:        3,
		expected: 18670,
	})
	testCases = append(testCases, TestCase{
		c:        []int{1000000, 1000000},
		b:        1000000,
		a:        1,
		expected: 9400003,
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := Paint(testCase.a, testCase.b, testCase.c)
			if testCase.expected != got {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
