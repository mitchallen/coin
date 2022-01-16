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

	m := map[bool]int{
		true:  0,
		false: 0,
	}

	for i := 0; i < limit; i++ {
		m[Weighted(0.5)]++
	}

	if m[true] < threshold || m[false] < threshold {
		t.Errorf("Weighted(0.5) true false not very balanced")
	}
}

func TestWeightAverage75(t *testing.T) {

	limit := 1000

	m := map[bool]int{
		true:  0,
		false: 0,
	}

	for i := 0; i < limit; i++ {
		m[Weighted(0.75)]++
	}

	if m[true] < m[false] {
		t.Errorf("Weighted(0.75) true should happen more than false")
	}
}

func TestWeightAverage25(t *testing.T) {

	limit := 1000

	m := map[bool]int{
		true:  0,
		false: 0,
	}

	for i := 0; i < limit; i++ {
		m[Weighted(0.25)]++
	}

	if m[false] < m[true] {
		t.Errorf("Weighted(0.25) false should happen more than true")
	}
}
