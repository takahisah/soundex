# Soundex

Source: [Programming in Go](http://www.qtrac.eu/gobook.html) chapter 3 exercise 2

## Reqs
- Create a web app
- `/` Present a form which the user can enter names to get soundex
- `/test` Show list of strings, soundex values, expectations
- `soundex` Robert -> R163. See [soundex wikipedia article](https://en.wikipedia.org/wiki/Soundex)

## Soundex algorithm
1. Retain the first letter of the name and drop all other occurrences of a, e, i, o, u, y, h, w.
2. Replace consonants with digits as follows (after the first letter):
    - b, f, p, v → 1
    - c, g, j, k, q, s, x, z → 2
    - d, t → 3
    - l → 4
    - m, n → 5
	- r → 6
3. If two or more letters with the same number are adjacent in the original name (before step 1), only retain the first letter; also two letters with the same number separated by 'h' or 'w' are coded as a single number, whereas such letters separated by a vowel are coded twice. This rule also applies to the first letter.
4. Iterate the previous step until you have one letter and three numbers. If you have too few letters in your word that you can't assign three numbers, append with zeros until there are three numbers. If you have more than 3 letters, just retain the first 3 numbers.


## Setup

	go get https://github.com/takahisah/soundex.git
	export HOST=localhost
	export PATH=5000
	go build
	./soundex

## Deploy

	git push heroku master
