package array_with_consecutive_elements

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	a        []int
	expected int
}

func TestArrayWithConsecutiveElements(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		a:        []int{3, 2, 1, 4, 5},
		expected: 1,
	})

	testCases = append(testCases, TestCase{
		a:        []int{48, 7, 42, 41, 118, 42, 118, 29, 72, 90, 49, 59, 115, 72, 74, 42, 73, 119, 64, 28, 98, 47, 47, 38, 43, 25, 22, 102, 80, 109, 82, 45, 80, 122, 40, 74, 9, 57, 103, 28, 42, 37, 41, 19, 117, 55, 43, 37, 83, 47, 87, 66, 88, 72, 89, 123, 42, 123, 117, 11, 75, 6, 89, 4, 120, 4, 83, 38, 4, 21, 31, 116, 53, 88, 112, 84, 26, 39, 85, 40, 110, 117, 44, 46, 66, 48, 106, 46, 83, 73, 96, 56, 55, 3, 89, 92, 52, 60, 119, 78, 92, 53, 64, 107, 102, 95, 111, 8, 69, 50, 118, 29, 28, 92, 10, 62, 59, 5, 81, 34, 84, 17, 12, 111, 65, 9, 122, 101, 46, 97, 110, 32, 123, 118, 121, 34, 52, 123, 58, 34, 44, 10, 114, 4, 56, 122, 91, 23, 81},
		expected: 0,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := ArrayWithConsecutiveElements(testCase.a)

			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
