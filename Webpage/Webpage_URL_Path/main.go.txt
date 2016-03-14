//CSCI 130
//Leeladhar Reddy Munnangi
//Assignment 7 (Create a webpage that displays the URL path using req.URL.path)
//Location: AW

package main

import (
    "fmt"
    "net/http"
)

func urlName(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "%v", req.URL.Path)
}

func main() {
    http.HandleFunc("/", urlName)
    http.ListenAndServe(":8080", nil)
}