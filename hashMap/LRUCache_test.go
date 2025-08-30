package hashMap

import (
	"testing"
)

func TestLRUCache_BasicOperations(t *testing.T) {
	// Test case 1: Basic get/put operations
	lru := New(2)

	// Test put operation
	lru.put(1, 100)
	lru.put(2, 200)

	// Test get operation
	result := lru.get(1)
	if result != 100 {
		t.Errorf("Expected 100, got %d", result)
	}

	result = lru.get(2)
	if result != 200 {
		t.Errorf("Expected 200, got %d", result)
	}

	// Test non-existent key
	result = lru.get(3)
	if result != -1 {
		t.Errorf("Expected -1 for non-existent key, got %d", result)
	}
}

func TestLRUCache_CapacityEviction(t *testing.T) {
	// Test case 2: Capacity eviction
	lru := New(2)

	lru.put(1, 100)
	lru.put(2, 200)

	// Access key 1 to make it recently used
	lru.get(1)

	// Add new key, should evict key 2 (least recently used)
	lru.put(3, 300)

	// Key 2 should be evicted
	result := lru.get(2)
	if result != -1 {
		t.Errorf("Expected key 2 to be evicted, but got %d", result)
	}

	// Key 1 and 3 should still exist
	result = lru.get(1)
	if result != 100 {
		t.Errorf("Expected 100 for key 1, got %d", result)
	}

	result = lru.get(3)
	if result != 300 {
		t.Errorf("Expected 300 for key 3, got %d", result)
	}
}

func TestLRUCache_UpdateExistingKey(t *testing.T) {
	// Test case 3: Update existing key
	lru := New(2)

	lru.put(1, 100)
	lru.put(2, 200)

	// Update existing key
	lru.put(1, 150)

	result := lru.get(1)
	if result != 150 {
		t.Errorf("Expected 150 after update, got %d", result)
	}
}

func TestLRUCache_LRUOrder(t *testing.T) {
	// Test case 4: LRU order maintenance
	lru := New(3)

	lru.put(1, 100)
	lru.put(2, 200)
	lru.put(3, 300)

	// Access in order: 1, 2, 3 (3 is most recent, 1 is least recent)
	lru.get(1)
	lru.get(2)
	lru.get(3)

	// Add new key, should evict least recently used (key 1)
	lru.put(4, 400)

	// Key 1 should be evicted
	result := lru.get(1)
	if result != -1 {
		t.Errorf("Expected key 1 to be evicted, got %d", result)
	}

	// Keys 2, 3, 4 should exist
	expectedKeys := map[int]int{2: 200, 3: 300, 4: 400}
	for key, expectedValue := range expectedKeys {
		result := lru.get(key)
		if result != expectedValue {
			t.Errorf("Expected %d for key %d, got %d", expectedValue, key, result)
		}
	}
}

func TestLRUCache_EdgeCases(t *testing.T) {
	// Test case 5: Edge cases

	// Capacity 1
	lru := New(1)
	lru.put(1, 100)

	result := lru.get(1)
	if result != 100 {
		t.Errorf("Expected 100, got %d", result)
	}

	// Add another key, should evict the first
	lru.put(2, 200)

	result = lru.get(1)
	if result != -1 {
		t.Errorf("Expected key 1 to be evicted, got %d", result)
	}

	result = lru.get(2)
	if result != 200 {
		t.Errorf("Expected 200 for key 2, got %d", result)
	}
}

// Benchmark test
func BenchmarkLRUCache_Operations(b *testing.B) {
	lru := New(1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := i % 1000
		lru.put(key, i)
		lru.get(key)
	}
}

// Test to identify specific bugs in current implementation
func TestLRUCache_IdentifyBugs(t *testing.T) {
	t.Log("Testing current implementation to identify bugs...")

	lru := New(2)

	// This will likely fail due to bugs in the current implementation
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Panic occurred: %v", r)
		}
	}()

	// Test basic put
	t.Log("Testing put operation...")
	lru.put(1, 100)

	// Test basic get
	t.Log("Testing get operation...")
	result := lru.get(1)
	t.Logf("Get result: %d", result)

	// Test capacity overflow
	t.Log("Testing capacity overflow...")
	lru.put(2, 200)
	lru.put(3, 300) // Should evict something

	t.Log("Current implementation test completed")
}
