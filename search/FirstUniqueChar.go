package main

// First unique string.
func mapUniqueChars(s string) map[rune][]int {
	// NOTE: `rune` := `int8` in Golang (8 bits := 1 bytes).
	strMap := make(map[rune][]int)
	for i, c := range s {
		strMap[c] = append(strMap[c], i)
	}
	return strMap
}

func detectPosUnique(s string) int {
	mapStr := mapUniqueChars(s)
	index := -1
	for _, r := range s {
		if len(mapStr[r]) == 1 {
			index = mapStr[r][0]
			break
		}
	}
	return index
}
