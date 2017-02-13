package main

import (
	"fmt"
	"os"
)

// Category is a category.
type Category struct {
	// Topics.
	tops []*Topic
	// Full name.
	title string
	// URL-ized full name.
	url string
}

// NewCategory creates a new Category.
func NewCategory(title string) *Category {
	return &Category{nil, title, Categorify(title)}
}

// Print prints the category page.
func (c *Category) Print() {
	out, err := os.Create(c.url)

	if err != nil {
		fmt.Printf("Could not create file [%s].\n", c.url)
		panic(err)
	}
	defer out.Close()

	FprintHeader(out, c.title)

	fmt.Fprintf(out, "<h1 id=\"top\">Category: \"%s\"</h1>\n\n", c.title)

	fmt.Fprintf(out, "The following topics are under this category:\n\n")
	fmt.Fprintln(out, "<ol>")
	for _, t := range c.tops {
		fmt.Fprintf(out, "  <li><a href=\"%s\">%s</a></li>\n", "#"+t.id, t.title)
	}
	fmt.Fprintln(out, "</ol>\n<br>")

	for i, t := range c.tops {
		t.Print(out, i)
	}

	FprintFooter(out)
}
