package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"app.devbook/src/config"
	"app.devbook/src/cookies"
	"app.devbook/src/requests"
	"app.devbook/src/responses"
	"github.com/gorilla/mux"
)

// CreateUser chama a API para cadastrar um usuário
func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"nick":     r.FormValue("nick"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users", config.APIURL)

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

	responses.JSON(w, res.StatusCode, nil)
}

// UnfollowUser chama a API para parar de seguir um usuário
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/unfollow", config.APIURL, userId)
	res, err := requests.RequestWithAuth(r, http.MethodDelete, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		responses.ErrorHandling(w, res)
		return
	}

	responses.JSON(w, res.StatusCode, nil)
}

// FollowUser chama a API para seguir um usuário
func FollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/follow", config.APIURL, userId)
	res, err := requests.RequestWithAuth(r, http.MethodPost, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		responses.ErrorHandling(w, res)
		return
	}

	responses.JSON(w, res.StatusCode, nil)
}

// EditUser chama a API para editar um usuário
func EditUser(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":  r.FormValue("name"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d", config.APIURL, userId)
	res, err := requests.RequestWithAuth(r, http.MethodPut, url, bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		responses.ErrorHandling(w, res)
		return
	}

	responses.JSON(w, res.StatusCode, nil)
}

// UpdatePassword chama a API para atualizar a senha do usuário
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	password, err := json.Marshal(map[string]string{
		"new":     r.FormValue("new"),
		"current": r.FormValue("current"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d/updatePassword", config.APIURL, userId)
	res, err := requests.RequestWithAuth(r, http.MethodPost, url, bytes.NewBuffer(password))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		responses.ErrorHandling(w, res)
		return
	}

	responses.JSON(w, res.StatusCode, nil)
}

// DeleteUser chama a API para deletar um usuário
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userId, err := strconv.ParseUint(cookie["id"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d", config.APIURL, userId)
	res, err := requests.RequestWithAuth(r, http.MethodDelete, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		responses.ErrorHandling(w, res)
		return
	}

	responses.JSON(w, res.StatusCode, nil)
}
