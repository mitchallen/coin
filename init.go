/**
 * Author: Mitch Allen
 * File: init.go
 */

package coin

import (
	"math/rand"
	"time"
)

// Seed rand
func init() {
	rand.Seed(time.Now().UnixNano())
}
