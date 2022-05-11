package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Req: %s %s \n", r.Host, r.URL.Path) 

		fmt.Fprint(w, "Hello world")
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}