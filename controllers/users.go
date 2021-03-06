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
	signIn(w,&user)
	//fmt.Fprintln(w, user)
	http.Redirect(w,r,"/cookietest",http.StatusFound)

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

	user, err := u.us.Authenticate(form.Email, form.Password)
	//do something with the login form
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	// fmt.Fprintln(w, "yaha aaya @xx")
	// fmt.Fprintln(w, user)
	signIn(w,user)
	http.Redirect(w,r,"/cookietest",http.StatusFound)
	//we would like to write after setting cookie
	fmt.Fprintln(w, user)
}

func signIn(w http.ResponseWriter, user *Users){
	cookie := http.Cookie{
		Name:  "email",
		Value: user.Email,
	}
	http.SetCookie(w, &cookie)
}

func (u *UsersSite) CookieTest(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("email")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintln(w, "Email is : ", cookie.Value)
	fmt.Fprintln(w, "\n ")
	fmt.Fprintln(w, "Email is : ", cookie)

}
