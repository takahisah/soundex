package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numToConsonant = [][]rune{
	[]rune{'b', 'f', 'p', 'v'},
	[]rune{'c', 'g', 'j', 'k', 'q', 's', 'x', 'z'},
	[]rune{'d', 't'},
	[]rune{'l'},
	[]rune{'m', 'n'},
	[]rune{'r'},
}

var consonantMapping = invertRuneArray(numToConsonant)

func invertRuneArray(matrix [][]rune) map[rune]int {
	var invertedArray = make(map[rune]int)
	for i, runeArr := range numToConsonant {
		for _, c := range runeArr {
			invertedArray[c] = i + 1
		}
	}
	return invertedArray
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
	vowelReplacement := func(char rune) rune {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'y' {
			return '*'
		} else if char == 'h' || char == 'w' {
			return '|'
		}
		return char
	}

	remainingLetters = strings.Map(vowelReplacement, remainingLetters)

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
