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
	cats map[string]*Category
}

// NewFaq creates a new FAQ page given a title.
func NewFaq(title, desc, filename string) *Faq {
	return &Faq{title, desc, "<br><hr><br>", nil, make(map[string]*Topic), filename,
		make(map[string][]*Topic), make(map[string]*Category)}
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
	cat, exists := f.cats[category]
	if !exists {
		cat = NewCategory(category)
		f.cats[category] = cat
	}
	cat.tops = append(cat.tops, t)

	return t
}

// This function does the actual ParseRef work.
func (f *Faq) parseRefsText(text string) string {
	exists := strings.Index(text, "<topic-ref")
	if exists < 0 {
		return text
	}

	insts := strings.Split(text, "</topic-ref>")

	var new string
	for _, token := range insts {
		if strings.TrimSpace(token) == "" {
			continue
		}
		tag := "<topic-ref=\""
		start := strings.Index(token, tag)
		if start < 0 {
			continue
		}
		data := token[start+len(tag):]
		end := strings.Index(data, "\">")
		ref := data[0:end]
		data = ""
		name := token[start+len(tag)+len(ref)+2:]
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
func (f *Faq) Print() {
	index, err := os.Create(f.burl)

	if err != nil {
		fmt.Println("Could not create file [%s].\n", f.burl)
		panic(err)
	}
	defer index.Close()

	// Print header (HTML5).
	FprintHeader(index, f.title)

	// Print header.
	fmt.Fprintf(index, "<h1 div=\"top\">%s</h1>\n\n", f.title)

	// Print FAQ description.
	fmt.Fprintf(index, "%s\n\n", Markdown(f.desc))

	f.PrintSep(index)

	// Print category list.
	fmt.Fprintf(index, "<h2>%s</h2>\n\n", "Categories:")
	fmt.Fprintln(index, "<ul>")
	// Create category directory.
	os.Mkdir("category", 0777)
	for k, v := range f.cats {
		fmt.Fprintf(index, "  <li><a href=\"%s\">%s</a></li>\n", GetURL(v.url), k)
		v.Print()
	}
	fmt.Fprintln(index, "</ul>\n")

	f.PrintSep(index)

	// Print the main list of topics.
	fmt.Fprintf(index, "<h2>%s</h2>\n\n", "Frequently Asked Questions:")
	fmt.Fprintln(index, "<ol>")
	for _, t := range f.tops {
		fmt.Fprintf(index, "  <li><a href=\"%s\">%s</a></li>\n", "#"+t.id, t.title)
	}
	fmt.Fprintln(index, "</ol>\n")

	fmt.Fprintln(index, "<br>")

	// Create page dir.
	os.Mkdir("page", 0777)
	for i, t := range f.tops {
		t.Print(index, i)
		t.PrintPage()
	}

	f.PrintSep(index)

	// Print list of tags.
	fmt.Fprintf(index, "<h2>%s</h2>\n\n", "List of Tags:")
	fmt.Fprintf(index, "<ul>")
	// Create tags directory.
	os.Mkdir("tags", 0777)
	f.PrintTags()
	for tag := range f.tags {
		fmt.Fprintf(index, "  <li><a href=\"%s\">%s</a></li>\n", GetURL("tags/"+URLize(tag)), tag)
	}
	fmt.Fprintln(index, "</lu>\n")

	FprintFooter(index)
}

// PrintTags prints a page for tags.
func (f *Faq) PrintTags() {
	for tag, tops := range f.tags {
		name := StringConcat("tags/", URLize(tag))
		out, err := os.Create(name)

		if err != nil {
			fmt.Println("Could not create file [%s].\n", name)
			panic(err)
		}
		defer out.Close()

		FprintHeader(out, tag)

		fmt.Fprintf(out, "<h1 id=\"top\">Tag: \"%s\"</h1>\n\n", tag)

		fmt.Fprintf(out, "The following topics are under this tag:\n\n")
		fmt.Fprintln(out, "<ol>")
		for _, t := range tops {
			fmt.Fprintf(out, "  <li><a href=\"%s\">%s</a></li>\n", "#"+t.id, t.title)
		}
		fmt.Fprintln(out, "</ol>\n<br>")

		for i, t := range tops {
			t.Print(out, i)
		}

		FprintFooter(out)
	}
}
