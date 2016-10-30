package sort

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{9, 3, 2, 5, 8, 4, 7, 10, 6, 1}
	sortedArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	arr = heapSort(arr)

	if ok := reflect.DeepEqual(arr, sortedArr); !ok {
		t.Error("Heapsort is failed.")
	}
}

func TestMergeSort(t *testing.T) {
	arr := []int{10, 1, 4, 3, 8, 2, 7, 9, 5, 6}
	sortedArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	arr = mergeSort(arr)

	if ok := reflect.DeepEqual(arr, sortedArr); !ok {
		t.Error("Mergesort is failed.")
	}
}
