package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	demoUrl, err := url.Parse("http://localhost:8080")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(demoUrl)
	http.ListenAndServe(":8081", proxy)
}