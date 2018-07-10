package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	f := flag.String("html", "ex1.html", "HTML File location")
	flag.Parse()

	file, err := os.Open(*f)
	if err != nil {
		log.Fatalln("Error opening file:", err)
	}

	links := extractLinks(html.NewTokenizer(file))
	fmt.Printf("%+v\n", links)
}

type link struct {
	href, text string
}

func extractLinks(z *html.Tokenizer) (l []link) {
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			if err := z.Err().Error(); err != "EOF" {
				log.Fatalln("HTML Error Token found:", err)
			}
			break
		}
		el, _ := z.TagName()
		if string(el) == "a" && tt.String() == "StartTag" {
			_, href, _ := z.TagAttr()
			text := extractText(z)
			l = append(l, extractLinks(z)...)
			l = append(l, link{string(href), strings.TrimSpace(text)})
		}
	}
	return l
}

func extractText(z *html.Tokenizer) string {
	var ret string
	for {
		z.Next()
		el, _ := z.TagName()
		if z.Token().Type == html.EndTagToken && string(el) == "a" {
			break
		}
		if z.Token().Type == html.TextToken {
			ret += " " + strings.TrimSpace(string(z.Raw()))
		}
	}
	return ret
}
