package armstrongnumbers

import (
	"fmt"
	"testing"
)

func TestCheckArmstrong(t *testing.T) {
	testCases := make(map[int]bool)
	testCases[5] = false
	testCases[15] = false
	testCases[153] = true
	testCases[1] = true
	for key, expected := range testCases {
		t.Run(fmt.Sprintf("testing for %d", key), func(tb *testing.T) {
			got := checkArmstrong(key)
			if got != expected {
				tb.Errorf("tested for %d expected %t and got %t", key, expected, got)
			}
		})
	}
}

func TestGetArmstrongInRange(t *testing.T) {
	testCases := make(map[int][]int)
	testCases[5] = []int{1}
	testCases[200] = []int{1, 153}

	for key, expected := range testCases {
		t.Run(fmt.Sprintf("testing for %d", key), func(tb *testing.T) {
			got := GetArmstrongInRange(key)
			if !IntArrayEquals(got, expected) {
				tb.Errorf("tested for %d expected %+v and got %+v", key, expected, got)
			}
		})
	}
}

func IntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
