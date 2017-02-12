package main

import (
	"bytes"
	"fmt"
	"github.com/gosimple/slug"
)

// EncodeURL takes a string and encodes it.
func EncodeURL(s string) string {
	return slug.MakeLang(s, "en")
}

// URLize turns a title into a filename dot html.
func URLize(t string) string {
	return StringConcat(slug.MakeLang(t, "en"), ".html")
}

// Pageify gives a name "page/pagen.html", where n is the index of the topic.
func Pageify(index int) string {
	return fmt.Sprintf("page/page%d.html", index)
}

// StringConcat concatenates two strings.
func StringConcat(s1, s2 string) string {
	var buffer bytes.Buffer
	buffer.WriteString(s1)
	buffer.WriteString(s2)
	return buffer.String()
}
