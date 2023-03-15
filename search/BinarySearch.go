package main

import "fmt"

// NOTE: `Binary Search` is a variation of `Divide and Conquer` algorithm.

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

func checkEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		pos := binarySearchV2(a[i], b)
		if pos == -1 {
			return false
		}
	}
	return true
}

func main() {
	arr0 := []int{0, 2, 4, 5, 7, 9, 11, 12, 13, 18}
	arr1 := []int{0, 2, 4, 7, 5, 9, 18, 11, 12, 13}
	fmt.Println(binarySearchV1(7, arr0), arr0[4])                        // Idx: 4, Val: 7.
	fmt.Println(binarySearchV2(7, arr0), arr0[4])                        // Idx: 4, Val: 7.
	fmt.Println(binarySearchRecursive(7, 0, len(arr0)-1, arr0), arr0[4]) // Idx: 4, Val: 7.
	fmt.Println(checkEquals(arr0, arr1))

	str := "aabaabbbcvbv"
	fmt.Println(mapUniqueChars(str))  // `map[97:[0 1 3 4] 98:[2 5 6 7 10] 99:[8] 118:[9 11]]`.
	fmt.Println(detectPosUnique(str)) // 8.

	nakedGraph := NewGraph()
	nakedGraph.AddVertexToRoot(
		Node{
			Label:     "dude0",
			Adjacents: map[string][]Node{"dude0": {Node{"dude0.0", nil}, Node{"dude0.1", nil}}},
		})
	nakedGraph.AddVertexToRoot(Node{"dude1", make(map[string][]Node)})
	fmt.Printf("%+v\n", nakedGraph)
}
