package main

import (
	"fmt"
	"log"
	"net/http"

	"app.devbook/src/config"
	"app.devbook/src/cookies"
	"app.devbook/src/router"
	"app.devbook/src/router/routes"
	"app.devbook/src/utils"
)

func main() {
	config.Load()
	cookies.Config()
	utils.LoadTemplates()

	r := router.Create()

	routes.Configure(r)

	fmt.Printf("App rodando na porta %d", config.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
