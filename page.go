package main

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
