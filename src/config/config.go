package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIURL representa a URL para comunicação com a API
	APIURL = ""
	// Port onde a aplicação web está rodando
	Port = 0
	// HashKey é utilizada para autenticar o cookie
	HashKey []byte
	// BlockKey é utilizado para criptografar os dados do cookie
	BlockKey []byte
)

// Load inicializa as váriaveis de ambiente
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("Port"))
	if err != nil {
		log.Fatal(err)
	}

	APIURL = os.Getenv("APIURL")
	HashKey = []byte(os.Getenv("HashKey"))
	BlockKey = []byte(os.Getenv("BlockKey"))
}
