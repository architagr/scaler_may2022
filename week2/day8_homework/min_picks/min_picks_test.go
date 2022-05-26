package min_picks

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputArr []int
	expected int
}

func TestGetMinPicks(t *testing.T) {
	testCases := make([]TestCase, 0)
	// testCases = append(testCases, TestCase{
	// 	inputArr: []int{0, 2, 9},
	// 	expected: -7,
	// })

	// testCases = append(testCases, TestCase{
	// 	inputArr: []int{5, 17, 100, 1},
	// 	expected: 99,
	// })

	testCases = append(testCases, TestCase{
		inputArr: []int{-98, 54, -52, 15, 23, -97, 12, -64, 52, 85},
		expected: 151,
	})

	for _, testcase := range testCases {
		t.Run(fmt.Sprintf("testing %v", testcase.inputArr), func(tb *testing.T) {
			got := GetMinPicks(testcase.inputArr)
			if got != testcase.expected {
				tb.Errorf("Tested %+v, expected %d but got %d", testcase.inputArr, testcase.expected, got)
			}
		})
	}
}
