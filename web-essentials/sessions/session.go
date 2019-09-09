package main

// do not forget to run: go get -u github.com/gorilla/sessions
import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// check if is authenticated
	if auth, ok :=session.Values["authenticated"].(bool); !ok || !auth{
		http.Error(w, "Forbidden!", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(w, "the cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Authentication goes here
	// ...
	
	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	//revoke user auth
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func main() {
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":81", nil)
}

// $ go run sessions.go

// $ curl -s http://localhost:8080/secret
// Forbidden

// $ curl -s -I http://localhost:8080/login
// Set-Cookie: cookie-name=MTQ4NzE5Mz...

// $ curl -s --cookie "cookie-name=MTQ4NzE5Mz..." http://localhost:8080/secret
// The cake is a lie!

// I did this and it worked, but I visited logout with the cookie 
// and it didint revoke the session...
// (used the same cookie afterwards and it still worked)
