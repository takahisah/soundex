package main

import (
	"fmt"
	"net/http"
	"os"
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
	return "write soundex func"
}

func processRequest(req *http.Request) (string, string, bool) {
	return "name", "success", true
}

func formatSoundex(code string) string {
	return code
}
