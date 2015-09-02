package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
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

func genTestCasesFromFile(reader io.Reader) (map[string]string, error) {
	var testCases = make(map[string]string)

	bufferedReader := bufio.NewReader(reader)
	eof := false
	var err error

	for !eof {
		var line string
		line, err = bufferedReader.ReadString('\n')
		if err == io.EOF {
			err = nil
			eof = true
		} else if err != nil {
			return nil, err
		}
		line = strings.Replace(line, "\n", "", -1)
		fields := strings.Split(line, " ")
		if len(fields) == 2 {
			testCases[fields[1]] = fields[0]
		}
	}

	return testCases, nil
}

func TestSoundexAlgorithm(t *testing.T) {
	var testdat io.Reader
	var err error
	var testCases map[string]string
	if testdat, err = os.Open("soundex-test-data.txt"); err != nil {
		log.Fatal(err)
	}

	if testCases, err = genTestCasesFromFile(testdat); err != nil {
		log.Fatal(err)
	}

	for name, expectedSoundex := range testCases {
		if soundex(name) != expectedSoundex {
			t.Errorf("soundex(%v) was '%v', expected '%v'",
				name, soundex(name), expectedSoundex)
		}
	}
}
