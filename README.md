# Soundex

Source: [Programming in Go](http://www.qtrac.eu/gobook.html) chapter 3 exercise 2

## Reqs
- Create a web app
- `/` Present a form which the user can enter names to get soundex
- `/test` Show list of strings, soundex values, expectations
- `soundex` Robert -> R163. See [soundex wikipedia article](https://en.wikipedia.org/wiki/Soundex)

## Setup

	go get https://github.com/takahisah/soundex.git
	PATH=5000 soundex

## Deploy

	git push heroku master
