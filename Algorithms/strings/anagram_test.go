package strings

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	words := []string{"listen", "silent", "programador", "recluta", "magia", "maiga"}
	wordsResponse := [][]string{
		{"listen", "silent"},
		{"programador"},
		{"recluta"},
		{"magia", "maiga"},
	}
	result := FindAnagrams(words)
	// Normalize matrix
	normalizeMatrix(wordsResponse)
	normalizeMatrix(result)
	// Compare the response with reflect
	if !reflect.DeepEqual(result, wordsResponse) {
		t.Errorf("Result was incorrect, got: %v, want: %v", result, wordsResponse)
	}
}

func normalizeMatrix(matrix [][]string) {
	// Sort each submatrix (anagrams group)
	for _, key := range matrix {
		sort.Strings(key)
	}
	// Order the matrix accord the first element of each submatrix
	sort.Slice(matrix, func(i, j int) bool {
		return matrix[i][0] < matrix[j][0]
	})
}
