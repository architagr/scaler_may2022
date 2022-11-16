package special_median

import "fmt"

type MinHeap struct {
	data []float64
	size int
	cap  int
}

func InitMinHeap(cap int) *MinHeap {
	return &MinHeap{
		data: make([]float64, 0, cap),
		size: 0,
		cap:  cap,
	}
}

func (minheap *MinHeap) Add(value float64) {
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
		minheap.data = append(make([]float64, 0, cap(minheap.data)+minheap.cap), minheap.data...)
	}
}
func (minHeap *MinHeap) RemoveMin() (float64, error) {
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
func (minHeap *MinHeap) GetMin() (float64, error) {
	if minHeap.size == 0 {
		return 0, fmt.Errorf("no data in heap")
	}
	return minHeap.data[0], nil
}

type MaxHeap struct {
	data []float64
	size int
	cap  int
}

func InitMaxHeap(cap int) *MaxHeap {
	return &MaxHeap{
		data: make([]float64, 0, cap),
		size: 0,
		cap:  cap,
	}
}

func (maxheap *MaxHeap) Add(value float64) {
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
		maxheap.data = append(make([]float64, 0, cap(maxheap.data)+maxheap.cap), maxheap.data...)
	}
}
func (maxHeap *MaxHeap) RemoveMax() (float64, error) {
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
func (maxHeap *MaxHeap) GetMax() (float64, error) {
	if maxHeap.size == 0 {
		return 0, fmt.Errorf("no data in heap")
	}
	return maxHeap.data[0], nil
}
func SpecialMedian(A []int) int {

	minHeap := InitMinHeap(len(A))
	maxHeap := InitMaxHeap(len(A))
	maxHeap.Add(float64(A[0]))
	median, _ := maxHeap.GetMax()

	for i := 1; i < len(A); i++ {
		if median == float64(A[i]) {
			return 1
		}
		max, _ := maxHeap.GetMax()
		if float64(A[i]) > max {
			minHeap.Add(float64(A[i]))
		} else {
			maxHeap.Add(float64(A[i]))
		}

		maxSize := maxHeap.size
		minSize := minHeap.size

		if maxSize-minSize > 1 {
			max, _ = maxHeap.RemoveMax()
			minHeap.Add(max)
		} else if minSize-maxSize > 0 {
			min, _ := minHeap.RemoveMin()
			maxHeap.Add(min)
		}
		if i&1 == 1 {
			min, _ := minHeap.GetMin()
			max, _ := maxHeap.GetMax()
			median = (min + max) / 2
		} else {
			median, _ = maxHeap.GetMax()
		}
	}

	n := len(A)
	minHeap = InitMinHeap(len(A))
	maxHeap = InitMaxHeap(len(A))
	maxHeap.Add(float64(A[n-1]))
	median, _ = maxHeap.GetMax()

	for i := n - 2; i >= 0; i-- {
		if median == float64(A[i]) {
			return 1
		}
		max, _ := maxHeap.GetMax()
		if float64(A[i]) > max {
			minHeap.Add(float64(A[i]))
		} else {
			maxHeap.Add(float64(A[i]))
		}

		maxSize := maxHeap.size
		minSize := minHeap.size

		if maxSize-minSize > 1 {
			max, _ = maxHeap.RemoveMax()
			minHeap.Add(max)
		} else if minSize-maxSize > 0 {
			min, _ := minHeap.RemoveMin()
			maxHeap.Add(min)
		}
		if (n-i-1)&1 == 1 {
			min, _ := minHeap.GetMin()
			max, _ := maxHeap.GetMax()
			median = (min + max) / 2
		} else {
			median, _ = maxHeap.GetMax()
		}
	}

	return 0
}
