package urlshort

import (
	"encoding/json"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path, ok := pathsToUrls[r.URL.Path]
		if ok {
			http.Redirect(w, r, path, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
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
func JSONHandler(json []byte, fallback http.Handler) (http.HandlerFunc, error) {
	return MapHandler(jtom(parseJSON(json)), fallback), nil
}

type path struct {
	Path string `json:"path"`
	URL  string `json:"url"`
}

func parseJSON(s []byte) []path {
	var paths []path
	if err := json.Unmarshal(s, &paths); err != nil {
		panic(err)
	}
	return paths
}

func jtom(in []path) map[string]string {
	des := make(map[string]string)
	for _, line := range in {
		des[line.Path] = line.URL
	}
	return des
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
	return MapHandler(ytom(parseYAML(yml)), fallback), nil
}

func parseYAML(in []byte) (out []map[string]string) {
	err := yaml.Unmarshal(in, &out)
	if err != nil {
		panic(err)
	}
	return out
}

func ytom(y []map[string]string) map[string]string {
	paths := make(map[string]string)
	for _, line := range y {
		paths[line["path"]] = line["url"]
	}
	return paths
}
