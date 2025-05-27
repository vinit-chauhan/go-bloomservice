package bloom

import "math"

func CalculateSize(n, p float64) int {
	// Calculate the size of the bloom filter using the formula:
	// size = ceil((n * log(p)) / log(1 / pow(2, log(2))));
	// Where,
	// 		n : number of items to be added
	// 		p : desired false positive rate.

	size := math.Ceil((n * math.Log(p))) / math.Log((1 / math.Pow(2, math.Log(2))))

	return int(size)
}

func CalculateNumHashFunctions(size, n int) int {
	// Calculate the number of hash functions using the formula:
	// k = ceil((size / n) * log(2))
	// Where,
	// 		size : size of the bloom filter
	// 		n : number of items to be added

	if n == 0 {
		return 0
	}

	k := math.Ceil((float64(size) / float64(n)) * math.Log(2))

	return int(k)
}

func CalculateOptimalParameters(n int, p float64) (int, int) {
	size := CalculateSize(float64(n), p)
	return size, CalculateNumHashFunctions(size, n)
}
