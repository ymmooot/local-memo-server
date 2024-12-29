package main

import (
	_ "embed"
	"net/http"
)

//go:embed index.html
var indexHTML string

//go:embed favicon.ico
var favicon []byte

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.ListenAndServe(":1234", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(indexHTML))
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/x-icon")
	w.Write(favicon)
}
