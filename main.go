package main

//package main => this is the part of the program we want to run when we start executeable
//or it has the main function ,and the main function is the first thing that runs
//init functions for packages will run first

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Welcome to my awesome site to1243o </h1>")

}

func main() {

	http.HandleFunc("/", handlerFunc)

	http.ListenAndServe(":3000", nil)

}
