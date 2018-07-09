package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func Test_extract_links(t *testing.T) {
	files := []struct {
		name, loc string
		want      []link
	}{
		{"should find a single anchor tag", "ex1.html",
			[]link{link{"/other-page", "A link to another page"}},
		},
		{"should strip inner html but keep text", "ex2.html",
			[]link{
				link{"https://www.twitter.com/joncalhoun", "Check me out on twitter"},
				link{"https://github.com/gophercises", "Gophercises is on Github!"},
			},
		},
		{"should find many anchor tags", "ex3.html",
			[]link{
				link{"#", "Login"},
				link{"/lost", "Lost? Need help?"},
				link{"https://twitter.com/marcusolsson", "@marcusolsson"},
			},
		},
		{"should not include comments or trailing whitespace", "ex4.html",
			[]link{
				link{"/dog-cat", "dog cat"},
				link{"/dog", "Something in a span Text not in a span Bold text!"},
			},
		},
	}

	for _, f := range files {
		location, err := os.Open(f.loc)
		if err != nil {
			log.Fatalln("Error opening file:", f.loc, "\nError:", err)
		}
		t.Run(f.name, func(t *testing.T) {
			got := extractLinks(html.NewTokenizer(location))
			assert.ElementsMatch(t, got, f.want, f.name)
		})
	}

}
