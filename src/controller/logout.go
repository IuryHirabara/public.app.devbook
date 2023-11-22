package controller

import (
	"net/http"

	"app.devbook/src/cookies"
)

// Logout remove os dados de autenticação salvos no cookie
func Logout(w http.ResponseWriter, r *http.Request) {
	cookies.Delete(w)
	http.Redirect(w, r, "/login", 302)
}
