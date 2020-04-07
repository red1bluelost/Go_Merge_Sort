package basic_recursive_merge

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	check := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	MergeSort(check)
	if !reflect.DeepEqual(check, expect) {
		t.Errorf("First sort should be %v was %v", expect, check)
	}

	check = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	MergeSort(check)
	if !reflect.DeepEqual(check, expect) {
		t.Errorf("Second sort should be %v was %v", expect, check)
	}

	check = []int{2, 10, 9, 7, 6, 4, 3, 8, 1, 5}
	MergeSort(check)
	if !reflect.DeepEqual(check, expect) {
		t.Errorf("Third sort should be %v was %v", expect, check)
	}

	check = []int{98, 58, 684, 857, 472, 1325, 324, 22, 945, 784, 5, 458, 69853, 41, 475, 842}
	expect = []int{5, 22, 41, 58, 98, 324, 458, 472, 475, 684, 784, 842, 857, 945, 1325, 69853}
	MergeSort(check)
	if !reflect.DeepEqual(check, expect) {
		t.Errorf("Fourth sort should be %v was %v", expect, check)
	}
}

func randomArrayOfLen(n int) []int {
	arr := make([]int, n)
	rand.Seed(0) //int64(time.Now().Nanosecond()))
	for len(arr) < n {
		arr = append(arr, rand.Int())
	}
	return arr
}

func BenchmarkRandomArrayOfLen(b *testing.B) {
	for i := 2; i < b.N; i++ {
		randomArrayOfLen(i)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 2; i < b.N; i++ {
		arr := randomArrayOfLen(i)
		MergeSort(arr)
	}
}