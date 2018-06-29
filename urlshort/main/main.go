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

	yamlFile := flag.String("yaml", "default.yaml", "Name of YAML file.")
	jsonFile := flag.String("json", "default.json", "Name of JSON file.")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/aric":  "https://github.com/aricalves",
		"/dernz": "https://twitch.tv/jdernz",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yaml, err := os.Open(*yamlFile)
	if err != nil {
		panic(err)
	}
	defer yaml.Close()

	yamlLines, err := ioutil.ReadAll(yaml)
	if err != nil {
		panic(err)
	}

	// Build the YAMLHandler using the mapHandler as the fallback
	yamlHandler, err := urlshort.YAMLHandler(yamlLines, mapHandler)
	if err != nil {
		panic(err)
	}

	json, err := os.Open(*jsonFile)
	if err != nil {
		panic(err)
	}
	defer json.Close()

	jsonLines, err := ioutil.ReadAll(json)
	if err != nil {
		panic(err)
	}

	jsonHandler, err := urlshort.JSONHandler(jsonLines, yamlHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
