package bth_smallest_prime_fraction

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA   []int
	inputB   int
	expected []int
}

func TestBthSmallestPrimeFraction(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   []int{1, 2, 3, 5},
		inputB:   3,
		expected: []int{2, 5},
	})

	testCases = append(testCases, TestCase{
		inputA:   []int{1, 7},
		inputB:   1,
		expected: []int{1, 7},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := BthSmallestPrimeFraction(testCase.inputA, testCase.inputB)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
