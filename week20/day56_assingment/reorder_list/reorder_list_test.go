package reorder_list

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	head1, head2 *listNode
	expected     []int
}

func TestMergeLinkedLists(t *testing.T) {
	testCases := make([]TestCase, 0)

	h1, h2 := GetTestCase1()
	testCases = append(testCases, TestCase{
		head1:    h1,
		head2:    h2,
		expected: []int{1, 4, 2, 5, 3, 6},
	})
	h3, h4 := GetTestCase2()
	testCases = append(testCases, TestCase{
		head1:    h3,
		head2:    h4,
		expected: []int{1, 5, 2, 6, 3, 7, 4},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := listToArr(mergeLinkedLists(testCase.head1, testCase.head2))

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}

func GetTestCase1() (*listNode, *listNode) {
	node11 := listNode_new(1)
	node12 := listNode_new(2)
	node13 := listNode_new(3)
	node11.next = node12
	node12.next = node13

	node21 := listNode_new(4)
	node22 := listNode_new(5)
	node23 := listNode_new(6)
	node21.next = node22
	node22.next = node23

	return node11, node21
}

func GetTestCase2() (*listNode, *listNode) {
	node11 := listNode_new(1)
	node12 := listNode_new(2)
	node13 := listNode_new(3)
	node14 := listNode_new(4)

	node11.next = node12
	node12.next = node13
	node13.next = node14

	node21 := listNode_new(5)
	node22 := listNode_new(6)
	node23 := listNode_new(7)
	node21.next = node22
	node22.next = node23

	return node11, node21
}

func TestReorderList(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		head1:    GetTestCase3(),
		expected: []int{1, 5, 2, 4, 3},
	})
	testCases = append(testCases, TestCase{
		head1:    GetTestCase4(),
		expected: []int{1, 4, 2, 3},
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := listToArr(ReorderList(testCase.head1))

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}

func GetTestCase3() *listNode {
	node11 := listNode_new(1)
	node12 := listNode_new(2)
	node13 := listNode_new(3)
	node14 := listNode_new(4)
	node15 := listNode_new(5)

	node11.next = node12
	node12.next = node13
	node13.next = node14
	node14.next = node15

	return node11
}

func GetTestCase4() *listNode {
	node11 := listNode_new(1)
	node12 := listNode_new(2)
	node13 := listNode_new(3)
	node14 := listNode_new(4)

	node11.next = node12
	node12.next = node13
	node13.next = node14

	return node11
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
