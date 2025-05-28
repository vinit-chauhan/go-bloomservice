// Copyright 2025 Vinit Chauhan

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bloom

import (
	"fmt"
	"hash"
	"math/rand"
	"sync"

	"github.com/spaolacci/murmur3"
)

var Filter *BloomFilter

type Parameters struct {
	// Size of the bloom filter in bits
	Size uint `json:"size"`
	// Number of hash functions used
	NumHashFunctions uint8 `json:"num_hash_functions"`
	// Estimated false positive rate
	FalsePositiveRate float64 `json:"false_positive_rate"`
}

type Statistics struct {
	// Number of items added to the bloom filter
	AddedItems uint64 `json:"added_items"`
	// Number of items checked in the bloom filter
	CheckedItems uint64 `json:"checked_items"`
}

type BloomFilter struct {
	// parameters for the bloom filter
	params Parameters

	// statistics for the bloom filter
	stats Statistics

	// bit array to store the bloom filter
	bitArray []bool // TODO: use a bit array instead of a slice of bools

	// hash functions to use
	hashFuncList []hash.Hash32

	// TODO: use mutex for concurrent access
	// mutex for concurrent access
	mu *sync.RWMutex

	// wait group for concurrent operations
	wg *sync.WaitGroup
}

// New Creates a new BloomFilter based on the size and number of hash functions
// parameters:
//
//	size			: number of bits in the bit array
//	hashFunctions	: number of times to hash the input
//
// returns:
//
//	*BloomFilter	: pointer to the BloomFilter struct
func New(params Parameters) *BloomFilter {
	// create hash functions
	hashFuncList := make([]hash.Hash32, params.NumHashFunctions)
	for i := range params.NumHashFunctions {
		hashFuncList[i] = murmur3.New32WithSeed(rand.Uint32())
	}

	return &BloomFilter{
		params:       params,
		bitArray:     make([]bool, params.Size),
		hashFuncList: hashFuncList,
		stats:        Statistics{},
		mu:           &sync.RWMutex{},
		wg:           &sync.WaitGroup{},
	}
}

func Init(estimatedKeyCount int, falsePositivePct float64) {
	if Filter == nil {
		params := CalculateOptimalParameters(estimatedKeyCount, falsePositivePct)
		Filter = New(params)
	}
}

// Add Adds an item to the bloom filter
// parameters:
//
//	item	: item to add to the bloom filter
//
// returns:
//
//	none
func (b *BloomFilter) Add(item string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.stats.AddedItems++
	idx := b.doHash(item)
	for _, index := range idx {
		b.bitArray[index] = true
	}
}

// Exists Checks if an item is in the bloom filter
// parameters:
//
//	item	: item to check in the bloom filter
//
// returns:
//
//	bool	: true if the item is in the bloom filter, false otherwise
func (b *BloomFilter) Exists(item string) bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	b.stats.CheckedItems++
	idx := b.doHash(item)
	for _, index := range idx {
		if !b.bitArray[index] {
			return false
		}
	}
	return true
}

// Clear Clears the bloom filter
// parameters:
//
//	none
//
// returns:
//
//	none
func (b *BloomFilter) Clear() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.bitArray = make([]bool, b.params.Size)
}

// doHash Hashes the input string using the hash functions
// parameters:
//
//	input			: input string to hash
//	hashFunctions	: number of hash functions to use
//
// returns:
//
//	[]int	: slice of indices to set in the bit array
func (b *BloomFilter) doHash(input string) []uint {
	idx := make([]uint, 0, b.params.NumHashFunctions)

	for _, hashFunc := range b.hashFuncList {
		hashFunc.Reset()
		hashFunc.Write([]byte(input))
		hashValue := hashFunc.Sum32()
		index := uint(hashValue) % b.params.Size
		idx = append(idx, index)
	}

	return idx
}

func (b *BloomFilter) String() string {
	return fmt.Sprintf("BloomFilter{size: %d, hashFunctions: %d, bitArray: %v}", b.params.Size, b.params.NumHashFunctions, b.bitArray)
}

func (b *BloomFilter) GetParameters() Parameters {
	return b.params
}

func (b *BloomFilter) GetStatistics() Statistics {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.stats
}
