package bloom

import (
	"testing"

	"github.com/vinit-chauhan/go-bloomservice/test"
)

func TestNew(t *testing.T) {
	params := CalculateOptimalParameters(100, 0.01)
	filter := New(params)
	if len(filter.bitArray) != int(params.Size) {
		t.Errorf("Expected size %d, got %d", params.Size, len(filter.bitArray))
	}
	if len(filter.hashFuncList) != int(params.NumHashFunctions) {
		t.Errorf("Expected %d hash functions, got %d", params.NumHashFunctions, len(filter.hashFuncList))
	}
}

func TestAddAndExists(t *testing.T) {
	items := []string{"apple", "banana", "cherry", "date", "elderberry"}
	nonItems := []string{"fig", "grape", "honeydew", "kiwi", "lemon"}

	params := CalculateOptimalParameters(len(items), 0.01)
	filter := New(params)

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

	params := CalculateOptimalParameters(n, 0.01)
	filter := New(params)

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
	items := test.GenerateStringsOfLength(10, 100)
	params := CalculateOptimalParameters(len(items), 0.01)
	filter := New(params)

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
