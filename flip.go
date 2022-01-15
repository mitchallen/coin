/**
 * Author: Mitch Allen
 * File: flip.go
 */

package coin

import (
	"math/rand"
	"time"
)

// Flip returns a random boolean true or false
func Flip() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Float32() > 0.5
}
