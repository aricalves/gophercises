package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/aricalves/gophercises/urlshort"
	"github.com/boltdb/bolt"
)

func main() {
	yamlFile := flag.String("yaml", "./data/default.yaml", "Name of YAML file.")
	jsonFile := flag.String("json", "./data/default.json", "Name of JSON file.")
	dbLocation := flag.String("db", "./data/urlshort.db", "Name of BoltDB file.")

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

	yamlSlice, err := ioutil.ReadAll(yaml)
	if err != nil {
		panic(err)
	}

	// Build the YAMLHandler using the mapHandler as the fallback
	yamlHandler, err := urlshort.YAMLHandler(yamlSlice, mapHandler)
	if err != nil {
		panic(err)
	}

	json, err := os.Open(*jsonFile)
	if err != nil {
		panic(err)
	}
	defer json.Close()

	jsonSlice, err := ioutil.ReadAll(json)
	if err != nil {
		panic(err)
	}

	// Build the JSONHandler using the yamlHandler as the fallback
	jsonHandler, err := urlshort.JSONHandler(jsonSlice, yamlHandler)
	if err != nil {
		panic(err)
	}

	db, err := bolt.Open(*dbLocation, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Init and populate DB
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("urlshort"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		err = b.Put([]byte("/twitter"), []byte("https://twitter.com/aric_alves"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	paths, err := urlshort.ReadDB(db)
	if err != nil {
		panic(db)
	}
	dbHandler := urlshort.MapHandler(paths, jsonHandler)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", dbHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
