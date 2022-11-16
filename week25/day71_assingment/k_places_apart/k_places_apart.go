package k_places_apart

import "fmt"

type MinHeap struct {
	data []int
	size int
	cap  int
}

func InitMinHeap(cap int) *MinHeap {
	return &MinHeap{
		data: make([]int, 0, cap),
		size: 0,
		cap:  cap,
	}
}

func (minheap *MinHeap) Add(value int) {
	if len(minheap.data) <= minheap.size {
		minheap.data = append(minheap.data, value)
	} else {
		minheap.data[minheap.size] = value
	}

	i := minheap.size
	for i > 0 {
		parentIndex := (i - 1) / 2
		if minheap.data[parentIndex] > minheap.data[i] {
			minheap.data[parentIndex], minheap.data[i] = minheap.data[i], minheap.data[parentIndex]
		}
		i = parentIndex
	}

	minheap.size++
	if minheap.size >= cap(minheap.data)-10 {
		minheap.data = append(make([]int, 0, cap(minheap.data)+minheap.cap), minheap.data...)
	}
}
func (minHeap *MinHeap) RemoveMin() (int, error) {
	if minHeap.size == 0 {
		return 0, fmt.Errorf("no data in heap")
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
			if left < minHeap.size && minHeap.data[i] > minHeap.data[left] {
				pos = left
			}
			if right < minHeap.size && minHeap.data[i] > minHeap.data[right] && minHeap.data[left] > minHeap.data[right] {
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
func (minHeap *MinHeap) GetMin() (int, error) {
	if minHeap.size == 0 {
		return 0, fmt.Errorf("no data in heap")
	}
	return minHeap.data[0], nil
}
func KPlacesApart(A []int, B int) []int {
	n := len(A)
	result := make([]int, 0, n)
	minHeap := InitMinHeap(B)
	for i := 0; i <= B && i < n; i++ {
		minHeap.Add(A[i])
	}
	for i := B + 1; i < n; i++ {
		v, _ := minHeap.RemoveMin()
		result = append(result, v)
		minHeap.Add(A[i])
	}
	for minHeap.size > 0 {
		v, _ := minHeap.RemoveMin()
		result = append(result, v)
	}
	return result
}
