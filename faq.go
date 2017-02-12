package main

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
