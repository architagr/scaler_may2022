package spellingchecker

type Node struct {
	count   int
	isEnd   bool
	hashmap map[byte]*Node
}

func createTries(arr []string) *Node {
	root := new(Node)
	root.hashmap = make(map[byte]*Node)
	for i := 0; i < len(arr); i++ {
		word := arr[i]
		temp := root
		for j := 0; j < len(word); j++ {
			if node, ok := temp.hashmap[word[j]]; ok {
				node.count++
				temp = node
			} else {
				node = new(Node)
				node.count++
				node.hashmap = make(map[byte]*Node)
				temp.hashmap[word[j]] = node
				temp = node
			}
		}
		temp.isEnd = true
	}
	return root
}
func SpellingChecker(A []string, B []string) []int {
	root := createTries(A)
	result := make([]int, len(B))
	for i := 0; i < len(B); i++ {
		word := B[i]
		temp := root
		flag := true
		for j := 0; j < len(word); j++ {
			if node, ok := temp.hashmap[word[j]]; ok {
				temp = node
			} else {
				result[i] = 0
				flag = false
				break
			}
		}
		if flag && temp.isEnd {
			result[i] = 1
		}

	}
	return result
}
