package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	f, err := os.Open("ex1.html")
	if err != nil {
		log.Fatalln("Error opening file:", err)
	}

	t := html.NewTokenizer(f)
	links := extractLinks(t)
	fmt.Println(links)
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
		_, href, _ := z.TagAttr()
		if string(el) == "a" && tt.String() == "StartTag" {
			z.Next()
			l = append(l, link{string(href), string(z.Text())})
		}
	}
	return l
}
