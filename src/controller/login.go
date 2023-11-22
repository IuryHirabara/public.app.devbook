package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"app.devbook/src/config"
	"app.devbook/src/cookies"
	"app.devbook/src/models"
	"app.devbook/src/responses"
)

// Login faz a requisição para realizar a autenticação pela API
func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.APIURL)

	res, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(user),
	)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		responses.ErrorHandling(w, res)
		return
	}

	var authData models.AuthData
	if err := json.NewDecoder(res.Body).Decode(&authData); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.Error{Error: err.Error()})
		return
	}

	if err = cookies.Save(w, authData.ID, authData.Token); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.Error{Error: err.Error()})
		return
	}

	responses.JSON(w, http.StatusOK, map[string]string{"mensagem": "Usuário autenticado com sucesso"})
}
