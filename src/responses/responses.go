package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// Error representa a resposta de erro da API
type Error struct {
	Error string `json:"error"`
}

// JSON json retorna uma resposta em formato json para o cliente
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

// ErrorHandling trata as requisições que o status code for do rang 400~500
func ErrorHandling(w http.ResponseWriter, r *http.Response) {
	var err Error
	json.NewDecoder(r.Body).Decode(&err)
	JSON(w, r.StatusCode, err)
}
