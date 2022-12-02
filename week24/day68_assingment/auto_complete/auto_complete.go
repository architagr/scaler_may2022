package auto_complete

import "sort"

type Weight struct {
	PrefixWeight int
	word         string
}
type Node struct {
	Data    byte
	Child   map[byte]*Node
	Weights []Weight
	Weight  int
	count   int
}

func CreateDictonary(words []string, weight []int) *Node {
	root := new(Node)
	root.Child = make(map[byte]*Node)
	root.Weights = make([]Weight, 0, 5)
	root.Weight = 0

	for index, word := range words {
		addWord(word, weight[index], root)
	}
	return root
}
func addWord(word string, weight int, root *Node) {
	temp := root
	for i := 0; i < len(word); i++ {
		node, ok := temp.Child[word[i]]
		if !ok {
			node = new(Node)
			node.Data = word[i]
			node.Child = make(map[byte]*Node)
			node.Weights = make([]Weight, 0, 5)
			node.Weight = 0
			temp.Child[word[i]] = node
		}
		temp = node

		weightNode := Weight{PrefixWeight: weight, word: word}
		temp.Weights = append(temp.Weights, weightNode)
		if len(temp.Weights) > 5 {
			sort.Slice(temp.Weights, func(i, j int) bool {
				return temp.Weights[i].PrefixWeight < temp.Weights[j].PrefixWeight
			})
			temp.Weights = temp.Weights[1:]
		}
	}
	temp.Weight = weight
	temp.count++
}

func Search(prefix string, root *Node) []string {
	result := make([]string, 0, 5)
	temp := root
	for i := 0; i < len(prefix); i++ {
		if node, ok := temp.Child[prefix[i]]; ok {
			temp = node
		} else {
			return result
		}
	}

	sort.Slice(temp.Weights, func(i, j int) bool {
		return temp.Weights[i].PrefixWeight > temp.Weights[j].PrefixWeight
	})
	for i := 0; i < len(temp.Weights); i++ {
		result = append(result, temp.Weights[i].word)
	}
	return result
}
