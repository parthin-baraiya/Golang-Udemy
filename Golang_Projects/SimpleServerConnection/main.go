package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(res, "Path is invaild", http.StatusNotFound)
		return
	}

	if req.Method != "GET" {
		http.Error(res, "Method is not allowed", http.StatusNotFound)
		return
	}

	fmt.Fprintln(res, "Hello!")
}

func formHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/form" {
		http.Error(res, "Path is invaild", http.StatusNotFound)
		return
	}

	if req.Method != "POST" {
		http.Error(res, "Method is not allowed", http.StatusNotFound)
		return
	}

	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(res, "form parseForm() err: %v\n", err)
		return
	}

	fmt.Fprintln(res, "Post request is successful")

	firstName := req.FormValue("firstName")
	lastName := req.FormValue("lastName")

	fmt.Fprintf(res, "FirstName:%v\n", firstName)
	fmt.Fprintf(res, "LastName:%v\n", lastName)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Server is running in port 8000")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
