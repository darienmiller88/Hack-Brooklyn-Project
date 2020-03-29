package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server is running")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":5000", nil)
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "<h1>landing page :O</h1>")
}

func aboutHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "<h1>about page! :O</h1>")
}
