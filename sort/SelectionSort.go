package main

import (
	"fmt"
	"sort"
)

func findSmallestIndex(arr []int) int {
	smallest := arr[0]
	smallestIdx := 0
	for i := range arr {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIdx = i
		}
	}
	return smallestIdx
}

func selectionSort(arr []int) []int {
	sortedArr := make([]int, 0, len(arr))
	for range arr {
		smallestIdx := findSmallestIndex(arr)
		sortedArr = append(sortedArr, arr[smallestIdx])
		arr = append(arr[:smallestIdx], arr[smallestIdx+1:]...)
	}
	return sortedArr
}

func main() {
	arr0 := []int{1, 3, 2, 5, 7, 6, 9, 12, 15, 13, 22, 20, 36}

	arr1 := make([]int, 0, len(arr0))
	arr1 = append(arr1, arr0...)
	sort.Ints(arr1)

	sortedArr := selectionSort(arr0)

	fmt.Println(sortedArr)
	fmt.Println(arr1)
}
