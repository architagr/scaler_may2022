package b_closest_points_to_origin

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	a, expected [][]int
	b           int
}

func TestBClosestPointToOrigin(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		a:        [][]int{{48, 18}, {46, 34}, {47, 30}, {19, 9}, {1, 39}, {95, 77}, {95, 106}, {78, 75}, {92, 108}, {89, 83}, {80, 76}},
		expected: [][]int{{19, 9}, {1, 39}, {48, 18}, {47, 30}, {46, 34}},
		b:        5,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := BClosestPointToOrigin(testCase.a, testCase.b)

			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
