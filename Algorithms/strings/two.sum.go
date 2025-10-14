package strings

// Sum to values inside the array to equals to target
func TwoSum(nums []int, target int) []int {
	// Create a complement map to know how it works, example:
	// Find [1,2,3,9], target = 8
	// The complement of one is 8 - 1 = 7, 8-2 = 6and so on...
	//
	complements := make(map[int]int)
	for i, value := range nums {
		// Find if the complement is already on the map
		val, ok := complements[value]
		if ok {
			return []int{val, i}
		}
		// Ad the complement, that is target - value.
		// Then for example for 8-4 = 0 -> complement-> response = 4....
		complements[target-value] = i
	}
	return []int{}
}
