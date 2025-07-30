package hashMap

import (
	"hash/fnv"
	"math"
)

// Bit array
// Multiple hash functions

type BloomFilter struct {
	m    uint64 // Size of bit array (number of bits)
	k    uint64 // Number of hash functions
	bits []byte // Bit array
}

func NewBloomFilter(expectedN uint64, falsePositiveRate float64) *BloomFilter {
	m, k := estimateParams(expectedN, falsePositiveRate)
	if m == 0 || k == 0 {
		panic("invalid parameters for NewBloomFilter: m or k is 0")
	}
	return NewBloomFilterWithMK(m, k)
}

func NewBloomFilterWithMK(m, k uint64) *BloomFilter {
	if m == 0 || k == 0 {
		panic("Invalid parameters for Bloom filter: m or k is zero")
	}
	// 整数除法（/），会自动向下取整
	numBytes := (m + 7) / 8
	return &BloomFilter{
		m:    m,
		k:    k,
		bits: make([]byte, numBytes),
	}
}

func estimateParams(n uint64, p float64) (m, k uint64) {
	if n == 0 || p <= 0 || p >= 1 {
		return 1000, 10
	}
	mFloat := -(float64(n) * math.Log(p)) / (math.Ln2 * math.Ln2)
	kFloat := (mFloat / float64(n)) * math.Ln2

	m = uint64(math.Ceil(mFloat))
	k = uint64(math.Floor(kFloat))

	if k < 1 {
		k = 1
	}
	return
}

// gitHashes uses double hashing to generate k hash values for the data
func (bf *BloomFilter) getHashes(data []byte) []uint64 {
	hashes := make([]uint64, bf.k)

	// Use two different versions (or seeds) of FNV-1a as base hash functions
	h1 := fnv.New64a()
	h1.Write(data)
	hash1Val := h1.Sum64()

	h2 := fnv.New64a()
	h2.Write(data)
	hash2Val := h2.Sum64()

	for i := uint64(0); i < bf.k; i++ {
		if hash2Val == 0 && i > 0 {
			hash2Val = 1
		}
		hashes[i] = (hash1Val + i*hash2Val) % bf.m
	}
	return hashes
}

// Add Insert data into the Bloom filter
func (bf *BloomFilter) Add(data []byte) {
	hashes := bf.getHashes(data)
	for _, hash := range hashes {
		byteIndex := hash / 8
		bitOffset := hash % 8
		// TODO:explain
		bf.bits[byteIndex] |= 1 << bitOffset
	}
}

// MightContain Check whether data might exist
func (bf *BloomFilter) MightContain(data []byte) bool {
	hashes := bf.getHashes(data)
	for _, hash := range hashes {
		byteIndex := hash / 8
		bitOffset := hash % 8
		if bf.bits[byteIndex]&(1<<bitOffset) == 0 {
			// If any bit corresponding to a hash is 0, the element definitely does not exist
			return false
		}
	}
	// If all bits corresponding to hashes are 1, the element might exist
	return true
}

// Reset clears the Bloom filter (sets all bits to 0)
func (bf *BloomFilter) Reset() {
	for i := range bf.bits {
		bf.bits[i] = 0
	}
}
