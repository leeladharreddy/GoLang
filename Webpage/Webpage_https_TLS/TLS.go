//CSCI 130
//Leeladhar Reddy Munnangi
//Assignment Webpage(Create webpage which serves at localhost server with https and tls )

package main

import (
    "io"
    "net/http"
    "log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "hello, world!\n")
}

func main() {
    http.HandleFunc("/hello", HelloServer)
    err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
