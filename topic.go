package main

import (
	"fmt"
	"os"
)

// Topic is a FAQ topic.
type Topic struct {
	// Title for topic.
	title string
	// Short description in Markdown.
	short string
	// Page for topic.
	page *Page
	// Id of this topic.
	id string
	// Slice index.
	index int
	// Whether there is more to show.
	more bool
	// Back url.
	burl string
}

// SetPage sets a page for this topic and returns it.
func (t *Topic) SetPage(content string) *Page {
	t.page = NewPage(t.title, content, t.id, Pageify(t.index))
	t.more = true
	return t.page
}

// Print writes the contents of this topic to a file.
func (t *Topic) Print(file *os.File) {
	// Print title.
	fmt.Fprintf(file, "<h2>%s</h2>\n\n", t.title)

	// Print short description.
	fmt.Fprintf(file, "%s\n\n", Markdown(t.short))

	// Print top and more links.
	fmt.Fprintf(file, "<a href=\"%s\">[top]</a>  ", GetURL(t.burl))
	if t.more {
		fmt.Fprintf(file, "<a href=\"%s\">[more]</a>\n\n", GetURL(t.page.url))
	}
}

// PrintPage prints this topic's page.
func (t *Topic) PrintPage() {
	t.page.Print()
}
