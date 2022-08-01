package max_distance

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    []int
	expected int
}

func TestMaxDistance(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{1, 10},
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		input:    []int{3, 5, 4, 2},
		expected: 2,
	})
	testCases = append(testCases, TestCase{
		input:    []int{1, 1, 1, 1, 1},
		expected: 4,
	})
	testCases = append(testCases, TestCase{
		input:    []int{67163940, 5180365, 17412933, 1372133, 3379743, 2336349, 387636},
		expected: 2,
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := MaxDistance(testCase.input)

			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
