package recursive

import (
	"reflect"
	"sort"
	"testing"
)

func TestCombine(t *testing.T) {
	testCases := []struct {
		n        int
		k        int
		expected [][]int
	}{
		{
			n: 4,
			k: 2,
			expected: [][]int{
				{1, 2}, {1, 3}, {1, 4},
				{2, 3}, {2, 4},
				{3, 4},
			},
		},
		{
			n:        1,
			k:        1,
			expected: [][]int{{1}},
		},
		{
			n: 3,
			k: 3,
			expected: [][]int{
				{1, 2, 3},
			},
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := Combine(tc.n, tc.k)

			// 对结果排序以便比较
			sortCombinations(result)
			sortCombinations(tc.expected)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Combine(%d, %d) = %v, expected %v", tc.n, tc.k, result, tc.expected)
			}
		})
	}
}

// sortCombinations 对组合数组进行排序，便于测试比较
func sortCombinations(combinations [][]int) {
	// 首先对每个组合内部排序
	for _, combo := range combinations {
		sort.Ints(combo)
	}

	// 然后对组合之间排序
	sort.Slice(combinations, func(i, j int) bool {
		a, b := combinations[i], combinations[j]
		for k := 0; k < len(a) && k < len(b); k++ {
			if a[k] != b[k] {
				return a[k] < b[k]
			}
		}
		return len(a) < len(b)
	})
}
