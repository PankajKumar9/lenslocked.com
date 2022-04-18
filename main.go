package main

//package main => this is the part of the program we want to run when we start executeable
//or it has the main function ,and the main function is the first thing that runs
//init functions for packages will run first

import (
	"net/http"

	"github.com/PankajKumar9/lenslocked.com/controllers"
	"github.com/PankajKumar9/lenslocked.com/models"
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

	r.HandleFunc("/api/movie", controllers.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movies", controllers.GetMyAllMovies).Methods("GET")
	r.HandleFunc("/api/movie/{id}", controllers.MarkAsWatched).Methods("PUT")

	r.HandleFunc("/api/movie/{id}", controllers.DeleteAMovie).Methods("DELETE")
	r.HandleFunc("/api/deleteallmovie", controllers.DeleteAllMovies).Methods("DELETE")

	staticC := controllers.NewStatic()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")

	// r.HandleFunc("/", home).Methods("GET")
	// r.HandleFunc("/contact", contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	// r.HandleFunc("/faq", faq)
	// r.NotFoundHandler = http.HandlerFunc(NotFound)
	var u models.Users
	controllers.CreateOrder(u,1001,"Fake description #1")


	http.ListenAndServe(":3000", r)

}

func must(err error) {
	if err != nil {
		panic(err)
	}

}
