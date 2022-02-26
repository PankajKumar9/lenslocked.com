package main

import (
	"fmt"
	"net/http"
)

/* this is the function called anytime
somebody comes to our webserver */
func handlerFunc(w http.ResponseWriter, r *http.Request) { /* w http.ResponseWriter, this is where we write our response */
	/*fprint me first argument is ,print kaha krna h
	  jaise yaha w pe krna h */
	fmt.Fprint(w, "<h1> Welcome to my awesome site </h1>")

	/*r *http.Request : has the information sort of about the user is making
	  we can use that for sort of checking the user's ip address and stuff like that */

}
