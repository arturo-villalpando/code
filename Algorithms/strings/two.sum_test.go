package strings

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	// First test
	first := []int{9, 2, 5, 6}
	firstResponse := []int{1, 2}
	result := TwoSum(first, 7)
	// Compare responses
	if !reflect.DeepEqual(result, firstResponse) {
		t.Errorf("Result was incorrect, got: %v, want: %v", result, firstResponse)
	}
	// Second test
	second := []int{9, 2, 5, 6}
	secondResponse := []int{}
	resultSecond := TwoSum(second, 100)
	// Compare responses
	if !reflect.DeepEqual(resultSecond, secondResponse) {
		t.Errorf("Result was incorrect, got: %v, want: %v", resultSecond, secondResponse)
	}
}
