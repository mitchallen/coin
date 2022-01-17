/**
 * Author: Mitch Allen
 * File: weighted.go
 */

package coin

import (
	"math/rand"
)

// Weighted returns a random boolean true or false based on weight (0.0..<1.0)
func Weighted(weight float32) bool {
	return rand.Float32() <= weight
}
