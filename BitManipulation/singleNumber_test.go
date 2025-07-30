package bitmanipulation

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "example 1: [2,2,1]",
			nums:     []int{2, 2, 1},
			expected: 1,
		},
		{
			name:     "example 2: [4,1,2,1,2]",
			nums:     []int{4, 1, 2, 1, 2},
			expected: 4,
		},
		{
			name:     "single element",
			nums:     []int{42},
			expected: 42,
		},
		{
			name:     "negative numbers",
			nums:     []int{-1, -1, -2, -2, 3},
			expected: 3,
		},
		{
			name:     "large numbers",
			nums:     []int{1000000, 999999, 1000000, 999999, 123456},
			expected: 123456,
		},
		{
			name:     "zero included",
			nums:     []int{0, 1, 0, 1, 5},
			expected: 5,
		},
		{
			name:     "single number is zero",
			nums:     []int{7, 7, 0},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := singleNumber(tt.nums)
			if result != tt.expected {
				t.Errorf("singleNumber(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}

// TestXORProperties tests the mathematical properties of XOR operation
func TestXORProperties(t *testing.T) {
	testCases := []struct {
		name string
		test func() bool
	}{
		{
			name: "x ^ x = 0",
			test: func() bool {
				values := []int{1, 5, 100, -50, 0}
				for _, x := range values {
					if x^x != 0 {
						return false
					}
				}
				return true
			},
		},
		{
			name: "x ^ 0 = x",
			test: func() bool {
				values := []int{1, 5, 100, -50, 42}
				for _, x := range values {
					if x^0 != x {
						return false
					}
				}
				return true
			},
		},
		{
			name: "commutative property: a ^ b = b ^ a",
			test: func() bool {
				pairs := [][2]int{{1, 2}, {5, 7}, {-3, 4}, {0, 10}}
				for _, pair := range pairs {
					a, b := pair[0], pair[1]
					if a^b != b^a {
						return false
					}
				}
				return true
			},
		},
		{
			name: "associative property: (a ^ b) ^ c = a ^ (b ^ c)",
			test: func() bool {
				triplets := [][3]int{{1, 2, 3}, {5, 7, 9}, {-1, 0, 4}}
				for _, triplet := range triplets {
					a, b, c := triplet[0], triplet[1], triplet[2]
					if (a^b)^c != a^(b^c) {
						return false
					}
				}
				return true
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if !tc.test() {
				t.Errorf("XOR property test failed: %s", tc.name)
			}
		})
	}
}

// TestXORBitLevel demonstrates XOR operation at bit level
func TestXORBitLevel(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{
			name:     "5 ^ 3 = 6 (0101 ^ 0011 = 0110)",
			a:        5, // 0101
			b:        3, // 0011
			expected: 6, // 0110
		},
		{
			name:     "12 ^ 8 = 4 (1100 ^ 1000 = 0100)",
			a:        12, // 1100
			b:        8,  // 1000
			expected: 4,  // 0100
		},
		{
			name:     "7 ^ 7 = 0 (same numbers)",
			a:        7,
			b:        7,
			expected: 0,
		},
		{
			name:     "15 ^ 0 = 15 (XOR with zero)",
			a:        15,
			b:        0,
			expected: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.a ^ tt.b
			if result != tt.expected {
				t.Errorf("%d ^ %d = %d, want %d", tt.a, tt.b, result, tt.expected)
				// Print binary representation for debugging
				t.Logf("Binary: %08b ^ %08b = %08b (expected %08b)",
					tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestSingleNumberPerformance tests with larger arrays
func TestSingleNumberPerformance(t *testing.T) {
	// Create a large array where all numbers appear twice except one
	size := 10000
	nums := make([]int, size*2+1)

	// Fill with pairs
	for i := 0; i < size; i++ {
		nums[i*2] = i + 1
		nums[i*2+1] = i + 1
	}

	// Add the single number
	singleNum := 99999
	nums[len(nums)-1] = singleNum

	result := singleNumber(nums)
	if result != singleNum {
		t.Errorf("singleNumber() = %d, want %d", result, singleNum)
	}
}

func BenchmarkSingleNumber(b *testing.B) {
	nums := []int{4, 1, 2, 1, 2, 3, 3, 5, 5, 6, 6}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		singleNumber(nums)
	}
}

func BenchmarkSingleNumberLarge(b *testing.B) {
	// Create array with 1000 pairs + 1 single number
	size := 1000
	nums := make([]int, size*2+1)

	for i := 0; i < size; i++ {
		nums[i*2] = i
		nums[i*2+1] = i
	}
	nums[len(nums)-1] = 12345 // single number

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		singleNumber(nums)
	}
}
