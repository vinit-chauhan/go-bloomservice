package bloom

import (
	"testing"

	"github.com/vinit-chauhan/go-bloomservice/test"
)

func BenchmarkBloomFilter_FalsePositives(b *testing.B) {
	n := 100_000     // Number of items to add
	testN := 100_000 // Number of items to test
	size := n * 20   // Size of the bloom filter
	k := 5           // Number of hash functions
	filter := New(size, k)

	items := test.GenerateStringsOfLength(10, n)
	for _, item := range items {
		filter.Add(item)
	}

	falsePositives := 0
	testItems := test.GenerateStringsOfLength(10, testN)
	for _, item := range testItems {
		if filter.Exists(item) {
			falsePositives++
		}
	}

	b.ReportMetric(float64(falsePositives), "false_positives")
	falsePositivePct := float64(falsePositives) / float64(testN) * 100
	b.Logf("False positivity rate: %.2f%%", falsePositivePct)

	// Ensure the false positive rate is within acceptable limits
	// For a well-configured bloom filter, this should be less than 1%
	if falsePositivePct > 0.1 {
		b.Fatalf("False positivity rate exceeded 1%%: %.2f%%", falsePositivePct)
	}
	b.ResetTimer()
}
