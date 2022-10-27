package valid_binary_search_tree

import (
	"fmt"
	"testing"
)

type TestCase struct {
	head     *treeNode
	expected int
}

func TestReorderList(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		head:     GetTestCase1(),
		expected: 0,
	})
	testCases = append(testCases, TestCase{
		head:     GetTestCase2(),
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		head:     GetTestCase3(),
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		head:     GetTestCase4(),
		expected: 0,
	})

	testCases = append(testCases, TestCase{
		head:     GetTestCase5(),
		expected: 0,
	})

	testCases = append(testCases, TestCase{
		head:     GetTestCase6(),
		expected: 0,
	})

	testCases = append(testCases, TestCase{
		head:     GetTestCase7(),
		expected: 1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := isValidBST(testCase.head)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
func GetTestCase1() *treeNode {
	node1 := treeNode_new(1)
	node21 := treeNode_new(2)
	node22 := treeNode_new(3)
	node1.left = node21
	node1.right = node22
	return node1

}
func GetTestCase2() *treeNode {
	node1 := treeNode_new(2)
	node21 := treeNode_new(1)
	node22 := treeNode_new(3)
	node1.left = node21
	node1.right = node22
	return node1

}

func GetTestCase3() *treeNode {
	node1 := treeNode_new(10)

	node21 := treeNode_new(5)
	node22 := treeNode_new(15)
	node1.left = node21
	node1.right = node22

	node31 := treeNode_new(3)
	node32 := treeNode_new(7)
	node21.left = node31
	node21.right = node32

	node33 := treeNode_new(13)
	node34 := treeNode_new(18)
	node22.left = node33
	node22.right = node34

	return node1
}

func GetTestCase4() *treeNode {
	node1 := treeNode_new(10)

	node21 := treeNode_new(5)
	node22 := treeNode_new(15)
	node1.left = node21
	node1.right = node22

	node31 := treeNode_new(3)
	node32 := treeNode_new(7)
	node21.left = node31
	node21.right = node32

	node33 := treeNode_new(13)
	node34 := treeNode_new(18)
	node22.left = node33
	node22.right = node34

	node41 := treeNode_new(2)
	node42 := treeNode_new(4)
	node31.left = node41
	node31.right = node42

	node43 := treeNode_new(6)
	node44 := treeNode_new(8)
	node32.left = node43
	node32.right = node44

	node45 := treeNode_new(12)
	node46 := treeNode_new(1)
	node33.left = node45
	node33.right = node46

	return node1
}

func GetTestCase5() *treeNode {
	node1 := treeNode_new(4)

	node21 := treeNode_new(2)
	node22 := treeNode_new(5)
	node1.left = node21
	node1.right = node22

	node31 := treeNode_new(1)
	node32 := treeNode_new(5)
	node21.left = node31
	node21.right = node32

	return node1
}

func GetTestCase6() *treeNode {
	node1 := treeNode_new(3)

	node21 := treeNode_new(2)
	node22 := treeNode_new(4)
	node1.left = node21
	node1.right = node22

	node31 := treeNode_new(1)
	node32 := treeNode_new(3)
	node21.left = node31
	node21.right = node32

	return node1
}

func GetTestCase7() *treeNode {
	node1 := treeNode_new(40)

	node21 := treeNode_new(35)
	node22 := treeNode_new(41)
	node1.left = node21
	node1.right = node22

	node31 := treeNode_new(34)
	node32 := treeNode_new(37)
	node21.left = node31
	node21.right = node32

	node33 := treeNode_new(46)
	node22.right = node33

	node41 := treeNode_new(31)
	node31.left = node41

	node42 := treeNode_new(36)
	node43 := treeNode_new(38)
	node32.left = node42
	node32.right = node43

	node44 := treeNode_new(45)
	node45 := treeNode_new(47)
	node33.left = node44
	node33.right = node45

	node51 := treeNode_new(29)
	node52 := treeNode_new(32)
	node41.left = node51
	node41.right = node52

	node53 := treeNode_new(39)
	node43.right = node53

	node54 := treeNode_new(43)
	node44.left = node54

	node61 := treeNode_new(28)
	node62 := treeNode_new(30)
	node51.left = node61
	node51.right = node62

	node63 := treeNode_new(33)
	node52.right = node63

	node64 := treeNode_new(42)
	node65 := treeNode_new(44)
	node54.left = node64
	node54.right = node65

	node71 := treeNode_new(25)
	node61.left = node71

	node81 := treeNode_new(24)
	node82 := treeNode_new(26)
	node71.left = node81
	node71.right = node82

	node91 := treeNode_new(27)
	node82.right = node91

	return node1
}
