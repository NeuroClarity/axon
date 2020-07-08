package middleware

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

func Log(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\nTrivial log of request form: %+v\n", r.Form)
		next.ServeHTTP(w, r)
	}
}

func Authenticate(next http.HandlerFunc, sessionStore *sessions.FilesystemStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := sessionStore.Get(r, "auth-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, ok := session.Values["profile"]; !ok {
			http.Redirect(w, r, "/api/permissions", http.StatusSeeOther)
		} else {
			next.ServeHTTP(w, r)
		}
	}
}
