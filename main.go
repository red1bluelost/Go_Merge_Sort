package main

import (
	"fmt"
	"time"

	bim "github.com/red1bluelost/Go_Merge_Sort/basic_iterative_merge"
	brm "github.com/red1bluelost/Go_Merge_Sort/basic_recursive_merge"
)

func main() {
	fmt.Println("Hello World")

	start := time.Now()
	slice1 := []int{8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println("Slice1 before:", slice1)

	brm.MergeSort(slice1)

	fmt.Println("Slice1 after: ", slice1)

	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s\n", elapsed)

	slice2 := []int{8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println("Slice2 before:", slice2)

	bim.MergeSort(slice2)

	fmt.Println("Slice2 after: ", slice2)
}
