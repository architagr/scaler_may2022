package contact_finder

/**
 * @input A : Integer array
 * @input B : array of strings
 *
 * @Output Integer array.
 */
func ContactFinder(A []int, B []string) []int {
	root := createTries(A, B)
	result := make([]int, 0, len(A))
	for i := 0; i < len(A); i++ {
		if A[i] == 0 {
			continue
		}
		word := B[i]
		temp := root
		for j := 0; j < len(word); j++ {
			if node, ok := temp.hashmap[word[j]]; ok {
				if j == len(word)-1 {
					result = append(result, node.count)
					break
				}
				temp = node
			} else {
				result = append(result, 0)
				break
			}
		}
	}
	return result
}

type Node struct {
	count   int
	hashmap map[byte]*Node
}

func createTries(A []int, arr []string) *Node {
	root := new(Node)
	root.hashmap = make(map[byte]*Node)
	for i := 0; i < len(arr); i++ {
		if A[i] == 1 {
			continue
		}
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
