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
	t.page = NewPage(t.title, content, t.burl, Pageify(t.index))
	t.more = true
	return t.page
}

// Print writes the contents of this topic to a file.
func (t *Topic) Print(file *os.File, index int) {
	// Print title.
	fmt.Fprintf(file, "<h3 id=\"%s\">%d. %s</h3>\n\n", t.id, index+1, t.title)

	// Print short description.
	fmt.Fprintf(file, "%s\n\n", Markdown(t.short))

	// Print top, index and more links.
	fmt.Fprintf(file, "<a href=\"%s\">[top]</a>  ", "#top")
	fmt.Fprintf(file, "<a href=\"%s\">[index]</a>  ", GetURL(t.burl))
	if t.more {
		fmt.Fprintf(file, "<a href=\"%s\">[more]</a>\n\n", GetURL(t.page.url))
	}
	fmt.Fprintf(file, "<br>\n")
}

// PrintPage prints this topic's page.
func (t *Topic) PrintPage() {
	t.page.Print()
}
