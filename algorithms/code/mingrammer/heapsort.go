package sort

type Heap struct {
	arr     []int
	arrSize int
	size    int
}

func maxHeapify(heap Heap, i int) {
	var largestIdx int

	lIdx, rIdx := -1, -1

	if 2*i+1 < heap.size {
		lIdx = 2*i + 1
	}

	if 2*(i+1) < heap.size {
		rIdx = 2 * (i + 1)
	}

	if lIdx > 0 && heap.arr[lIdx] > heap.arr[i] {
		largestIdx = lIdx
	} else {
		largestIdx = i
	}

	if rIdx > 0 && heap.arr[rIdx] > heap.arr[largestIdx] {
		largestIdx = rIdx
	}

	if largestIdx != i {
		heap.arr[i], heap.arr[largestIdx] = heap.arr[largestIdx], heap.arr[i]
		maxHeapify(heap, largestIdx)
	}
}

func buildMaxHeap(heap Heap) {
	for i := heap.arrSize - 1; i >= 0; i-- {
		maxHeapify(heap, i)
	}
}

func heapSort(arr []int) []int {
	arrSize := len(arr)
	heapSize := arrSize

	heap := Heap{arr: arr, arrSize: arrSize, size: heapSize}

	buildMaxHeap(heap)

	for i := heap.arrSize - 1; i > 0; i-- {
		heap.arr[i], heap.arr[0] = heap.arr[0], heap.arr[i]
		heap.size--
		maxHeapify(heap, 0)
	}

	return heap.arr
}
