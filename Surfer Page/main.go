
package main

import (
	"html/template"
	"log"
	"net/http"
)

func surfer(res http.ResponseWriter, req *http.Request) {
	tmp, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	tmp.Execute(res, nil)
}

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/pic/", http.StripPrefix("/pic/", http.FileServer(http.Dir("pic"))))

	http.HandleFunc("/", surfer)
	http.ListenAndServe(":8080", nil)
}