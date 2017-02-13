package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Parse takes a faq-script file and stores its information into a Faq.
func Parse(filename, index string) *Faq {
	dat, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Printf("Could not open file [%s].\n", filename)
		panic(err)
	}

	input := string(dat)
	start := strings.Index(input, "\n")
	title := strings.TrimSpace(input[0:start])

	end := strings.Index(input, "<topic")
	desc := input[start+1 : end]

	faq := NewFaq(title, desc, index)

	input = input[end:len(input)]
	topics := strings.Split(input, "</topic>")
	input = ""

	for _, top := range topics {
		if strings.TrimSpace(top) == "" {
			break
		}

		// Extract reference tag.
		ttag := "<topic ref=\""
		start = strings.Index(top, ttag) + len(ttag)
		data := top[start:len(top)]
		end = strings.Index(data, "\"")
		ref := data[0:end]

		// Extract category tag.
		ttag = "category=\""
		start = strings.Index(data, ttag)
		data = data[start+len(ttag) : len(data)]
		end = strings.Index(data, "\">")
		category := data[0:end]

		// Extract title (which is always the first line).
		start = strings.Index(data, "\n") + 1
		data = data[start:len(data)]
		end = strings.Index(data, "\n")
		title := strings.TrimSpace(data[0:end])

		// Extract short description.
		start = end + 1
		data = data[start:len(data)]
		end = strings.Index(data, "<page>")
		var short string
		if end < 0 {
			short = data[0:strings.Index(data, "<tags=")]
		} else {
			short = data[0:end]
		}

		// Extract page contents if it exists.
		var content string
		if end >= 0 {
			ctag := strings.Index(data, "</page>")
			content = data[end+6 : ctag]
			data = data[end+1 : len(data)]
		}

		// Extract tags.
		start = strings.Index(data, "<tags=")
		data = data[start+6 : len(data)]
		end = strings.Index(data, ">")
		utags := data[0:end]
		tags := strings.Split(utags, ";")

		t := faq.AddTopic(title, ref, category, short, tags)
		if content != "" {
			t.SetPage(content)
		}
	}

	faq.ParseRefs()

	return faq
}
