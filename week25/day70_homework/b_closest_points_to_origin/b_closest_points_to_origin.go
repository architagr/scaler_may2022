package b_closest_points_to_origin

import "fmt"

type Data struct {
	distance int
	index    int
}
type MinHeap struct {
	data []Data
	size int
	cap  int
}

func InitMinHeap(cap int) *MinHeap {
	return &MinHeap{
		data: make([]Data, 0, cap),
		size: 0,
		cap:  cap,
	}
}

func (minheap *MinHeap) Add(value Data) {
	if len(minheap.data) <= minheap.size {
		minheap.data = append(minheap.data, value)
	} else {
		minheap.data[minheap.size] = value
	}

	i := minheap.size
	for i > 0 {
		parentIndex := (i - 1) / 2
		if minheap.data[parentIndex].distance > minheap.data[i].distance {
			minheap.data[parentIndex], minheap.data[i] = minheap.data[i], minheap.data[parentIndex]
		}
		i = parentIndex
	}

	minheap.size++
	if minheap.size >= cap(minheap.data)-10 {
		minheap.data = append(make([]Data, 0, cap(minheap.data)+minheap.cap), minheap.data...)
	}
}
func (minHeap *MinHeap) RemoveMin() (Data, error) {
	if minHeap.size == 0 {
		return Data{}, fmt.Errorf("no data in heap")
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
			if left < minHeap.size && minHeap.data[i].distance > minHeap.data[left].distance {
				pos = left
			}
			if right < minHeap.size && minHeap.data[i].distance > minHeap.data[right].distance && minHeap.data[left].distance > minHeap.data[right].distance {
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
func (minHeap *MinHeap) GetMin() (Data, error) {
	if minHeap.size == 0 {
		return Data{}, fmt.Errorf("no data in heap")
	}
	return minHeap.data[0], nil
}
func BClosestPoints(A [][]int, B int) [][]int {
	minHeap := InitMinHeap(len(A))
	for i := 0; i < len(A); i++ {
		distance := (A[i][0] * A[i][0]) + (A[i][1] * A[i][1])
		minHeap.Add(Data{
			distance: distance,
			index:    i,
		})
	}
	result := make([][]int, B)

	for i := 0; i < B; i++ {
		v, _ := minHeap.RemoveMin()
		result[i] = A[v.index]
	}
	return result
}
