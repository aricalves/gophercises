package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	s := `
	<p>Links:</p>
	<ul>
		<li>
		<a href="foo">Foo</a>
		<li>
		<a href="/bar/baz">BarBaz</a>
	</ul>`
	t := html.NewTokenizer(strings.NewReader(s))
	printAnchorTags(t)
}

type link struct {
	href, text string
}

func printAnchorTags(z *html.Tokenizer) {
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			fmt.Println("\nEOF")
			return
		}
		a, _ := z.TagName()
		fmt.Printf("name:%v | text:%v\n", string(a), string(z.Text()))
	}
}
