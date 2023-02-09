package edges_in_mst

import (
	"fmt"
	"sort"
)

type Node struct {
	node1, node2, weight, index int
}

func InitNode(node1, node2, weight, index int) Node {
	return Node{
		node1:  node1,
		node2:  node2,
		weight: weight,
		index:  index,
	}
}

type MinHeap struct {
	data []Node
	size int
	cap  int
}

func InitMinHeap(cap int) *MinHeap {
	return &MinHeap{
		data: make([]Node, 0, cap),
		size: 0,
		cap:  cap,
	}
}

func (minheap *MinHeap) Add(value Node) {
	if len(minheap.data) <= minheap.size {
		minheap.data = append(minheap.data, value)
	} else {
		minheap.data[minheap.size] = value
	}

	i := minheap.size
	for i > 0 {
		parentIndex := (i - 1) / 2
		if minheap.data[parentIndex].weight > minheap.data[i].weight {
			minheap.data[parentIndex], minheap.data[i] = minheap.data[i], minheap.data[parentIndex]
		}
		i = parentIndex
	}

	minheap.size++
	if minheap.size >= cap(minheap.data)-10 {
		minheap.data = append(make([]Node, 0, cap(minheap.data)+minheap.cap), minheap.data...)
	}
}
func (minHeap *MinHeap) RemoveMin() (Node, error) {
	if minHeap.size == 0 {
		return Node{}, fmt.Errorf("no data in heap")
	}
	min := minHeap.data[0]
	minHeap.data[0], minHeap.data[minHeap.size-1] = minHeap.data[minHeap.size-1], minHeap.data[0]
	minHeap.size--
	if minHeap.size > 1 {
		i := 0
		for i < minHeap.size {
			left := (2 * i) + 1
			right := (2 * i) + 2
			pos := 0
			if left < minHeap.size && minHeap.data[i].weight > minHeap.data[left].weight {
				pos = left
			}
			if right < minHeap.size && minHeap.data[i].weight > minHeap.data[right].weight && minHeap.data[left].weight > minHeap.data[right].weight {
				pos = right

			}

			if pos == 0 {
				break
			} else {
				minHeap.data[i], minHeap.data[pos] = minHeap.data[pos], minHeap.data[i]
				i = pos
			}
		}
	}

	return min, nil
}
func (minHeap *MinHeap) GetMin() (Node, error) {
	if minHeap.size == 0 {
		return Node{}, fmt.Errorf("no data in heap")
	}
	return minHeap.data[0], nil
}

// prism
func EdgesInMST2(A int, B [][]int) []int {
	visited := make([]bool, A+1)
	ans := make([]int, len(B))
	heap := InitMinHeap(len(B))
	adjencyList := getAdjencyList(B)
	nodes := adjencyList[1]
	visited[1] = true
	updateHashMap(visited, nodes, heap)
	for heap.size > 0 {
		minNode, _ := heap.RemoveMin()
		node := 0
		if visited[minNode.node1] && visited[minNode.node2] {
			continue
		}
		if !visited[minNode.node1] {
			node = minNode.node1
		} else {
			node = minNode.node2
		}
		visited[node] = true
		ans[minNode.index] = 1
		updateHashMap(visited, adjencyList[node], heap)
	}
	return ans
}

func updateHashMap(visited []bool, nodes []Node, heap *MinHeap) {
	for _, node := range nodes {
		if !visited[node.node1] {
			heap.Add(node)
		}
	}
}
func getAdjencyList(B [][]int) map[int][]Node {
	adjencyList := make(map[int][]Node)
	for i := 0; i < len(B); i++ {
		n1, ok1 := adjencyList[B[i][0]]
		n2, ok2 := adjencyList[B[i][1]]
		if !ok1 {
			n1 = make([]Node, 0)
		}
		if !ok2 {
			n2 = make([]Node, 0)
		}

		n1 = append(n1, InitNode(B[i][1], B[i][0], B[i][2], i))
		n2 = append(n2, InitNode(B[i][0], B[i][1], B[i][2], i))

		adjencyList[B[i][0]] = n1
		adjencyList[B[i][1]] = n2

	}
	return adjencyList
}

