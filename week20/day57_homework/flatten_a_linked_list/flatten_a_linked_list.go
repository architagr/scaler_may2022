package flatten_a_linked_list

type listNode struct {
	right *listNode
	down  *listNode
	val   int
}

func Flatten(root *listNode) *listNode {

	if root == nil || root.right == nil {
		return root
	}
	temp := root
	ans := merge(temp, temp.right)
	temp = temp.right
	if temp.right != nil {
		for temp != nil {
			temp = temp.right
			ans = merge(ans, temp)
		}
	}
	return ans
}

func merge(a, b *listNode) *listNode {
	var head *listNode
	if a == nil && b == nil {
		return head
	} else if a == nil {
		return b
	} else if b == nil {
		return a
	}
	if a.val < b.val {
		head = a
		a = a.down
	} else {
		head = b
		b = b.down
	}
	temp := head
	for a != nil && b != nil {
		if a.val < b.val {
			temp.down = a
			a = a.down
		} else {
			temp.down = b
			b = b.down
		}
		temp = temp.down
	}
	if a != nil {
		temp.down = a
	}

	if b != nil {
		temp.down = b
	}

	return head
}
