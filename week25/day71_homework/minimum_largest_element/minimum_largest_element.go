package minimum_largest_element

import "fmt"

type Node struct {
	nextValue, originalValue, currentValue int
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
	return a.nextValue > b.nextValue ||
		(a.nextValue == b.nextValue && a.originalValue > b.originalValue)
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
func MinLargestElement(A []int, B int) int {
	max := A[0]
	minHeap := InitMinHeap(len(A))
	for i := 0; i < len(A); i++ {
		minHeap.Add(Node{
			nextValue:     2 * A[i],
			currentValue:  A[i],
			originalValue: A[i],
		})
		if A[i] > max {
			max = A[i]
		}
	}
	for i := 0; i < B; i++ {
		min, _ := minHeap.RemoveMin()
		if min.nextValue > max {
			max = min.nextValue
		}
		min.currentValue = min.nextValue
		min.nextValue = min.nextValue + min.originalValue

		minHeap.Add(min)
	}

	return max
}
