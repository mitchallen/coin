/**
 * Author: Mitch Allen
 * File: flip_test.go
 */

package coin

import (
	"testing"
)

func TestWeighted(t *testing.T) {
	if got := Weighted(0.75); got != true && got != false {
		t.Errorf("Weighted() = %t, didn't return true or false", got)
	}
}

func TestWeightAverage50(t *testing.T) {

	limit := 1000
	threshold := 400
	heads := 0
	tails := 0

	for i := 1; i < limit; i++ {
		if Weighted(0.5) {
			heads++
		} else {
			tails++
		}
	}

	if heads < threshold || tails < threshold {
		t.Errorf("Weighted(0.5) true false not very balanced")
	}
}

func TestWeightAverage75(t *testing.T) {

	limit := 1000
	heads := 0
	tails := 0

	for i := 1; i < limit; i++ {
		if Weighted(0.75) {
			heads++
		} else {
			tails++
		}
	}

	if heads < tails {
		t.Errorf("Weighted(0.75) heads should happen more than tails")
	}
}

func TestWeightAverage25(t *testing.T) {

	limit := 1000
	heads := 0
	tails := 0

	for i := 1; i < limit; i++ {
		if Weighted(0.25) {
			heads++
		} else {
			tails++
		}
	}

	if tails < heads {
		t.Errorf("Weighted(0.75) heads should happen more than tails")
	}
}
