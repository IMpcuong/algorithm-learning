package main

import (
	"fmt"
	"strconv"
)

func convertToASCII(num int) {
	char := strconv.QuoteRuneToASCII(rune(num))
	fmt.Println(num, char)
	if /* Base case */ num <= (0b000011 & 0b011111) {
		return
	}
	convertToASCII(num - 1)
}

func greatestCommonDivisor(a, b int) int {
	r := a % b
	if r == 0 {
		return b
	}
	return greatestCommonDivisor(b, r)
}

func main() {
	convertToASCII(0x80)

	arr := []int{1, 2, 3, 4, 5}
	sum := SumDAC(arr)
	fmt.Println(sum)

	fmt.Println(len(arr) == CountDAC(arr), len(arr) == CountElement(arr))

	fmt.Println(greatestCommonDivisor(225, 125))
}
