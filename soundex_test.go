package main

import (
	"testing"
)

func TestSoundexAlgorithm(t *testing.T) {
	name := "Robert"
	expectedSoundex := "R163"

	if soundex(name) != expectedSoundex {
		t.Errorf("soundex(%v) was '%v', expected '%v'",
			name, soundex(name), expectedSoundex)
	}
}
