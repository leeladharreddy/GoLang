//CSCI 130
//Leeladhar Reddy Munnangi
//Assignment 11 (Create a webpage that serves a form and allows the user to upload a txt file)

package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	var value string
	http.ListenAndServe(":8080", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			file, _, err := req.FormFile("my-file")
			if err != nil {
				http.Error(res, err.Error(), 500)
				return
			}
			defer file.Close()

			src := io.LimitReader(file, 400)

			dst, err := os.Create(filepath.Join(".", "read.txt"))
			if err != nil {
				http.Error(res, err.Error(), 500)
				return
			}
			defer dst.Close()

			io.Copy(dst, src)

			contents, err := ioutil.ReadFile("read.txt")
			value = string(contents)

		}

		res.Header().Set("Content-Type", "text/html")
		io.WriteString(res, `
      <form method="POST" enctype="multipart/form-data">
        <input type="file" name="my-file">
        <input type="submit">
      </form>
      `+`<br/>`+value)
	}))
}