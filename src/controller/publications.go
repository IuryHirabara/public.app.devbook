package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"app.devbook/src/config"
	"app.devbook/src/requests"
	"app.devbook/src/responses"
	"github.com/gorilla/mux"
)

// CreatePublication faz a requisição para a API para criar uma nova publicação
func CreatePublication(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	publication, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications", config.APIURL)

	res, err := requests.RequestWithAuth(r, http.MethodPost, url, bytes.NewBuffer(publication))
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

// LikePublication faz a requisição para API para curtir uma publicação
func LikePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	publicationId, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d/like", config.APIURL, publicationId)
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

// DislikePublication faz a requisição para API para descurtir uma publicação
func DislikePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	publicationId, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d/dislike", config.APIURL, publicationId)
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

// EditPublication chama a API para alterar uma publicação
func EditPublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	publicationId, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	r.ParseForm()

	publication, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d", config.APIURL, publicationId)
	res, err := requests.RequestWithAuth(r, http.MethodPut, url, bytes.NewBuffer(publication))
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

// DeletePublication chama a API para deletar uma publicação
func DeletePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	publicationId, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d", config.APIURL, publicationId)

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
