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

	"github.com/spaolacci/murmur3"
)

var Filter *BloomFilter

type BloomFilter struct {
	// size of the bloom filter
	size uint

	// number of hash functions
	hashFunctions int

	// bit array to store the bloom filter
	bitArray []bool // TODO: use a bit array instead of a slice of bools

	// hash functions to use
	hashFuncList []hash.Hash32
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
func New(size, hashFunctions int) *BloomFilter {
	hashFuncList := make([]hash.Hash32, hashFunctions)
	for i := range hashFunctions {
		hashFuncList[i] = murmur3.New32WithSeed(rand.Uint32())
	}
	Filter = &BloomFilter{
		size:          uint(size),
		hashFunctions: hashFunctions,
		bitArray:      make([]bool, size),
		hashFuncList:  hashFuncList,
	}

	return Filter
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
	idx := b.doHash(item)
	for _, index := range idx {
		b.bitArray[index] = true
	}
}

// Check Checks if an item is in the bloom filter
// parameters:
//
//	item	: item to check in the bloom filter
//
// returns:
//
//	bool	: true if the item is in the bloom filter, false otherwise
func (b *BloomFilter) Check(item string) bool {
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
	b.bitArray = make([]bool, b.size)
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
	idx := make([]uint, 0, b.hashFunctions)

	for _, hashFunc := range b.hashFuncList {
		hashFunc.Reset()
		hashFunc.Write([]byte(input))
		hashValue := hashFunc.Sum32()
		index := uint(hashValue) % b.size
		idx = append(idx, index)
	}

	return idx
}

func (b *BloomFilter) String() string {
	return fmt.Sprintf("BloomFilter{size: %d, hashFunctions: %d, bitArray: %v}", b.size, b.hashFunctions, b.bitArray)
}
