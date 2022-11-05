package maximum_xor

type Node struct {
	count   int
	hashmap map[int]*Node
}

func createTries(arr []int) *Node {
	root := new(Node)
	root.hashmap = make(map[int]*Node)
	for i := 0; i < len(arr); i++ {
		number := arr[i]
		temp := root
		for j := 31; j >= 0; j-- {
			bit := number >> j & 1
			if node, ok := temp.hashmap[bit]; ok {
				node.count++
				temp = node
			} else {
				node = new(Node)
				node.count++
				node.hashmap = make(map[int]*Node)
				temp.hashmap[bit] = node
				temp = node
			}
		}
	}
	return root
}
func MaxXor(A []int) int {
	root := createTries(A)
	max := 0
	for i := 0; i < len(A); i++ {
		number := A[i]
		temp := root
		num := 0
		for j := 31; j >= 0; j-- {
			bit := number >> j & 1
			searchBit := bit
			if bit == 0 {
				searchBit = 1
			} else {
				searchBit = 0
			}
			if node, ok := temp.hashmap[searchBit]; ok {
				num += searchBit << j
				temp = node
			} else if node, ok := temp.hashmap[bit]; ok {
				num += bit << j
				temp = node
			}
		}
		max = maxvalue(num^number, max)
	}
	return max
}

func maxvalue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
