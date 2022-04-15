package main

//package main => this is the part of the program we want to run when we start executeable
//or it has the main function ,and the main function is the first thing that runs
//init functions for packages will run first

import (
	"net/http"

	"github.com/PankajKumar9/lenslocked.com/controllers"
	"github.com/PankajKumar9/lenslocked.com/views"
	"github.com/gorilla/mux"
)

var (
	homeView    *views.View
	contactView *views.View
)

// func home(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Contact-Type", "text/html")

// 	must(homeView.Render(w, nil))
// }
// func contact(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Contact-Type", "text/html")
// 	must(contactView.Render(w, nil))
// }

// func NotFound(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	w.WriteHeader(http.StatusNotFound)
// 	fmt.Fprint(w, "<h1>no found bro</h1>")
// }

func main() {

	// 	homeView = views.NewView("bootstrap", "views/home.gohtml")
	// 	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	usersC := controllers.NewUsers()

	r := mux.NewRouter()
	staticC := controllers.NewStatic()
	r.Handle("/",staticC.Home).Methods("GET")
	r.Handle("/contact",staticC.Contact).Methods("GET")

	// r.HandleFunc("/", home).Methods("GET")
	// r.HandleFunc("/contact", contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	// r.HandleFunc("/faq", faq)
	// r.NotFoundHandler = http.HandlerFunc(NotFound)
	http.ListenAndServe(":3000", r)

}

func must(err error) {
	if err != nil {
		panic(err)
	}

}
