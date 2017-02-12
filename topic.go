package main

// Topic is a FAQ topic.
type Topic struct {
	// Title for topic.
	title string
	// Short description in Markdown.
	short string
	// Page for topic. If nil, then generates page with short.
	page *Page
	// Id of this topic.
	id string
	// Slice index.
	index int
}

// SetPage sets a page for this topic and returns it.
func (t *Topic) SetPage(content string) *Page {
	t.page = NewPage(t.title, content, t.id, Pageify(t.index))
	return t.page
}
