package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Parse takes a faq-script file and stores its information into a Faq.
func Parse(filename string) *Faq {
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

	faq := NewFaq(title, desc)

	input = input[end:len(input)]
	topics := strings.Split(input, "</topic>")
	input = ""

	for _, top := range topics {
		var ref string
		fmt.Sscanf(top, "<topic ref=\"%s\">", &ref)

		start = strings.Index(top, "\n") + 1
		data := top[start:len(top)]
		end = strings.Index(data, "\n")
		title := data[0:end]

		start = end + 1
		data = data[start:len(data)]
		end = strings.Index(data, "<page>")
		short := data[0:end]
		var content string
		if end >= 0 {
			ctag := strings.Index(data, "</page>")
			content = data[end+6 : ctag]
			data = data[end+1 : len(data)]
		}
		var utags string
		fmt.Sscanf(data, "<tags=%s>", &utags)
		tags := strings.Split(utags, "; ")

		t := faq.AddTopic(title, ref, short, tags)
		if content != "" {
			t.SetPage(content)
		}
	}

	return faq
}
