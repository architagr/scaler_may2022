package magician_and_chocolates

import "fmt"

type MaxHeap struct {
	data []int
	size int
	cap  int
}

func InitMaxHeap(cap int) *MaxHeap {
	return &MaxHeap{
		data: make([]int, 0, cap),
		size: 0,
		cap:  cap,
	}
}

func (maxheap *MaxHeap) Add(value int) {
	if len(maxheap.data) <= maxheap.size {
		maxheap.data = append(maxheap.data, value)
	} else {
		maxheap.data[maxheap.size] = value
	}

	i := maxheap.size
	for i > 0 {
		parentIndex := (i - 1) / 2
		if maxheap.data[parentIndex] < maxheap.data[i] {
			maxheap.data[parentIndex], maxheap.data[i] = maxheap.data[i], maxheap.data[parentIndex]
		}
		i = parentIndex
	}

	maxheap.size++
	if maxheap.size >= cap(maxheap.data)-10 {
		maxheap.data = append(make([]int, 0, cap(maxheap.data)+maxheap.cap), maxheap.data...)
	}
}
func (maxHeap *MaxHeap) RemoveMax() (int, error) {
	if maxHeap.size == 0 {
		return 0, fmt.Errorf("no data in heap")
	}
	min := maxHeap.data[0]
	maxHeap.data[0], maxHeap.data[maxHeap.size-1] = maxHeap.data[maxHeap.size-1], maxHeap.data[0]
	maxHeap.size--
	if maxHeap.size > 1 {
		i := 0
		for i < maxHeap.size {
			left := (2 * i) + 1
			right := (2 * i) + 2
			pos := 0
			if left < maxHeap.size && maxHeap.data[i] < maxHeap.data[left] {
				pos = left
			}
			if right < maxHeap.size && maxHeap.data[i] < maxHeap.data[right] && maxHeap.data[left] < maxHeap.data[right] {
				pos = right
			}

			if pos == 0 {
				break
			} else {
				maxHeap.data[i], maxHeap.data[pos] = maxHeap.data[pos], maxHeap.data[i]
				i = pos
			}
		}
	}

	return min, nil
}
func (maxHeap *MaxHeap) GetMax() (int, error) {
	if maxHeap.size == 0 {
		return 0, fmt.Errorf("no data in heap")
	}
	return maxHeap.data[0], nil
}

func nchoc(A int, B []int) int {
	result := 0
	maxHeap := InitMaxHeap(len(B))
	for i := 0; i < len(B); i++ {
		maxHeap.Add(B[i])
	}
	for i := 0; i < A; i++ {
		v, _ := maxHeap.RemoveMax()
		result = (result + v) % 1000000007
		v /= 2
		maxHeap.Add(v)
	}
	return result
}
