/**
 * Author: Mitch Allen
 * File: flip.go
 */

package coin

import (
	"math/rand"
)

// Flip returns a random boolean true or false
func Flip() bool {
	return rand.Float32() > 0.5
}
