package very_large_power

import (
	"fmt"
	"testing"
)

type TestCase struct {
	a, b, expected int
}

func TestLargePower(t *testing.T) {
	testCases := make([]TestCase, 0)
	// testCases = append(testCases, TestCase{
	// 	a:        1,
	// 	b:        1,
	// 	expected: 1,
	// })
	// testCases = append(testCases, TestCase{
	// 	a:        2,
	// 	b:        2,
	// 	expected: 4,
	// })
	testCases = append(testCases, TestCase{
		a:        2,
		b:        27,
		expected: 666348826,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := LargePower(testCase.a, testCase.b)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
