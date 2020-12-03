package basic_heap_sort

import u "github.com/red1bluelost/Go_Merge_Sort/utils"

func HeapSort(arr []int) {
	heapSortIter(arr, len(arr))
}

func heapSortIter(arr []int, size int) {
	buildHeap(arr, size)

	for i := size - 1; i > 0; i-- {
		u.Swap(arr, 0, i)
		downHeap(arr, i, 0)
	}
}

func buildHeap(arr []int, size int) {
	// build heap in place
	for i := 0; i < size; i++ {
		upHeap(arr, size, i)
	}
}

func upHeap(arr []int, size, idx int) {
	for idx > 0 {
		if p := parent(idx); arr[idx] > arr[p] {
			u.Swap(arr, idx, p)
			idx = p
		} else {
			break
		}
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func downHeap(arr []int, size, idx int) {
	for idx < size {
		if lc, rc := leftChild(idx), rightChild(idx); !(lc < size && rc < size) {
			break
		} else if rc < size && (arr[idx] < arr[lc] || arr[idx] < arr[rc]) {
			if arr[rc] < arr[lc] {
				u.Swap(arr, idx, lc)
				idx = lc
			} else {
				u.Swap(arr, idx, rc)
				idx = rc
			}
		} else if arr[idx] < arr[lc] {
			u.Swap(arr, idx, lc)
			idx = lc
		} else {
			break
		}
	}
}

func leftChild(i int) int {
	return i*2 + 1
}

func rightChild(i int) int {
	return i*2 + 2
}
