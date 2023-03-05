package main

import (
	"fmt"
	"math"
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

// NOTE: `pow(x, n)`.
// The time complexity: `O(log n)`
func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	// NOTE: `pow(x, -n)` := `1 / (x^n)`.
	if n < 0 {
		n = -1 * n
		x = 1 / x
	}
	if n%2 == 0 {
		return pow(x*x, n/2) // `(x^2)^(n/2)` := `x^(2*n/2)` := `x^n`.
	} else {
		return x * pow(x, n-1) // `x * x^(n-1)` := `x^n`.
	}
}

func main() {
	convertToASCII(0x80)

	arr := []int{1, 2, 3, 4, 5}
	sum := SumDAC(arr)
	fmt.Println(sum) // 15.

	fmt.Println(len(arr) == CountDAC(arr), len(arr) == CountElement(arr))

	fmt.Println(greatestCommonDivisor(225, 125)) // 25.

	fmt.Println(pow(-2, 10) == math.Pow(-2, 10)) // 1024.
}
