package utils

import (
	"math/big"
	"math/rand"
	"time"
)

func RandomBigInt(min, max *big.Int) *big.Int {
	// Calculate the range (max - min + 1)
	diff := new(big.Int).Sub(max, min)
	diff.Add(diff, big.NewInt(1))

	// Generate a random number in range [0, diff-1]
	result := new(big.Int)
	result.Rand(rand.New(rand.NewSource(time.Now().UnixNano())), diff)

	// Add min to shift to the desired range [min, max]
	result.Add(result, min)

	return result
}
