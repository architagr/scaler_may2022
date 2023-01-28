package valid_path

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	x, y, N, R int
	X, Y       []int
	expected   string
}

func TestPermute(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		x:        2,
		y:        3,
		N:        1,
		R:        1,
		X:        []int{2},
		Y:        []int{3},
		expected: "NO",
	})
	testCases = append(testCases, TestCase{
		x:        41,
		y:        67,
		N:        5,
		R:        0,
		X:        []int{17, 16, 12, 0, 40},
		Y:        []int{52, 61, 61, 25, 31},
		expected: "YES",
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := solve(testCase.x, testCase.y, testCase.N, testCase.R, testCase.X, testCase.Y)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
