/**
 * Author: Mitch Allen
 * File: flip_test.go
 */

package coin

import (
	"testing"
)

func TestFlip(t *testing.T) {
	if got := Flip(); got != true && got != false {
		t.Errorf("Flip() = %t, didn't return true or false", got)
	}
}

func TestFlipAverage(t *testing.T) {

	limit := 1000
	threshold := 400
	heads := 0
	tails := 0

	for i := 1; i < limit; i++ {
		if Flip() {
			heads++
		} else {
			tails++
		}
	}

	if heads < threshold || tails < threshold {
		t.Errorf("Flip() true false not very balanced")
	}
}
