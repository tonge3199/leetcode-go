package hashMap

import (
	"testing"
)

// TestBloomFilter tests the core functionality of the Bloom filter
func TestBloomFilter(t *testing.T) {
	// Setup: Create a Bloom filter with expected 1000 elements and 1% false positive rate
	expectedN := uint64(1000)
	falsePositiveRate := 0.01

	// Test case 1: Verify estimated parameters
	m, k := estimateParams(expectedN, falsePositiveRate)
	if m == 0 || k == 0 {
		t.Errorf("estimateParameters(%d, %f) returned invalid parameters: m=%d, k=%d", expectedN, falsePositiveRate, m, k)
	}
	t.Logf("Estimated parameters: m = %d, k = %d", m, k)

	// Test case 2: Create Bloom filter and test Add/MightContain
	bf := NewBloomFilter(expectedN, falsePositiveRate)

	// Define test items
	item1 := []byte("apple")
	item2 := []byte("banana")
	item3 := []byte("cherry")
	item4 := []byte("grape")

	// Add elements
	bf.Add(item1)
	bf.Add(item2)

	// Test membership
	tests := []struct {
		name     string
		item     []byte
		expected bool
	}{
		{"apple (added)", item1, true},
		{"banana (added)", item2, true},
		{"cherry (not added)", item3, false},
		{"grape (not added)", item4, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := bf.MightContain(tt.item)
			if result != tt.expected {
				t.Errorf("MightContain(%s) = %t; want %t", tt.name, result, tt.expected)
			}
		})
	}

	// Test case 3: Test Reset functionality
	bf.Reset()
	t.Run("After Reset", func(t *testing.T) {
		if bf.MightContain(item1) {
			t.Errorf("MightContain('apple') after Reset = true; want false")
		}
	})
}
