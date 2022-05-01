package controllers

import (
	"fmt"
	"net/http"

	"github.com/PankajKumar9/lenslocked.com/views"
)

const userPwPepper = ""

type UsersSite struct {
	NewView   *views.View
	LoginView *views.View
	us        UserService
}

func NewUsers() *UsersSite {
	return &UsersSite{
		NewView:   views.NewView("bootstrap", "users/new"),
		LoginView: views.NewView("bootstrap", "users/login"),
		us:        UserService{},
	}
}

//get signup
func (u *UsersSite) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

//struct tags k baare me dekhne ka
type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
	Name     string `schema:"name"`
}

//post signup
func (u *UsersSite) Create(w http.ResponseWriter, r *http.Request) {

	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user := Users{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}
	u.us = UserService{}
	_, err := u.us.Create(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, user)

	// fmt.Fprintln(w, r.PostForm["email"])
	// // fmt.Fprintln(w, r.PostFormValue("email"))
	// fmt.Fprintln(w, r.PostForm["password"])
	// //fmt.Fprintln(w, r.PostFormValue("password"))
	// fmt.Fprintln(w, "Pretend theis is a fake message")

}

type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func (u *UsersSite) Login(w http.ResponseWriter, r *http.Request) {
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Println(form.Password)
	fmt.Println("the email is:", form.Email)
	user, err := u.us.Authenticate(form.Email, form.Password)
	//do something with the login form
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, "yaha aaya @xx")
	fmt.Fprintln(w, user)
}
