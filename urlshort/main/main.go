package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aricalves/gophercises/urlshort"
)

func main() {

	yamlFile := flag.String("yaml", "default.yaml", "Filename of paths and corresponding urls.")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/aric":  "https://github.com/aricalves",
		"/dernz": "https://twitch.tv/jdernz",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	f, err := os.Open(*yamlFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	// Build the YAMLHandler using the mapHandler as the fallback
	yamlHandler, err := urlshort.YAMLHandler(lines, mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
