package main

import "fmt"

// Verbose implementation:
func binarySearchV1(num int, arr []int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]
		if guess == num {
			return mid
		} else if guess > num {
			high = mid - 1
		} else if guess < num {
			low = mid + 1
		}
	}
	return -1
}

func binarySearchV2(num int, arr []int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]
		if guess > num {
			high = mid - 1
		} else if guess < num {
			low = mid + 1
		}
		return mid
	}
	return -1
}

func binarySearchRecursive(num, low, high int, arr []int) int {
	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]
		if guess > num {
			binarySearchRecursive(num, low, mid-1, arr)
		}
		if guess < num {
			binarySearchRecursive(num, mid+1, high, arr)
		}
		return mid
	}
	return -1
}

func main() {
	arr := []int{0, 2, 4, 5, 7, 9, 11, 12, 13, 18}
	fmt.Println(binarySearchV1(7, arr), arr[4])                       // Idx: 4, Val: 7.
	fmt.Println(binarySearchV2(7, arr), arr[4])                       // Idx: 4, Val: 7.
	fmt.Println(binarySearchRecursive(7, 0, len(arr)-1, arr), arr[4]) // Idx: 4, Val: 7.
}
