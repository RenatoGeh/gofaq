# GoFAQ

[![Go Report Card](https://goreportcard.com/badge/github.com/renatogeh/gofaq)](https://goreportcard.com/report/github.com/renatogeh/gofaq)
[![Build Status](https://travis-ci.org/RenatoGeh/gofaq.svg?branch=master)](https://travis-ci.org/RenatoGeh/gofaq)
[![GoDoc](https://godoc.org/github.com/RenatoGeh/gofaq?status.svg)](https://godoc.org/github.com/RenatoGeh/gofaq)

Ever wondered how to make a proper FAQ? Well, now you can GoFAQ yourself
in an easy and satisfying way! Automatically generate all questions,
answers, their pages, tags and categories from a single `faq-script`,
letting you enjoy the FAQing like never before.

The `faq-script` allows you to write your questions and answers in
Markdown so you don't have to go down and dirty with all the HTML tags.

## Installation

With Go and Git:

```
go get -u github.com/RenatoGeh/gofaq
```

Or import it:

```
import "github.com/RenatoGeh/gofaq"
```

## Usage

```
Usage of ./gofaq:
  -author string
        The author's name. (default "Dog")
  -desc string
        A short description of the website. (default "Hi! This is Dog, and this is my dog-site!")
  -footer string
        Optional footer.
  -lang string
        Please refer to http://www.w3schools.com/tags/ref_language_codes.asp for more information. (default "en")
  -url string
        The base URL to be prepended to all URLs.
```

See `main.go` to see how to use the code.

## Example

See `example.faq` for an example on `faq-script` syntax.

To compile the example, simply run `make`. See Makefile's recipe `test`
for an example on how to run GoFAQ.

## Dependencies

GoFAQ uses:

- [BlackFriday's](https://github.com/russross/blackfriday) Markdown renderer.
- [GoSimple's](https://github.com/gosimple/slug) slugifier.

## Todo

- "Beautify" pages.
