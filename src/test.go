package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func test1() {
	println("test1")
	url := "http://localhost:8080/persons"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		println(err.Error())
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		println(err.Error())
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func test2() {
	println("test2")
	url := "http://localhost:8080/persons/1"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		println(err.Error())
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		println(err.Error())
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func test3() {
	println("test3")
	url := "http://localhost:8080/persons/"
	data := strings.NewReader(`{"id":3, "name": "petr" , "age": 42, "address":"47", "work":"46"}`)
	req, err := http.NewRequest("POST", url, data)
	req.Header.Add("Content-type", "application/json")
	if err != nil {
		println("004")
		println(err.Error())
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		println(005)
		println(err.Error())
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func test4() {
	println("test4")
	url := "http://localhost:8080/persons/3"
	data := strings.NewReader(`{"id" : 3 , "name": "oleg" , "age": 42, "address":"47", "work":"46"}`)
	req, err := http.NewRequest("PATCH", url, data)
	req.Header.Add("Content-type", "application/json")
	if err != nil {
		println("004")
		println(err.Error())
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		println(005)
		println(err.Error())
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func test5() {
	println("test5")
	url := "http://localhost:8080/persons/3"
	req, err := http.NewRequest("DELETE", url, nil)
	req.Header.Add("Content-type", "application/json")
	if err != nil {
		println("004")
		println(err.Error())
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		println(005)
		println(err.Error())
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func test6() {
	println("test6")
	println(HOST_URL + "/persons/")
	req, err := http.NewRequest("Get", HOST_URL+"/persons/", nil)
	println(":0")
	if err != nil {
		println(err.Error())
		return
	}
	println(":1")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		println(":1_1")
		println(err.Error())
		return
	}
	println(":2")
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		println(":1_1")
		println(err.Error())
		return
	}
	fmt.Println(string(body))
	println(":3")
}
