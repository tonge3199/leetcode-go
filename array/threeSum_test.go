package array

import (
	"reflect"
	"sort"
	"testing"
)

// Helper function to sort a slice of integer slices.
// This is crucial for comparing the expected and actual results,
// as the order of triplets or the order of numbers within triplets might vary
// but still be semantically correct.
func sortResult(result [][]int) [][]int {
	// Sort each inner slice (triplet)
	for _, triplet := range result {
		sort.Ints(triplet)
	}

	// Sort the outer slice based on the first element, then second, then third
	sort.Slice(result, func(i, j int) bool {
		for k := 0; k < len(result[i]); k++ {
			if result[i][k] != result[j][k] {
				return result[i][k] < result[j][k]
			}
		}
		return false // They are identical
	})
	return result
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "Example case",
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, 0, 1}, {-1, -1, 2}},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: [][]int{},
		},
		{
			name:     "Array with less than 3 elements",
			nums:     []int{0, 0},
			expected: [][]int{},
		},
		{
			name:     "No solution",
			nums:     []int{1, 2, 3, 4, 5},
			expected: [][]int{},
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			name:     "Duplicates, multiple solutions",
			nums:     []int{-2, 0, 0, 2, 2},
			expected: [][]int{{-2, 0, 2}},
		},
		{
			name:     "Mixed positives and negatives",
			nums:     []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4},
			expected: [][]int{{-4, 0, 4}, {-4, 1, 3}, {-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:     "Single solution",
			nums:     []int{-5, -1, 0, 1, 6},
			expected: [][]int{{-5, -1, 6}, {-1, 0, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)

			// Sort both the actual and expected results for consistent comparison
			sortedGot := sortResult(got)
			sortedExpected := sortResult(tt.expected)

			if !reflect.DeepEqual(sortedGot, sortedExpected) {
				t.Errorf("threeSum(%v) got = %v, want %v", tt.nums, sortedGot, sortedExpected)
			}
		})
	}
}
