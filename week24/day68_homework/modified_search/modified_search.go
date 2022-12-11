package modified_search

type Node struct {
	count   int
	isEnd   bool
	hashmap map[byte]*Node
}

func NewNode() *Node {
	node := new(Node)
	node.hashmap = make(map[byte]*Node)
	node.count = 0
	node.isEnd = false
	return node
}
func createTries(arr []string) *Node {
	root := NewNode()
	for i := 0; i < len(arr); i++ {
		word := arr[i]
		temp := root
		for j := 0; j < len(word); j++ {
			node, ok := temp.hashmap[word[j]]
			if ok {
				node.count++
			}
			if !ok {
				temp.hashmap[word[j]] = NewNode()
			}
			temp = temp.hashmap[word[j]]
		}
		temp.isEnd = true
	}
	return root
}

func search(node *Node, word string, N, idx int, isModified bool) bool {
	if idx == N && isModified {
		return true
	}
	if node == nil || (idx == N && !isModified) {
		return false
	}

	res := false
	for key, value := range node.hashmap {
		if isModified && key == word[idx] {
			res = res || search(value, word, N, idx+1, isModified)
		} else if !isModified {
			if key == word[idx] {
				res = res || search(value, word, N, idx+1, isModified)
			} else {
				res = res || search(value, word, N, idx+1, true)
			}
		}
	}

	return res
}
func ModifiesSearch(A, B []string) string {
	result := make([]byte, len(B))
	root := createTries(A)
	for i := 0; i < len(B); i++ {
		if search(root, B[i], len(B[i]), 0, false) {
			result[i] = '1'
		} else {
			result[i] = '0'
		}
	}
	return string(result)
}
