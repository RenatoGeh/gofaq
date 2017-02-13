package main

import (
	"flag"
)

func main() {
	flag.StringVar(&Author, "author", "Dog", "The author's name.")
	flag.StringVar(&Description, "desc", "Hi! This is Dog, and this is my dog-site!",
		"A short description of the website.")
	flag.StringVar(&Lang, "lang", "en",
		"Please refer to http://www.w3schools.com/tags/ref_language_codes.asp for more information.")
	flag.StringVar(&BaseURL, "url", "", "The base URL to be prepended to all URLs.")
	flag.StringVar(&Footer, "footer", "", "Optional footer.")

	var input string
	flag.StringVar(&input, "in", "example.faq", "The faq-script to be parsed.")

	var output string
	flag.StringVar(&output, "out", "index.html", "The output filename.")

	flag.Parse()

	faq := Parse(input, output)
	faq.Print()
}
