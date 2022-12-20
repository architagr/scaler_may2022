package dijsktra

import (
	"fmt"
	"math"
)

func solve(A int, B [][]int, C int) []int {
	dis := make([]int, A)
	for i := 0; i < A; i++ {
		dis[i] = math.MaxInt
	}
	at := make(map[int]map[int]int)
	for i := 0; i < len(B); i++ {
		val, ok := at[B[i][0]]
		if !ok {
			val = make(map[int]int)
		}
		if _, ok := val[B[i][1]]; !ok {
			val[B[i][1]] = B[i][2]
			at[B[i][0]] = val
		}

		val, ok = at[B[i][1]]
		if !ok {
			val = make(map[int]int)
		}
		if _, ok := val[B[i][0]]; !ok {
			val[B[i][0]] = B[i][2]
			at[B[i][1]] = val
		}
	}
	minHeap := InitMinHeap(A)
	minHeap.Add(Node{dis: 0, node: C})
	dis[C] = 0
	for minHeap.size > 0 {
		n, _ := minHeap.RemoveMin()
		if dis[n.node] == n.dis {
			nodes := at[n.node]
			for no, d := range nodes {
				if n.dis+d < dis[no] {
					dis[no] = n.dis + d
					minHeap.Add(Node{dis: n.dis + d, node: no})
				}
			}
		}
	}
	for i := 0; i < A; i++ {
		if dis[i] == math.MaxInt {
			dis[i] = -1
		}
	}
	return dis
}

type Node struct {
	dis, node int
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
		if minheap.data[parentIndex].dis > minheap.data[i].dis {
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
			if left < minHeap.size && minHeap.data[i].dis > minHeap.data[left].dis {
				pos = left
			}
			if right < minHeap.size && minHeap.data[i].dis > minHeap.data[right].dis && minHeap.data[left].dis > minHeap.data[right].dis {
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
