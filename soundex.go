package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	pageTop = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Statistics</title>
<body><h3>Soundex</h3>
<p>Converts name to soundex</p>`
	form = `<form action="/" method="POST">
<label for="name">Name:</label><br />
<input type="text" name="name" size="30"><br />
<input type="submit" value="Calculate">
</form>`
	pageBottom = `</body></html>`
	anError    = `<p class="error">%s</p>`
)

func main() {
	http.HandleFunc("/", home)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func home(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	fmt.Fprint(res, pageTop, form)
	if err != nil {
		fmt.Fprintf(res, anError, err)
	} else {
		if name, message, ok := processRequest(req); ok {
			soundexCode := soundex(name)
			fmt.Fprint(res, formatSoundex(soundexCode))
		} else if message != "" {
			fmt.Fprintf(res, anError, message)
		}
	}
	fmt.Fprint(res, pageBottom)
}

func soundex(name string) string {
	firstLetter := name[:1]
	remainingLetters := name[1:]

	//TODO: rule 3

	// drop all vowels
	for _, vowel := range []string{"a", "e", "i", "o", "u", "y", "h", "w"} {
		remainingLetters = strings.Replace(remainingLetters, vowel, "", -1)
	}

	// replace consonants
	consonantMapping := map[rune]int{
		'b': 1,
		'f': 1,
		'p': 1,
		'v': 1,
		'c': 2,
		'g': 2,
		'j': 2,
		'k': 2,
		'q': 2,
		's': 2,
		'x': 2,
		'z': 2,
		'd': 3,
		't': 3,
		'l': 4,
		'm': 5,
		'n': 5,
		'r': 6,
	}

	var digits int
	for _, consonant := range remainingLetters {
		digits = digits*10 + consonantMapping[consonant]
	}

	// collapse numbers
	numStrings := strconv.Itoa(digits)
	var prevRune rune
	var collapsed string

	for _, num := range numStrings {
		if num != prevRune {
			prevRune = num
			collapsed += string(num)
		}
	}

	if len(collapsed) >= 3 {
		return firstLetter + collapsed[:3]
	} else {
		return firstLetter + collapsed + strings.Repeat("0", 3-len(collapsed))
	}
}

func processRequest(req *http.Request) (string, string, bool) {
	return "name", "success", true
}

func formatSoundex(code string) string {
	return code
}
