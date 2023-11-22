package router

import "github.com/gorilla/mux"

// Create retorna um router com todas as rotas configuradas
func Create() *mux.Router {
	return mux.NewRouter()
}
