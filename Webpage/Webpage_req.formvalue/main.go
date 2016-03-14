//CSCI 130
//Leeladhar Reddy Munnangi
//Assignment 9 (Display the name in the URL when url is localhost:8080)

package main

import (
	"net/http"
	"io"
)

func main(){
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request){
		key := "q"
		val := req.FormValue(key)
		io.WriteString(res, "Query String - Name :"+val)
	})

	http.ListenAndServe(":8080",nil)
}