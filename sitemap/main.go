package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aricalves/gophercises/sitemap/link"
	"github.com/dolab/colorize"
)

func main() {
	warn := colorize.New("yellow")
	e := colorize.New("red")
	url := flag.String("url", "", "URL to build a sitemap from.\nexample: "+warn.Paint("github.com"))
	flag.Parse()

	if *url == "" {
		fmt.Println(warn.Paint("Please enter a url to build a sitemap from with the -url flag.\nRun this program with the -h flag for help."))
		fmt.Println("Exiting with no errors...")
		os.Exit(0)
	}

	resp, err := http.Get("http://www." + *url)
	if err != nil {
		fmt.Println(e.Paint(fmt.Errorf("Could not GET from https://www.%s\n%v", *url, err)))
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println(e.Paint("Error not handled while reading response body\n", err))
	}

	links, err := link.Parse(bytes.NewReader(body))
	if err != nil {
		fmt.Println(e.Paint("Error not handled while parsing response body links\n", err))
	}

	fmt.Printf("%+v\n", links)
}
