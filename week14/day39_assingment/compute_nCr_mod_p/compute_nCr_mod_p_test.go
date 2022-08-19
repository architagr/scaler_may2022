package compute_nCr_mod_p

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	a, b, c  int
	expected int
}

func TestComputeNCR(t *testing.T) {
	testCases := make([]TestCase, 0)

	// testCases = append(testCases, TestCase{
	// 	a:        5,
	// 	b:        2,
	// 	c:        13,
	// 	expected: 10,
	// })

	// testCases = append(testCases, TestCase{
	// 	a:        300,
	// 	b:        116,
	// 	c:        643,
	// 	expected: 323,
	// })

	testCases = append(testCases, TestCase{
		a:        51299,
		b:        25646,
		c:        3685739,
		expected: 2391643,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := ComputeNCR(testCase.a, testCase.b, testCase.c)

			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
