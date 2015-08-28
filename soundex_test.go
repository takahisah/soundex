package main

import (
	"testing"
)

var testCases = map[string]string{
	"Robert":    "R163",
	"Ashcroft":  "A261",
	"Tymczak":   "T522",
	"Pfister":   "P236",
	"Burroughs": "B620",
	"Ellery":    "E460",
	"Example":   "E251",
	"Gauss":     "G200",
}

func TestSoundexAlgorithm(t *testing.T) {
	for name, expectedSoundex := range testCases {
		if soundex(name) != expectedSoundex {
			t.Errorf("soundex(%v) was '%v', expected '%v'",
				name, soundex(name), expectedSoundex)
		}
	}
}
