package misha_and_candies

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA   []int
	inputB   int
	expected int
}

func TestMishaAndCandies(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   []int{110, 289, 585, 135, 314, 259, 238, 798, 954, 399, 492, 282, 311, 177, 804, 769, 191, 539, 830, 806, 854, 50, 373, 329, 549, 202, 399, 542, 166, 986, 446, 346, 58, 269, 103, 267, 547, 171, 713, 408, 847, 988, 704, 162, 850, 493, 59, 584, 588, 117, 639, 310, 615, 871, 977, 136, 973, 324, 450, 737, 778, 139, 606, 98, 968, 275, 911, 558, 849, 62, 495, 512, 106, 682, 980, 9, 517, 104, 80, 540, 689, 696, 396, 681, 541, 468, 12, 482, 459, 438, 924, 507, 725, 388, 579, 754, 421, 30, 151},
		inputB:   80,
		expected: 131,
	})

	testCases = append(testCases, TestCase{
		inputA:   []int{1, 2, 1},
		inputB:   2,
		expected: 1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := MishaAndCandies(testCase.inputA, testCase.inputB)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
