package basic_iterative_merge

import u "github.com/red1bluelost/Go_Merge_Sort/utils"

func MergeSort(arr []int) {
	mergeSortIter(arr, len(arr))
}

func mergeSortIter(arr []int, n int) {
	for blockSize := 1; blockSize <= n-1; blockSize *= 2 {
		for offset := 0; offset < n-1; offset += 2 * blockSize {
			middle := u.Min(offset+blockSize-1, n-1)
			upper := u.Min(offset+2*blockSize-1, n-1)

			merge(arr, offset, middle, upper)
		}
	}
}

func merge(arr []int, l, m, r int) {
	lenL, lenR := m-l+1, r-m

	//create temporary slices
	arrL, arrR := make([]int, lenL), make([]int, lenR)
	copy(arrL, arr[l:m+1])
	copy(arrR, arr[m+1:r+1])

	//merge the temporary arrays into the main slice
	i, j, k := 0, 0, l
	for i < lenL && j < lenR {
		if arrL[i] <= arrR[j] {
			arr[k] = arrL[i]
			i++
		} else {
			arr[k] = arrR[j]
			j++
		}
		k++
	}

	//copy remaining elements into main slice if necessary
	for i < lenL {
		arr[k] = arrL[i]
		i++
		k++
	}
	for j < lenR {
		arr[k] = arrR[j]
		j++
		k++
	}
}
