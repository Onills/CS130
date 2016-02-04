package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	IQ  int
	Smart bool
}

func main() {
	p1 := person{
		Name: "Carter",
		IQ:  99,
	}
	if p1.Smart > 130 {
		p1.Smart = true
	}
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}