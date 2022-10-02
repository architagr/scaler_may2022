package merge_two_sorted_list

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	head1, head2 *listNode
	expected     []int
}

func TestReorderList(t *testing.T) {
	testCases := make([]TestCase, 0)

	h1, h2 := GetTestCase1()
	testCases = append(testCases, TestCase{
		head1:    h1,
		head2:    h2,
		expected: []int{4, 5, 8, 11, 15, 20},
	})
	h3, h4 := GetTestCase2()
	testCases = append(testCases, TestCase{
		head1:    h3,
		head2:    h4,
		expected: []int{1, 2, 3},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := listToArr(MergeTwoLists(testCase.head1, testCase.head2))

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}

func GetTestCase1() (*listNode, *listNode) {
	node11 := listNode_new(5)
	node12 := listNode_new(8)
	node13 := listNode_new(20)
	node11.next = node12
	node12.next = node13

	node21 := listNode_new(4)
	node22 := listNode_new(11)
	node23 := listNode_new(15)
	node21.next = node22
	node22.next = node23

	return node11, node21
}

func GetTestCase2() (*listNode, *listNode) {
	node11 := listNode_new(1)
	node12 := listNode_new(2)
	node13 := listNode_new(3)

	node11.next = node12
	node12.next = node13

	return node11, nil
}

func listToArr(head *listNode) []int {
	arr := make([]int, 0)
	if head == nil {
		return arr
	}
	for head != nil {
		arr = append(arr, head.value)
		head = head.next
	}
	return arr
}
