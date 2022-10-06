package balanced_paranthesis

type Node struct {
	val  byte
	prev *Node
}

/**
 * @input A : String
 *
 * @Output Integer
 */
func BalancedParanthesis(A string) int {
	var head *Node
	hashMap := make(map[byte]byte)
	hashMap['}'] = '{'
	hashMap[')'] = '('
	hashMap[']'] = '['

	for i := 0; i < len(A); i++ {
		if head == nil {
			newNode := new(Node)
			newNode.val = A[i]
			head = newNode
			continue
		}

		if v, ok := hashMap[A[i]]; ok && head.val == v {
			head = head.prev
		} else {

			newNode := new(Node)
			newNode.val = A[i]
			newNode.prev = head
			head = newNode

		}
	}
	if head != nil {
		return 0
	}
	return 1
}
