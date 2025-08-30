package array

import "testing"

func TestTrap(t *testing.T) {
	testCases := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "示例案例",
			height:   []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expected: 6,
		},
		{
			name:     "简单案例",
			height:   []int{3, 0, 2, 0, 4},
			expected: 7,
		},
		{
			name:     "递增序列",
			height:   []int{1, 2, 3, 4, 5},
			expected: 0,
		},
		{
			name:     "递减序列",
			height:   []int{5, 4, 3, 2, 1},
			expected: 0,
		},
		{
			name:     "空数组",
			height:   []int{},
			expected: 0,
		},
		{
			name:     "单个元素",
			height:   []int{5},
			expected: 0,
		},
		{
			name:     "两个元素",
			height:   []int{2, 5},
			expected: 0,
		},
		{
			name:     "山谷形状",
			height:   []int{3, 2, 0, 2, 3},
			expected: 4,
		},
		{
			name:     "复杂案例",
			height:   []int{4, 2, 0, 3, 2, 5},
			expected: 9,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Trap(tc.height)
			if result != tc.expected {
				t.Errorf("trap(%v) = %d, expected %d", tc.height, result, tc.expected)
			}
		})
	}
}

func BenchmarkTrap(b *testing.B) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Trap(height)
	}
}
