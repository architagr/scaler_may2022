package maximum_xor_subarray

import "math"

type Node struct {
	index   int
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
			node, ok := temp.hashmap[bit]
			if !ok {
				node = new(Node)
				node.hashmap = make(map[int]*Node)
				temp.hashmap[bit] = node
			}
			temp = node

		}
		temp.index = i
	}
	return root
}
func MaxXorSubarray(A []int) []int {
	prefixXor := getPrefixXor(A)
	prefixXor = append([]int{0}, prefixXor...)
	root := createTries(prefixXor)
	max := 0
	for i := 0; i < len(prefixXor); i++ {
		number := prefixXor[i]
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
	hashMap := make(map[int]int)
	minLen := math.MaxInt
	minStart := math.MaxInt
	minEnd := math.MaxInt
	for i := 0; i < len(prefixXor); i++ {
		val := max ^ prefixXor[i]
		if start, ok := hashMap[val]; ok {
			start++
			end := i
			len := end - start + 1
			if len < minLen {
				minLen = len
				minStart = start
				minEnd = end
			} else if len == minLen && start < minStart {
				minLen = len
				minStart = start
				minEnd = end
			}
		}
		hashMap[prefixXor[i]] = i
	}
	return []int{minStart, minEnd}

}
func maxvalue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getPrefixXor(A []int) []int {
	result := make([]int, len(A))
	result[0] = A[0]
	for i := 1; i < len(A); i++ {
		result[i] = result[i-1] ^ A[i]
	}
	return result
}
