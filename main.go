package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

var routes = map[string]string{
	"/":               "index.html",
	"/about/":         "about.html",
	"/politics/":      "politics.html",
	"/business/":      "business.html",
	"/philanthropy/":  "philanthropy.html",
	"/impact/":        "impact.html",
	"/news/":          "news.html",
	"/events/":        "events.html",
	"/contacts/":      "contacts.html",
	"/video-gallery/": "video-gallery.html",
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	mux.Handle("/styles.css", http.FileServer(http.Dir(".")))
	mux.Handle("/inner.css", http.FileServer(http.Dir(".")))
	mux.Handle("/business.css", http.FileServer(http.Dir(".")))
	mux.Handle("/brand.css", http.FileServer(http.Dir(".")))
	mux.Handle("/images.css", http.FileServer(http.Dir(".")))
	mux.Handle("/script.js", http.FileServer(http.Dir(".")))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if file, ok := routes[r.URL.Path]; ok {
			servePage(w, file)
			return
		}
		http.NotFound(w, r)
	})
	log.Println("Serving at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func servePage(w http.ResponseWriter, file string) {
	page, err := os.ReadFile(file)
	if err != nil {
		http.NotFound(w, nil)
		return
	}
	page = []byte(strings.Replace(string(page), "</head>", `<link rel="icon" type="image/jpeg" href="/assets/images/logo.jpeg?v=1"><link rel="stylesheet" href="/brand.css"><link rel="stylesheet" href="/images.css"></head>`, 1))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write(page)
}
