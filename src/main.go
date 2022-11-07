package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"io/ioutil"
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		b, err := ioutil.ReadFile("web_html/main_window.html")
		if err != nil {
			fmt.Println(err)
		}

		// convert bytes to string
		str := string(b)
		return str
		//return "Hello world!"
	})
	m.Get("/persons", func() string {
		return "Hello world!Persons"
	})
	m.Get("/persons/:id", func(params martini.Params) string {
		return "Hello world!Persons id:" + params["id"]
	})
	m.Post("/persons", func() string {
		return "Hello world!Post Persons"
	})
	m.Patch("/persons/:id", func(params martini.Params) string {
		return "Hello world!Persons patch"
	})
	m.Delete("/persons/:id", func(params martini.Params) string {
		fmt.Println("delete")
		return "Hello world!Persons delete"
	})

	m.Run()
}
