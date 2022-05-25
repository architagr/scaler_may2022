package searchelement

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputArr      []int
	searchElement int
	expected      bool
}

func TestSearchElement(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputArr:      []int{1, 2, 3, 4, 5},
		searchElement: 1,
		expected:      true,
	})

	testCases = append(testCases, TestCase{
		inputArr:      []int{1, 2, 3, 4, 5},
		searchElement: 10,
		expected:      false,
	})

	for _, testcase := range testCases {
		t.Run(fmt.Sprintf("testing %v", testcase.inputArr), func(tb *testing.T) {
			got := SearchElement(testcase.inputArr, testcase.searchElement)
			if got != testcase.expected {
				tb.Errorf("Tested %+v to contain %d, expected %t but got %t", testcase.inputArr, testcase.searchElement, testcase.expected, got)
			}
		})
	}
}
