package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// replace consonants
var consonantMapping = map[rune]int{
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

const (
	pageTop = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Soundex</title>
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
		if names, message, ok := processRequest(req); ok {
			soundexMappings := genSoundexMappings(names)
			fmt.Fprint(res, formatSoundex(soundexMappings))
		} else if message != "" {
			fmt.Fprintf(res, anError, message)
		}
	}
	fmt.Fprint(res, pageBottom)
}

func genSoundexMappings(names []string) map[string]string {
	mapping := make(map[string]string)
	for _, name := range names {
		mapping[name] = soundex(name)
	}

	return mapping

}

func shortenName(name string) string {
	var prevRune rune
	var shortenedName string

	for _, c := range name {
		if c != prevRune {
			prevRune = c
			shortenedName += string(c)
		}
	}

	return shortenedName
}

func soundex(name string) string {
	shortenedName := shortenName(name)

	firstLetter := shortenedName[:1]
	remainingLetters := shortenedName[1:]

	// drop all vowels
	for _, vowel := range []string{"a", "e", "i", "o", "u", "y"} {
		remainingLetters = strings.Replace(remainingLetters, vowel, "*", -1)
	}

	for _, hw := range []string{"h", "w"} {
		remainingLetters = strings.Replace(remainingLetters, hw, "|", -1)
	}

	var digits string
	for _, consonant := range remainingLetters {
		if consonant == '*' || consonant == '|' {
			digits += string(consonant)
		} else {
			digits += strconv.Itoa(consonantMapping[consonant])
		}
	}

	// first letter check
	i, err := strconv.Atoi(digits[:1])
	if err == nil {
		if consonantMapping[[]rune(strings.ToLower(firstLetter))[0]] == i {
			digits = digits[1:]
		}
	}

	// collapse numbers
	var prevRune rune
	var collapsed string

	for _, num := range digits {
		if num == '*' {
			prevRune = num
		} else if num == '|' {
			// skip
		} else if num != prevRune {
			prevRune = num
			collapsed += string(num)
		}
	}

	// return all digits or fill rest with zeros
	if len(collapsed) >= 3 {
		return firstLetter + collapsed[:3]
	} else {
		return firstLetter + collapsed + strings.Repeat("0", 3-len(collapsed))
	}
}

func processRequest(req *http.Request) ([]string, string, bool) {
	var names []string

	if slice, found := req.Form["name"]; found && len(slice) > 0 {
		text := strings.Replace(slice[0], ",", " ", -1)
		var validName = regexp.MustCompile(`^[A-Z][a-z]*`)

		for _, field := range strings.Fields(text) {
			if validName.MatchString(field) {
				names = append(names, field)
			} else {
				return names, "'" + field + "'" + "is invalid name", false
			}
		}
	}

	if len(names) == 0 {
		return names, "", false // no data first time shown
	}

	return names, "", true
}

//TODO: Pretty formatting of map to html table
func formatSoundex(code map[string]string) string {
	return fmt.Sprintf("<p>%v</p>", code)
}
