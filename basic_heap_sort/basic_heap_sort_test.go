package basic_heap_sort

import (
	u "github.com/red1bluelost/Go_Merge_Sort/utils"
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	check := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	HeapSort(check)
	if !reflect.DeepEqual(check, expect) {
		t.Errorf("First sort should be %v was %v", expect, check)
	}

	check = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	HeapSort(check)
	if !reflect.DeepEqual(check, expect) {
		t.Errorf("Second sort should be %v was %v", expect, check)
	}

	check = []int{2, 10, 9, 7, 6, 4, 3, 8, 1, 5}
	HeapSort(check)
	if !reflect.DeepEqual(check, expect) {
		t.Errorf("Third sort should be %v was %v", expect, check)
	}

	check = []int{98, 58, 684, 857, 472, 1325, 324, 22, 945, 784, 5, 458, 69853, 41, 475, 842}
	expect = []int{5, 22, 41, 58, 98, 324, 458, 472, 475, 684, 784, 842, 857, 945, 1325, 69853}
	HeapSort(check)
	if !reflect.DeepEqual(check, expect) {
		t.Errorf("Fourth sort should be %v was %v", expect, check)
	}
}

func TestBuildHeap(t *testing.T) {
	check := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect := []int{10, 9, 6, 7, 8, 2, 5, 1, 4, 3}
	buildHeap(check, len(check))
	if !reflect.DeepEqual(check, expect) {
		t.Errorf("First build heap should be %v was %v", expect, check)
	}

	check = []int{1, 2, 3}
	expect = []int{3, 1, 2}
	buildHeap(check, len(check))
	if !reflect.DeepEqual(check, expect) {
		t.Errorf("Second build heap should be %v was %v", expect, check)
	}

	check = []int{1, 2, 3}
	expect = []int{3, 1, 2}
	buildHeap(check, len(check))
	if !reflect.DeepEqual(check, expect) {
		t.Errorf("Third build heap should be %v was %v", expect, check)
	}
}

func TestDownHeap(t *testing.T) {
	check := []int{1, 10, 9, 8, 7, 6, 5, 4, 3, 2}
	expect := []int{10, 8, 9, 4, 7, 6, 5, 1, 3, 2}
	downHeap(check, len(check), 0)
	if !reflect.DeepEqual(check, expect) {
		t.Errorf("First down heap should be %v was %v", expect, check)
	}

	check = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	expect = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	downHeap(check, len(check), 0)
	if !reflect.DeepEqual(check, expect) {
		t.Errorf("Second down heap should be %v was %v", expect, check)
	}

	check = []int{3, 4, 5, 6, 7, 8, 9, 1, 2, 10}
	expect = []int{5, 4, 9, 6, 7, 8, 3, 1, 2, 10}
	downHeap(check, len(check), 0)
	if !reflect.DeepEqual(check, expect) {
		t.Errorf("Third down heap should be %v was %v", expect, check)
	}
}

func BenchmarkRandomArrayOfLen(b *testing.B) {
	for i := 2; i < b.N; i++ {
		u.RandomArrayOfLen(i)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	for i := 2; i < b.N; i++ {
		arr := u.RandomArrayOfLen(i)
		HeapSort(arr)
	}
}
