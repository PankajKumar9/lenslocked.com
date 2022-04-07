package main

//package main => this is the part of the program we want to run when we start executeable
//or it has the main function ,and the main function is the first thing that runs
//init functions for packages will run first

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contact-Type", "text/html")
	fmt.Fprint(w, "<h1>this is the home function</h1>")

}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch,please send an email to <a href = \"mailto:support@lenslock\">support @lenslocked</a>")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>faq</h1>")
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>no found bro</h1>")
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = http.HandlerFunc(NotFound)
	http.ListenAndServe(":3000", r)

}
