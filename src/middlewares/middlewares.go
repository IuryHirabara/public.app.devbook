package middlewares

import (
	"log"
	"net/http"

	"app.devbook/src/cookies"
)

// Logger escreve as requisições que são feitas para o app
func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}

// Auth verifica se os dados de autenticação estão salvos nos cookies do browser
func Auth(nexFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := cookies.Read(r)
		if err != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}

		nexFunc(w, r)
	}
}
