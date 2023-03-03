package main

func SumDAC[number int | uint | float64](arr []number) number {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + SumDAC(arr[1:])
}

func CountDAC[number int | uint | float64](arr []number) number {
	if len(arr) == 0 {
		return 0
	}
	return 1 + CountDAC(arr[1:])
}

func CountElement[number int | uint | float64](arr []number) number {
	var counter number = 0
	for i := range arr {
		if &arr[i] != nil {
			counter++
		}
	}
	return counter
}

func FindMax(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 {
		return arr[0]
	}
	max := arr[0]
	return max
}
