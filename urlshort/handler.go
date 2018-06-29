package urlshort

import (
	"encoding/json"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

type pathMap map[string]string

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls pathMap, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path, ok := pathsToUrls[r.URL.Path]
		if ok {
			http.Redirect(w, r, path, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

func makePathMap(in []pathMap) pathMap {
	m := make(pathMap)
	for _, line := range in {
		m[line["path"]] = line["url"]
	}
	return m
}

type unmarshaler func([]byte, interface{}) error

func unmarshal(in []byte, fn unmarshaler) (out []pathMap) {
	if err := fn(in, &out); err != nil {
		panic(err)
	}
	return out
}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
// JSON is expected in this format:
//
// 	[
// 	  {
// 	    "path": "/cheese",
// 	    "url": "https://cheese.com/"
// 	  }, {
// 	    "path": "/bar",
// 	    "url": "https://bar.com/"
// 	  }
// 	]
//
func JSONHandler(jsn []byte, fallback http.Handler) (http.HandlerFunc, error) {
	return MapHandler(makePathMap(unmarshal(jsn, json.Unmarshal)), fallback), nil
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	return MapHandler(makePathMap(unmarshal(yml, yaml.Unmarshal)), fallback), nil
}
