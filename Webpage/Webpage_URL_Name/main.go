//CSCI 130
//Leeladhar Reddy Munnangi
//Assignment 8(Create a webpage that serves at localhost:8080)

package main

import (
	"io"
	"net/http"
	"strings"
)

func handle(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	nm := strings.Split(req.URL.Path, "/")
	io.WriteString(res, nm[1])
}

func main() {

	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}