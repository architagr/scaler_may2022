package auto_complete

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	index    int
	words    []string
	weights  []int
	prefixes []string
	expected [][]string
}

func TestCreateDictonary(t *testing.T) {
	testCases := make([]TestCase, 0)
	i := 1
	testCases = append(testCases, TestCase{
		index:    i,
		words:    []string{"abcd", "aecd", "abaa", "abef", "acdcc", "acbcc"},
		weights:  []int{2, 1, 3, 4, 6, 5},
		expected: nil,
	})
	i++
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := CreateDictonary(testCase.words, testCase.weights)

			if len(x.Child) != 1 {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	testCases := make([]TestCase, 0)
	// testCases = append(testCases, TestCase{
	// 	index:    1,
	// 	words:    []string{"abcd", "abcde", "abaa", "abef", "acdcc", "acbcc"},
	// 	weights:  []int{2, 1, 3, 4, 6, 5},
	// 	prefixes: []string{"a"},
	// 	expected: [][]string{{"acdcc", "acbcc", "abef", "abaa", "abcd"}},
	// })
	// testCases = append(testCases, TestCase{
	// 	index:    2,
	// 	words:    []string{"abcd", "aecd", "abaa", "abef", "acdcc", "acbcc"},
	// 	weights:  []int{2, 1, 3, 4, 6, 5},
	// 	prefixes: []string{"a"},
	// 	expected: [][]string{{"acbcc", "acdcc", "abef", "abaa", "abcd"}},
	// })
	// testCases = append(testCases, TestCase{
	// 	index:    3,
	// 	words:    []string{"abcd", "aecd", "abaa", "abef", "acdcc", "acbcc"},
	// 	weights:  []int{2, 1, 3, 4, 6, 5},
	// 	prefixes: []string{"abc"},
	// 	expected: [][]string{{"abcd"}},
	// })
	testCases = append(testCases, TestCase{
		index:    4,
		words:    []string{"abcd", "aecd", "abaa", "abef", "acdcc", "acbcc"},
		weights:  []int{2, 1, 3, 4, 6, 5},
		prefixes: []string{"ab", "abc", "a"},
		expected: [][]string{{"abef", "abaa", "abcd"}, {"abcd"}, {"acdcc", "acbcc", "abef", "abaa", "abcd"}},
	})
	testCases = append(testCases, TestCase{
		index:    5,
		words:    []string{"aaaa", "aacd", "abaa", "abaa", "aadcc"},
		weights:  []int{3, 4, 1, 2, 6},
		prefixes: []string{"aa", "aba", "abc"},
		expected: [][]string{{"aadcc", "aacd", "aaaa"}, {"abaa", "abaa"}, {}},
	})
	testCases = append(testCases, TestCase{
		index:    6,
		words:    []string{"aeedd", "eaa"},
		weights:  []int{3, 11},
		prefixes: []string{"aee", "aeed", "aee"},
		expected: [][]string{{"aeedd"}, {"aeedd"}, {"aeedd"}},
	})
	testCases = append(testCases, TestCase{
		index:    7,
		words:    []string{"aacdbac", "dbaae", "a", "bbd", "cbbdec", "eaedbcdd", "bbddc", "ccced", "ed", "bddedee", "abbbdee", "dcddceb", "aede", "c"},
		weights:  []int{27, 19, 66, 35, 8, 54, 12, 21, 46, 76, 32, 58, 69, 36},
		prefixes: []string{"a", "eaedbcd", "eaedb", "cbb", "ccced", "cbbd", "a", "cc"},
		expected: [][]string{{"aede", "a", "abbbdee", "aacdbac"}, {"eaedbcdd"}, {"eaedbcdd"}, {"cbbdec"}, {"ccced"}, {"cbbdec"}, {"aede", "a", "abbbdee", "aacdbac"}, {"ccced"}},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := CreateDictonary(testCase.words, testCase.weights)
			for index, prefix := range testCase.prefixes {
				words := Search(prefix, x)
				if !reflect.DeepEqual(testCase.expected[index], words) {
					test.Fatalf("tested %d_%d expected %+v but got %+v", (i + 1), index, testCase.expected[index], words)
				}
			}

		})
	}
}

func sameStringSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	// create a map of string -> int
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		// 0 value for int is 0, so just increment a counter for the string
		diff[_x]++
	}
	for _, _y := range y {
		// If the string _y is not in diff bail out early
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y] -= 1
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	return len(diff) == 0
}
