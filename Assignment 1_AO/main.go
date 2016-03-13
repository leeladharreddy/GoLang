//CSCI 130
//Leeladhar Reddy Munnangi
//Assignment 1 (Create a templete that uses Conditional Logic)

package main

import (
	"log"
	"os"
	"text/template"
)
type person struct{
	Name string
}

type conditions struct {
	person 
	IsYou bool

}

func main() {

	p1 := conditions{
		person: person{
			Name: "abcdefg" ,
		},
		IsYou:false,
	}

	if p1.Name == "Leeladhar" {
		p1.IsYou = true
	}

	tmp, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tmp.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}
