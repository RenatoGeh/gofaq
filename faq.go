package main

import (
	"fmt"
	"os"
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
	// Tags.
	tags map[string][]*Topic
	// Categories.
	cats map[string][]*Topic
}

// NewFaq creates a new FAQ page given a title.
func NewFaq(title, desc string) *Faq {
	return &Faq{title, desc, "<br><hr><br>", nil, make(map[string]*Topic), URLize(title),
		make(map[string][]*Topic), make(map[string][]*Topic)}
}

// AddTopic adds a new topic to this FAQ and returns it.
func (f *Faq) AddTopic(title, ref, category, short string, tags []string) *Topic {
	id := EncodeURL(title)
	index := len(f.tops)
	t := &Topic{title, short, NewPage(title, short, StringConcat(f.burl+"#", id), Pageify(index)),
		id, index, false, f.burl}
	f.tops = append(f.tops, t)

	// Add topic tags.
	for _, s := range tags {
		f.tags[s] = append(f.tags[s], t)
	}

	f.refs[ref] = t
	f.cats[category] = append(f.cats[category], t)

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

// PrintSep prints the separator tag.
func (f *Faq) PrintSep(file *os.File) {
	fmt.Fprintf(file, "%s\n\n", f.sep)
}

// Print writes all the content from this FAQ to an HTML structure, with the main index file being
// named filename.
func (f *Faq) Print(filename string) {
	index, err := os.Create(filename)

	if err != nil {
		fmt.Println("Could not create file [%s].\n", filename)
		panic(err)
	}
	defer index.Close()

	// Print header (HTML5).
	FprintHeader(index, f.title)

	// Print header.
	fmt.Fprintf(index, "<h1>%s</h1>\n\n", f.title)

	// Print FAQ description.
	fmt.Fprintf(index, "%s\n\n", Markdown(f.desc))

	f.PrintSep(index)

	// Print alphabetical listing.

	fmt.Fprintf(index, "</body>\n</html>\n")
}