func EdgesInMST1(A int, B [][]int) []int {
	for i := 0; i < len(B); i++ {
		B[i] = append(B[i], i)
	}
	sort.Slice(B, func(i, j int) bool {
		return sortNode(B[i], B[j])
	})
	componentArray := make([]int, A+1)
	for i := 0; i < len(componentArray); i++ {
		componentArray[i] = i
	}
	ans := make([]int, len(B))
	for i := 0; i < len(B); i++ {
		pa, pb := parentNode(B[i][0], componentArray), parentNode(B[i][1], componentArray)
		if pa != pb {
			ans[B[i][3]] = 1
			componentArray[maxValue(pa, pb)] = minValue(pa, pb)
		}
	}
	return ans
}
func maxValue(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func minValue(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func sortNode(a, b []int) bool {
	if a[2] == b[2] {
		return a[3] < b[3]
	}
	return a[2] < b[2]
}

func parentNode(node int, componentArr []int) int {
	if componentArr[node] == node {
		return node
	}
	return parentNode(componentArr[node], componentArr)
}

type Pair struct {
	u, v, w int
}

func EdgesInMST(A int, B [][]int) []int {
	minCost := 0
	mod := 1000000007
	hashMap := make(map[string]int)
	list := make([]Pair, 0)
	for i := 0; i < len(B); i++ {
		list = append(list, Pair{
			u: B[i][0],
			v: B[i][1],
			w: B[i][2],
		})
		hashMap[fmt.Sprintf("%d_%d_%d", B[i][0], B[i][1], B[i][2])] = 0
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].w < list[j].w
	})
	parent := make([]int, A+1)
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}

	// check for edge if it can be considered in minimum spanning tree
	for i := 0; i < len(list); i++ {
		edge := list[i]

		// check union method just check 2 different components at a given time
		for j := i; j < len(list); j++ {
			parallelEdge := list[j]
			if parallelEdge.w == edge.w {
				if checkUnion(parallelEdge, parent) {
					hashMap[fmt.Sprintf("%d_%d_%d", parallelEdge.u, parallelEdge.v, parallelEdge.w)] = 1
				}
			} else {
				break
			}
		}

		// union method combines 2 different components at a given time
		for j := i; j < len(list); j++ {
			parallelEdge := list[j]
			if parallelEdge.w == edge.w {
				if union(parallelEdge, parent) {
					// if edge is considered, add weight to cost
					minCost = ((minCost % mod) + (parallelEdge.w % mod)) % mod
					hashMap[fmt.Sprintf("%d_%d_%d", parallelEdge.u, parallelEdge.v, parallelEdge.w)] = 1
				}
			} else {
				break
			}
		}
	}
	ans := make([]int, len(B))
	for i := 0; i < len(B); i++ {
		ans[i] = hashMap[fmt.Sprintf("%d_%d_%d", B[i][0], B[i][1], B[i][2])]
	}
	return ans
}
func checkUnion(edge Pair, parent []int) bool {

	// find parent of source and target node of edge
	u := edge.u
	v := edge.v
	su := find(u, parent)
	sv := find(v, parent)

	// if x and v are part of different components, only then combine them
	return su != sv
	// return false when x and v are part of same component
}

func union(edge Pair, parent []int) bool {

	// find parent of source and target node of edge
	u := edge.u
	v := edge.v
	su := find(u, parent)
	sv := find(v, parent)

	// if x and v are part of different components, only then combine them
	if su != sv {
		parent[su] = sv
		return true
	}
	// return false when x and v are part of same component
	return false
}

func find(x int, parent []int) int {

	// base condition when parent of x = x
	if parent[x] == x {
		return x
	}
	// find top most parent of x
	parent[x] = find(parent[x], parent)
	return parent[x]
}
