package main

import (
	"bytes"
	"fmt"
	"github.com/gosimple/slug"
	"github.com/russross/blackfriday"
	"os"
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

var (
	// BaseURL is the base URL.
	BaseURL = ""
	// Author is the author.
	Author = ""
	// Description is the description.
	Description = ""
	// Lang is the language.
	Lang = ""
)

// GetURL gets the URL based on the BaseURL path.
func GetURL(rel string) string {
	return StringConcat(StringConcat(BaseURL, "/"), rel)
}

// FprintHeader prints the standard HTML5 header.
func FprintHeader(file *os.File, title string) {
	fmt.Fprintf(file,
		"<!doctype html>\n"+
			"<html lang=\"%s\">\n"+
			"<head>\n"+
			"  <meta charset=\"utf-8\">\n"+
			"  <title>%s</title>\n"+
			"  <meta name=\"description\" content=\"%s\">\n"+
			"  <meta name=\"author\" content=\"%s\">\n"+
			"</head>\n"+
			"<body>\n\n",
		Lang, title, Description, Author)
}

// Markdown uses BlackFriday's markdown renderer to convert markdown to html.
func Markdown(mark string) string {
	return string(blackfriday.MarkdownCommon([]byte(mark)))
}

// StringConcat concatenates two strings.
func StringConcat(s1, s2 string) string {
	var buffer bytes.Buffer
	buffer.WriteString(s1)
	buffer.WriteString(s2)
	return buffer.String()
}
