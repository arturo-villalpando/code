package strings

import (
	"reflect"
	"testing"
)

func TestZeroMatrix(t *testing.T) {
	// First test
	firstMatrix := [][]int{
		{2, 1, 3, 0, 2},
		{7, 4, 1, 3, 8},
		{1, 0, 1, 2, 1},
		{9, 3, 4, 1, 9},
	}
	firstMatrixAnswer := [][]int{
		{0, 0, 0, 0, 0},
		{7, 0, 1, 0, 8},
		{0, 0, 0, 0, 0},
		{9, 0, 4, 0, 9},
	}
	// Execute the verifications, firstMatrix is the result
	ZeroMatrix(firstMatrix)
	if !reflect.DeepEqual(firstMatrix, firstMatrixAnswer) {
		t.Errorf("Result was incorrect, got: %v, want: %v", firstMatrix, firstMatrixAnswer)
	}
	// Second test
	secondMatrix := [][]int{
		{2, 0, 2},
		{0, 2, 1},
		{9, 3, 4},
	}
	secondMatrixAnswer := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 4},
	}
	ZeroMatrix(secondMatrix)
	if !reflect.DeepEqual(secondMatrix, secondMatrixAnswer) {
		t.Errorf("Result was incorrect, got: %v, want: %v", secondMatrix, secondMatrixAnswer)
	}
}
