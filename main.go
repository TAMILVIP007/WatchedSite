package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const (
	TMDB_API_KEY = "d56e51fb77b081a9cb5192eaaa7823ad"
)

var (
	PORT = getPort()
)

func main() {
	fmt.Println("Listening on port " + PORT)
	http.ListenAndServe(":"+PORT, nil)
}

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index := template.Must(template.ParseFiles("index.html"))
		index.Execute(w, nil)
	})
	http.HandleFunc("/search", SearchTmdb)
	http.HandleFunc("/search/movie", getMovie)
	http.HandleFunc("/search/tv", getTvShow)
	http.HandleFunc("/movie/", func(w http.ResponseWriter, r *http.Request) {
		movie := template.Must(template.ParseFiles("movie.html"))
		movie.Execute(w, nil)
	})
}

func SearchTmdb(w http.ResponseWriter, r *http.Request) {
	_url := "https://api.themoviedb.org/3/search/multi"
	_query := r.FormValue("query")
	if _query == "Trending" {
		parseTrending(w, _query)
		return
	}
	params := map[string]string{
		"api_key":       TMDB_API_KEY,
		"query":         url.QueryEscape(_query),
		"language":      "en-US",
		"page":          "1",
		"include_adult": "false",
	}
	resp, err := http.Get(_url + "?" + encodeParams(params))
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprint(w, string(data))
}

func parseTrending(w http.ResponseWriter, query string) {
	_url := "https://api.themoviedb.org/3/trending/all/day"
	resp, err := http.Get(_url + "?api_key=" + TMDB_API_KEY)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprint(w, string(data))
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	_url := "https://api.themoviedb.org/3/movie/"
	_id := r.FormValue("id")
	params := map[string]string{
		"api_key":       TMDB_API_KEY,
		"language":      "en-US",
		"page":          "1",
		"include_adult": "false",
	}
	resp, err := http.Get(_url + _id + "?" + encodeParams(params))
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprint(w, string(data))
}

func getTvShow(w http.ResponseWriter, r *http.Request) {
	_url := "https://api.themoviedb.org/3/tv/"
	_id := r.FormValue("id")
	params := map[string]string{
		"api_key":       TMDB_API_KEY,
		"language":      "en-US",
		"page":          "1",
		"include_adult": "false",
	}
	resp, err := http.Get(_url + _id + "?" + encodeParams(params))
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprint(w, string(data))
}

func encodeParams(params map[string]string) string {
	var query string
	for k, v := range params {
		query += k + "=" + v + "&"
	}
	return query[:len(query)-1]
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "80"
	}
	return port
}
