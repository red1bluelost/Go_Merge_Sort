package main

import (
	"fmt"
	"time"

	bim "github.com/red1bluelost/Go_Merge_Sort/basic_iterative_merge"
	brm "github.com/red1bluelost/Go_Merge_Sort/basic_recursive_merge"
	cim "github.com/red1bluelost/Go_Merge_Sort/concurrent_iterative_merge"
	u "github.com/red1bluelost/Go_Merge_Sort/utils"
)

func main() {
	size := 1 << 24 //don't go over 28 or pushing it
	slc := u.RandomArrayOfLen(size)
	slcCopy := make([]int, size)
	slcCopy2 := make([]int, size)
	copy(slcCopy, slc)
	copy(slcCopy2, slc)

	fmt.Println("Starting concurrent process")
	start := time.Now()
	cim.MergeSort(slc)
	elapsed := time.Since(start)
	fmt.Printf("Concurrent iterative took %s\n", elapsed)

	fmt.Println("Starting basic iterative process")
	start = time.Now()
	bim.MergeSort(slcCopy)
	elapsed = time.Since(start)
	fmt.Printf("Basic iterative took %s\n", elapsed)

	fmt.Println("Starting basic recursive process")
	start = time.Now()
	brm.MergeSort(slcCopy2)
	elapsed = time.Since(start)
	fmt.Printf("Basic recursive took %s\n", elapsed)
}
