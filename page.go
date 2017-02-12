package main

import (
	"fmt"
	"os"
)

// Page is a dedicated page for a certain topic.
type Page struct {
	// Title should be the same for topic.
	title string
	// Content in Markdown.
	content string
	// Back URL (topic's id).
	burl string
	// This page's URL.
	url string
}

// NewPage creates a new page.
func NewPage(title, content, burl, url string) *Page {
	return &Page{title, content, burl, url}
}

// Print writes the contents of this page to the designated page file.
func (p *Page) Print() {
	out, err := os.Create(p.url)

	if err != nil {
		fmt.Printf("Could not create file [%s].\n", p.url)
		panic(err)
	}
	defer out.Close()

	// Print header.
	FprintHeader(out, p.title)

	// Print title.
	fmt.Fprintf(out, "<h1>%s</h1>\n\n", p.title)

	// Print contents.
	fmt.Fprintf(out, "%s\n\n", Markdown(p.content))

	// Print back link.
	fmt.Fprintf(out, "<a href=\"%s\">[back]</a>\n", GetURL(p.burl))
}
