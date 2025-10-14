package sorts

import "fmt"

// BubbleSort call all bubble sort
func BubbleSort(array []int) {
	// Bubble sort and print
	fmt.Println("Classic", BubbleSortClassic(array))
	// Bubble sort sentinel and pring
	fmt.Println("Sentinel", BubbleSortSentinel(array))
	// Buuble sort sentinel do while
	fmt.Println("Sentinel Do While", BubbleSortSentinelDo(array))
}

// BubbleSortClassic order and int array
func BubbleSortClassic(array []int) []int {
	// From 1 to array len
	for i := 1; i < len(array); i++ {
		// From o to array len -1
		for j := 0; j < len(array)-i; j++ {
			// If the bubble j is bigger than j+1, then change the values...
			if array[j] > array[j+1] {
				aux := array[j]
				array[j] = array[j+1]
				array[j+1] = aux
			}
		}
	}
	// Response ordered array
	return array
}

// BubbleSortSentinel order an int array and check if it's already order
// to finish the process quickly
func BubbleSortSentinel(array []int) []int {
	i := 1
	ordered := false
	// For is golang while
	for i < len(array) && !ordered {
		i = i + 1
		ordered = true
		for j := 0; j < len(array)-i; j++ {
			if array[j] > array[j+1] {
				ordered = false
				aux := array[j]
				array[j] = array[j+1]
				array[j+1] = aux
			}
		}
	}
	// Response ordered array
	return array
}

// BubbleSortSentinelDo order an int array and check if it's already order
// to finish the process quickly
func BubbleSortSentinelDo(array []int) []int {
	i := 1
	for {
		i = i + 1
		ordered := true
		// Check the buble
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				ordered = false
				aux := array[j]
				array[j] = array[j+1]
				array[j+1] = aux
			}
		}
		if !(i < len(array)) || ordered {
			break
		}
	}
	// Response ordered array
	return array
}
