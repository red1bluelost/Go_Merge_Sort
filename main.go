package main

import (
	"fmt"

	brm "github.com/red1bluelost/Go_Merge_Sort/basic_recursive_merge"
)

func main() {
	fmt.Println("Hello World")
	slice1 := []int{8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println("Slice1 before:", slice1)

	brm.MergeSort(slice1)

	fmt.Println("Slice1 after: ", slice1)
}
