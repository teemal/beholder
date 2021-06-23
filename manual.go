package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func test_ping() {
	resp, err := http.Get("http://localhost:3000/ping")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}

func test_non_exist() {
	resp, err := http.Get("http://localhost:3000/poop")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}

func main() {
	// yeet := test_ping()
	// fmt.Println()
	test_ping()
	test_non_exist()
}
