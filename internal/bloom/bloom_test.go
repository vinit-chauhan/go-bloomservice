package bloom

import (
	"testing"

	"github.com/vinit-chauhan/go-bloomservice/test"
)

func TestNew(t *testing.T) {
	filter := New(100, 5)
	if filter.size != 100 {
		t.Errorf("Expected size 100, got %d", filter.size)
	}
	if filter.hashFunctions != 5 {
		t.Errorf("Expected 5 hash functions, got %d", filter.hashFunctions)
	}
	if len(filter.bitArray) != 100 {
		t.Errorf("Expected bit array length 100, got %d", len(filter.bitArray))
	}
	if len(filter.hashFuncList) != 5 {
		t.Errorf("Expected 5 hash functions in list, got %d", len(filter.hashFuncList))
	}
}

func TestAddAndExists(t *testing.T) {
	items := []string{"apple", "banana", "cherry", "date", "elderberry"}
	nonItems := []string{"fig", "grape", "honeydew", "kiwi", "lemon"}

	size, hashes := CalculateOptimalParameters(len(items), 0.01)
	filter := New(size, hashes)

	for _, item := range items {
		filter.Add(item)
	}

	for _, item := range items {
		if !filter.Exists(item) {
			t.Fatalf("Expected item %s to exist in the filter", item)
		}
	}

	for _, item := range nonItems {
		if filter.Exists(item) {
			t.Fatalf("Expected item %s to NOT exist in the filter", item)
		}
	}
}

func TestAddAndExistsRandomStrings(t *testing.T) {
	n := 100
	items := test.GenerateStringsOfLength(10, n)

	size, hashes := CalculateOptimalParameters(n, 0.01)
	filter := New(size, hashes)

	for _, item := range items {
		filter.Add(item)
	}

	for _, item := range items {
		if !filter.Exists(item) {
			t.Fatalf("Expected item %s to exist in the filter", item)
		}
	}
}

func TestClear(t *testing.T) {
	filter := New(100_000, 5)
	items := test.GenerateStringsOfLength(10, 100)

	for _, item := range items {
		filter.Add(item)
	}

	filter.Clear()

	for bit := range filter.bitArray {
		if filter.bitArray[bit] {
			t.Fatalf("Expected bit array to be cleared, but bit %d is still set", bit)
		}
	}
}
