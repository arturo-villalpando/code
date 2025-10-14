package strings

import "strings"

func IsUnique(s string) bool {
	if len(s) > 128 {
		return false
	}
	// split the string into array
	arr := strings.Split(s, "")
	// Creat array to know if is already work in this character
	used := make(map[string]string)
	// Run first for
	for i := range arr {
		_, ok := used[arr[i]]
		if ok {
			return false
		} else {
			used[arr[i]] = arr[i]
		}
	}
	return true
}
