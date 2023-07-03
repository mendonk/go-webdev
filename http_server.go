package main

import (
	"fmt"
	"net/http"
)

// process dynamic requests - HandleFunc registers handler function, first param is a path, second param is function to execute

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to my website!")
	})
	//create a file server and point to a URL path
	fs := http.FileServer(http.Dir("static/"))
	//strip static from url
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.ListenAndServe(":80", nil)
}
