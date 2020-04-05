package basic_recursive_merge

func MergeSort(arr []int) {
	mergeSortRec(arr, 0, len(arr)-1)
}

func mergeSortRec(arr []int, l, r int) {
	if l < r {
		m := l + (r-l)/2
		mergeSortRec(arr, l, m)
		mergeSortRec(arr, m+1, r)
		merge(arr, l, m, r)
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
