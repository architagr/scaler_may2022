package k_reverse_linked_list

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	head     *listNode
	b        int
	expected []int
}

func TestKReverseLinkedList(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		head:     GetTestCase2(),
		b:        2,
		expected: []int{2, 1, 4, 3, 6, 5},
	})

	testCases = append(testCases, TestCase{
		head:     GetTestCase2(),
		b:        3,
		expected: []int{3, 2, 1, 6, 5, 4},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := listToArr(KReverseLinkedList(testCase.head, testCase.b))

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}

func GetTestCase1() *listNode {
	node1 := listNode_new(46)
	node2 := listNode_new(76)
	node3 := listNode_new(35)
	node1.next = node2
	node2.next = node3
	return node1
}

func GetTestCase2() *listNode {
	node1 := listNode_new(1)
	node2 := listNode_new(2)
	node3 := listNode_new(3)
	node4 := listNode_new(4)
	node5 := listNode_new(5)
	node6 := listNode_new(6)

	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node6

	return node1
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
