package rain_water_trapped

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	expected int
}

func TestFindRainWaterTrapped(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{0, 1, 0, 2},
		expected: 1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d ", (i+1)), func(tb *testing.T) {
			got := FindRainWaterTrapped(testCase.input)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %d, got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
