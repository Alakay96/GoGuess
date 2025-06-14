package main

import (
	"testing"
)

func TestRandomRange(t *testing.T) {
	min := 1
	max := 100
	number := randInt(min, max)
	if number > 100 {
		t.Errorf("Number greater than 100, got '%d'", number)
	} else if number < 1 {
		t.Errorf("Number less than 1, got '%d'", number)
	}
}
