package controllers

import (
	"fmt"
	"net/http"

	"github.com/PankajKumar9/lenslocked.com/views"
)

type Users struct {
	NewView *views.View
}

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
	}
}

//get signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

//struct tags k baare me dekhne ka
type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

//post signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {

	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)

	// fmt.Fprintln(w, r.PostForm["email"])
	// // fmt.Fprintln(w, r.PostFormValue("email"))
	// fmt.Fprintln(w, r.PostForm["password"])
	// //fmt.Fprintln(w, r.PostFormValue("password"))
	// fmt.Fprintln(w, "Pretend theis is a fake message")

}
