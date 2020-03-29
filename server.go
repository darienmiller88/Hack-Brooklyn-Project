package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("Server is running")
	// fs := http.FileServer(http.Dir("css"))
	// http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)

	//I hope this works :/...

	http.ListenAndServe(":5000", nil)
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	myTemplate, _ := template.ParseFiles("templates/example.html")
	myTemplate.Execute(writer, request)
}

func aboutHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "<h1>about page! :O</h1>")
}
