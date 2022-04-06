package main

//package main => this is the part of the program we want to run when we start executeable
//or it has the main function ,and the main function is the first thing that runs
//init functions for packages will run first

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1> Welcome to my awesome site too </h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch,please send an email to <a href = \"mailto:support@lenslock\">support @lenslocked</a>")
	}

	//fmt.Fprint(w, "<h1> Welcome to my awesome site too </h1>")
	//fmt.Fprint(w, "To get in touch,please send an email to <a href = \"mailto:support@lenslock\">support @lenslocked</a>")
}

func main() {

	http.HandleFunc("/", handlerFunc)

	http.ListenAndServe(":3000", nil)

}
