package little_ponny_max_element

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputArr      []int
	searchElement int
	expected      int
}

func TestLittlePonnyMaxElement(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputArr:      []int{1, 2, 3, 4, 5},
		searchElement: 3,
		expected:      2,
	})

	testCases = append(testCases, TestCase{
		inputArr:      []int{1, 2, 3, 4, 5},
		searchElement: 10,
		expected:      -1,
	})

	for _, testcase := range testCases {
		t.Run(fmt.Sprintf("testing %v", testcase.inputArr), func(tb *testing.T) {
			got := LittlePonnyMaxElement(testcase.inputArr, testcase.searchElement)
			if got != testcase.expected {
				tb.Errorf("Tested %+v to contain %d, expected %d but got %d", testcase.inputArr, testcase.searchElement, testcase.expected, got)
			}
		})
	}
}
