package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadFile("web_html/main_window.html")
	// can file be opened?
	if err != nil {
		fmt.Println(err)
	}

	// convert bytes to string
	str := string(b)
	w.Write([]byte(str))
}

func web() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
