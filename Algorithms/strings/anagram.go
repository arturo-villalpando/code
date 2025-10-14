package strings

import (
	"strconv"
	"strings"
)

func FindAnagrams(s []string) [][]string {
	words := make(map[string][]string)
	for i := range s {
		hash := createStringHash(s[i])
		words[hash] = append(words[hash], s[i])
	}
	// Values
	result := make([][]string, 0, len(words))
	for _, value := range words {
		result = append(result, value)
	}
	// Return value...
	return result
}

func createStringHash(s string) string {
	array := make([]int, 26)
	for _, k := range s {
		array[int(k)-int('a')] += 1
	}
	return convertIntToString(array)
}

func convertIntToString(elems []int) string {
	var sb strings.Builder
	for _, v := range elems {
		sb.WriteString(strconv.Itoa(v))
	}
	return sb.String()
}
