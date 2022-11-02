package shortest_unique_prefix

type Node struct {
	count   int
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
	}
	return root
}

func Prefix(A []string) []string {
	root := createTries(A)
	result := make([]string, len(A))
	for i := 0; i < len(A); i++ {
		word := A[i]
		temp := root
		for j := 0; j < len(word); j++ {
			if node, ok := temp.hashmap[word[j]]; ok {
				if node.count == 1 {
					result[i] = string(word[:j+1])
					break
				}
				temp = node
			}
		}
	}
	return result
}
