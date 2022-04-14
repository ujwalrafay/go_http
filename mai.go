package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("got request from some one", req.Method)

	switch req.Method {
	case "GET":
		fmt.Fprintln(w, "got get method")
	case "POST":
		fmt.Fprintln(w, "got post method")

	case "PUTS":
		fmt.Fprintln(w, "got put method")
	}
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	fmt.Println(http.ListenAndServe(":8090", nil))
}
