package bth_smallest_prime_fraction

import "fmt"

func BthSmallestPrimeFraction(A []int, B int) []int {
	maxHeap := InitMaxHeap(B)
	for i := len(A) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			fr := float64(A[j]) / float64(A[i])
			if maxHeap.size == B {
				max, _ := maxHeap.GetMax()
				if max.fr > fr {
					maxHeap.Add(Node{fr: fr, p: A[j], q: A[i]})
					maxHeap.RemoveMax()
				}
			} else {
				maxHeap.Add(Node{fr: fr, p: A[j], q: A[i]})
			}

		}
	}
	max, _ := maxHeap.GetMax()
	return []int{max.p, max.q}
}

type Node struct {
	fr   float64
	p, q int
}
type MaxHeap struct {
	data []Node
	size int
	cap  int
}

func InitMaxHeap(cap int) *MaxHeap {
	return &MaxHeap{
		data: make([]Node, 0, cap),
		size: 0,
		cap:  cap,
	}
}

func (maxheap *MaxHeap) Add(value Node) {
	if len(maxheap.data) <= maxheap.size {
		maxheap.data = append(maxheap.data, value)
	} else {
		maxheap.data[maxheap.size] = value
	}

	i := maxheap.size
	for i > 0 {
		parentIndex := (i - 1) / 2
		if maxheap.data[parentIndex].fr < maxheap.data[i].fr {
			maxheap.data[parentIndex], maxheap.data[i] = maxheap.data[i], maxheap.data[parentIndex]
		}
		i = parentIndex
	}

	maxheap.size++
	if maxheap.size >= cap(maxheap.data)-10 {
		maxheap.data = append(make([]Node, 0, cap(maxheap.data)+maxheap.cap), maxheap.data...)
	}
}
func (maxHeap *MaxHeap) RemoveMax() (Node, error) {
	if maxHeap.size == 0 {
		return Node{}, fmt.Errorf("no data in heap")
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
			if left < maxHeap.size && maxHeap.data[i].fr < maxHeap.data[left].fr {
				pos = left
			}
			if right < maxHeap.size && maxHeap.data[i].fr < maxHeap.data[right].fr && maxHeap.data[left].fr < maxHeap.data[right].fr {
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
func (maxHeap *MaxHeap) GetMax() (Node, error) {
	if maxHeap.size == 0 {
		return Node{}, fmt.Errorf("no data in heap")
	}
	return maxHeap.data[0], nil
}
