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

	m := map[bool]int{
		true:  0,
		false: 0,
	}

	for i := 1; i < limit; i++ {
		m[Flip()]++
	}

	if m[true] < threshold || m[false] < threshold {
		t.Errorf("Flip() true false not very balanced")
	}
}
