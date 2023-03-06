package main

func sumDAC[number int | uint | float64](arr []number) number {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + sumDAC(arr[1:])
}

func countDAC[number int | uint | float64](arr []number) number {
	if len(arr) == 0 {
		return 0
	}
	return 1 + countDAC(arr[1:])
}

func countElement[number int | uint | float64](arr []number) number {
	var counter number = 0
	for i := range arr {
		if &arr[i] != nil {
			counter++
		}
	}
	return counter
}

func findMax(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 {
		return arr[0]
	}
	max := arr[0]
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
