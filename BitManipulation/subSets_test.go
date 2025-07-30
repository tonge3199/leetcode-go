package bitmanipulation

import (
	"sort"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name: "example case [1,2,3]",
			nums: []int{1, 2, 3},
			expected: [][]int{
				{},
				{1},
				{2},
				{1, 2},
				{3},
				{1, 3},
				{2, 3},
				{1, 2, 3},
			},
		},
		{
			name: "single element",
			nums: []int{1},
			expected: [][]int{
				{},
				{1},
			},
		},
		{
			name: "empty array",
			nums: []int{},
			expected: [][]int{
				{},
			},
		},
		{
			name: "two elements",
			nums: []int{4, 5},
			expected: [][]int{
				{},
				{4},
				{5},
				{4, 5},
			},
		},
		{
			name: "four elements",
			nums: []int{1, 2, 3, 4},
			expected: [][]int{
				{},
				{1},
				{2},
				{1, 2},
				{3},
				{1, 3},
				{2, 3},
				{1, 2, 3},
				{4},
				{1, 4},
				{2, 4},
				{1, 2, 4},
				{3, 4},
				{1, 3, 4},
				{2, 3, 4},
				{1, 2, 3, 4},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := subsets(tt.nums)

			// Verify the total number of subsets is 2^n
			expectedCount := 1 << len(tt.nums)
			if len(result) != expectedCount {
				t.Errorf("subsets(%v) returned %d subsets, want %d", tt.nums, len(result), expectedCount)
				return
			}

			// Convert to sets for comparison (order independent)
			if !subsetsEqual(result, tt.expected) {
				t.Errorf("subsets(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}

// subsetsEqual compares two slices of subsets for equality, ignoring order
func subsetsEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	// Convert slices to sets for comparison
	setA := make(map[string]bool)
	setB := make(map[string]bool)

	for _, subset := range a {
		sortedSubset := make([]int, len(subset))
		copy(sortedSubset, subset)
		sort.Ints(sortedSubset)
		key := subsetToString(sortedSubset)
		setA[key] = true
	}

	for _, subset := range b {
		sortedSubset := make([]int, len(subset))
		copy(sortedSubset, subset)
		sort.Ints(sortedSubset)
		key := subsetToString(sortedSubset)
		setB[key] = true
	}

	// Compare sets
	for key := range setA {
		if !setB[key] {
			return false
		}
	}

	for key := range setB {
		if !setA[key] {
			return false
		}
	}

	return true
}

// subsetToString converts a sorted subset to a string for use as map key
func subsetToString(subset []int) string {
	result := ""
	for i, num := range subset {
		if i > 0 {
			result += ","
		}
		result += string(rune(num + 1000)) // offset to avoid conflicts with small numbers
	}
	return result
}

// slicesEqual compares two int slices for equality
func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestSubsetsProperties(t *testing.T) {
	testCases := [][]int{
		{1, 2, 3},
		{4, 5, 6, 7},
		{0},
		{1, 2},
	}

	for _, nums := range testCases {
		t.Run("properties test", func(t *testing.T) {
			result := subsets(nums)

			// Property 1: Total number should be 2^n
			expectedCount := 1 << len(nums)
			if len(result) != expectedCount {
				t.Errorf("Expected %d subsets, got %d", expectedCount, len(result))
			}

			// Property 2: Should contain empty set
			hasEmptySet := false
			for _, subset := range result {
				if len(subset) == 0 {
					hasEmptySet = true
					break
				}
			}
			if !hasEmptySet {
				t.Errorf("Result should contain empty subset")
			}

			// Property 3: Should contain the full set
			hasFullSet := false
			for _, subset := range result {
				if len(subset) == len(nums) {
					sort.Ints(subset)
					sortedNums := make([]int, len(nums))
					copy(sortedNums, nums)
					sort.Ints(sortedNums)
					if slicesEqual(subset, sortedNums) {
						hasFullSet = true
						break
					}
				}
			}
			if !hasFullSet {
				t.Errorf("Result should contain the full set")
			}

			// Property 4: No duplicate subsets
			seen := make(map[string]bool)
			for _, subset := range result {
				sort.Ints(subset)
				key := ""
				for _, num := range subset {
					key += string(rune(num + '0'))
				}
				if seen[key] {
					t.Errorf("Found duplicate subset: %v", subset)
				}
				seen[key] = true
			}
		})
	}
}

func BenchmarkSubsets(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		subsets(nums)
	}
}

func BenchmarkSubsetsLarge(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		subsets(nums)
	}
}
