package utils

import "math/rand"

func RandomArrayOfLen(n int) []int {
	arr := make([]int, n)
	rand.Seed(0) //int64(time.Now().Nanosecond()))
	for len(arr) < n {
		arr = append(arr, rand.Int())
	}
	return arr
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Swap(arr []int, idx1, idx2 int) {
	temp := arr[idx1]
	arr[idx1] = arr[idx2]
	arr[idx2] = temp
}