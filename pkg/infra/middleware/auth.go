package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

// Log is a trival Middleware that demonstrates simple patterning.
func Log(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\nTrivial log of request form: %+v\n", r.Form)
		next.ServeHTTP(w, r)
	}
}

// Authenticate follows the Middleware pattern for authenticating users with a valid session struct.
func Authenticate(next http.HandlerFunc, sessionStore *sessions.FilesystemStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := sessionStore.Get(r, "auth-session")
		if err != nil {
			log.Print(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, ok := session.Values["profile"]; !ok {
			log.Printf("Unauth access to endpoint %s with session profile: %+v\n", r.URL, session.Values["profile"])
			http.Redirect(w, r, "/api/permissions", http.StatusSeeOther)
		} else {
			next.ServeHTTP(w, r)
		}
	}
}
