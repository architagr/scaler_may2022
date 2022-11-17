package free_cars

import (
	"fmt"
	"sort"
)

type Node struct {
	time, profit int
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
		if minheap.comp(minheap.data[parentIndex], minheap.data[i]) {
			minheap.data[parentIndex], minheap.data[i] = minheap.data[i], minheap.data[parentIndex]
		}
		i = parentIndex
	}

	minheap.size++
	if minheap.size >= cap(minheap.data)-10 {
		minheap.data = append(make([]Node, 0, cap(minheap.data)+minheap.cap), minheap.data...)
	}
}
func (minheap *MinHeap) comp(a, b Node) bool {
	return a.profit > b.profit
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
			if left < minHeap.size && minHeap.comp(minHeap.data[i], minHeap.data[left]) {
				pos = left
			}
			if right < minHeap.size && minHeap.comp(minHeap.data[i], minHeap.data[right]) && minHeap.comp(minHeap.data[left], minHeap.data[right]) {
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
func solve(A []int, B []int) int {
	MOD := 1000000007
	n := len(A)
	timeMap := make(map[int]Node)
	minHeap := InitMinHeap(n)
	cars := make([]Node, n)
	for i := 0; i < n; i++ {
		cars[i] = Node{
			time:   A[i],
			profit: B[i],
		}
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].time < cars[j].time
	})
	currentTime := 1
	timeMap[currentTime] = cars[0]
	currentTime++
	minHeap.Add(cars[0])
	sum := cars[0].profit
	for i := 1; i < n; i++ {
		time := cars[i].time
		profit := cars[i].profit

		if _, ok := timeMap[time]; !ok {
			timeMap[currentTime] = cars[i]
			currentTime++
			minHeap.Add(cars[i])
			sum = (sum + cars[i].profit) % MOD
			continue
		}

		min, _ := minHeap.GetMin()
		if min.profit < profit {
			min, _ = minHeap.RemoveMin()
			sum = (sum - min.profit) % MOD

			timeMap[time] = cars[i]
			minHeap.Add(cars[i])
			sum = (sum + cars[i].profit) % MOD

		}
	}
	return sum
}
