package utils

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/exp/constraints"
)

// MinVal compares 2 integer numbers to determine which one is smaller.
func MinVal[T constraints.Ordered](dst, src T) T {
	if dst < src {
		return dst
	}
	return src
}

// Itobytes converts an int number to a bytes slice.
func Itobytes(val int) []byte {
	return []byte(strconv.Itoa(val))
}

// Bytestoi converts a bytes slice to an int number.
func Bytestoi(val []byte) int {
	i, err := strconv.Atoi(string(val))
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %+v\n", err)
	}
	return i
}

// Strtoi converts a string to its integer representation.
func Strtoi(src string, bitSize int) int64 {
	const decBase int = 10
	destInt, err := strconv.ParseInt(src, decBase, bitSize)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %+v\n", err)
	}
	return destInt
}

// Contains returns true if slice `s` contains the given element `e`.
func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// Remove the given element `e` from slice `s`.
func Remove[T comparable](s []T, e T) []T {
	if !Contains(s, e) {
		return s
	}
	pos := IndexOf(s, e)
	return append(s[:pos], s[pos+1:]...)
}

// IndexOf returns the index of the first occurrence of the provided `e` in `s`.
func IndexOf[T comparable](s []T, e T) int {
	for pos, v := range s {
		if e == v {
			return pos
		}
	}
	return -1
}

// Unique returns a unique slice with no duplicated values.
func Unique[T comparable](s []T) []T {
	unqMap := make(map[T]bool)
	var res []T
	for _, v := range s {
		if _, ok := unqMap[v]; !ok {
			unqMap[v] = true
			res = append(res, v)
		}
	}
	return res
}

// ReverseBytes reverse the order of a byte slice.
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
