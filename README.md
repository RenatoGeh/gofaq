# GoFAQ

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

## Dependencies

GoFAQ uses:

- [BlackFriday's](https://github.com/russross/blackfriday) Markdown renderer.
- [GoSimple's](https://github.com/gosimple/slug) slugifier.

## Todo

- Example for clarity.
- "Beautify" pages.
