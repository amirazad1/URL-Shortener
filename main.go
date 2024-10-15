package main

import (
	"fmt"
	"net/http"
)

func main() {
	shortener := &URLShortener{
		urls: make(map[string]string),
	}

	http.HandleFunc("/", shortener.Home)
	http.HandleFunc("/shorten", shortener.Shorten)
	http.HandleFunc("/short/", shortener.Redirect)

	fmt.Println("URL Shortener is running on :8080")
	http.ListenAndServe(":8080", nil)
}
