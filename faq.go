package main

import (
	"fmt"
	"strings"
)

// Faq is a FAQ page.
type Faq struct {
	// Title for FAQ.
	title string
	// FAQ description in Markdown.
	desc string
	// Topic separator.
	sep string
	// Topics.
	tops []*Topic
	// Reference table (references must be unique).
	refs map[string]*Topic
	// Base url.
	burl string
	// Tags
	tags map[string][]*Topic
}

// NewFaq creates a new FAQ page given a title.
func NewFaq(title, desc string) *Faq {
	return &Faq{title, desc, "<br><hr><br>", nil, make(map[string]*Topic), URLize(title), make(map[string][]*Topic)}
}

// AddTopic adds a new topic to this FAQ and returns it.
func (f *Faq) AddTopic(title, ref, short string, tags []string) *Topic {
	id := EncodeURL(title)
	index := len(f.tops)
	t := &Topic{title, short, NewPage(title, short, StringConcat(f.burl+"#", id), Pageify(index)), id, index}
	f.tops = append(f.tops, t)

	// Add topic tags.
	for _, s := range tags {
		l := f.tags[s]
		l = append(l, t)
	}

	f.refs[ref] = t

	return t
}

// This function does the actual ParseRef work.
func (f *Faq) parseRefsText(text string) string {
	insts := strings.Split(text, "</topic-ref>")

	var new string
	for _, token := range insts {
		tag := "<topic-ref=\""
		start := strings.Index(text, tag)
		data := token[start:len(token)]
		end := strings.Index(data, "\">")
		ref := data[0:end]
		data = ""
		name := token[start+len(tag)+len(ref)+2 : len(token)]
		link := GetURL(f.refs[ref].page.url)
		mkdwn := fmt.Sprintf("[%s](%s)", name, link)
		new = StringConcat(new, StringConcat(token[0:start], mkdwn))
	}

	return new
}

// ParseRefs finds all instances of <topic-ref="something">text</topic-ref> and turns them into a
// Markdown link pointing to the referenced page.
func (f *Faq) ParseRefs() {
	for _, t := range f.tops {
		t.short = f.parseRefsText(t.short)
		t.page.content = f.parseRefsText(t.page.content)
	}
}
