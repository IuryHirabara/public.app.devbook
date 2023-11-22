package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"app.devbook/src/config"
	"app.devbook/src/cookies"
	"app.devbook/src/models"
	"app.devbook/src/requests"
	"app.devbook/src/responses"
	"app.devbook/src/utils"
	"github.com/gorilla/mux"
)

// LoadLoginPage carrega a tela de login
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}

	utils.ExecuteTemplate(w, "login.html", nil)
}

// LoadRegisterPage carrega a tela de cadastro de usuário
func LoadRegisterPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "register.html", nil)
}

// LoadHomePage carrega a home com as publições dos usuários
func LoadHomePage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publications", config.APIURL)
	res, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		responses.ErrorHandling(w, res)
		return
	}

	var publications []models.Publication
	if err = json.NewDecoder(res.Body).Decode(&publications); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.Error{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecuteTemplate(w, "home.html", struct {
		Publications []models.Publication
		UserID       uint64
	}{
		Publications: publications,
		UserID:       userID,
	})
}

// LoadEditPublicationPage Carrega a página para editar uma publicação com as informações da publicação
func LoadEditPublicationPage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationId, err := strconv.ParseUint(params["publicationId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d", config.APIURL, publicationId)
	res, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		responses.ErrorHandling(w, res)
		return
	}

	var publication models.Publication
	if err = json.NewDecoder(res.Body).Decode(&publication); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.Error{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "updatePublication.html", publication)
}

// LoadUsersPage carrega a página com os usuários que atendem o filtro passado
func LoadUsersPage(w http.ResponseWriter, r *http.Request) {
	search := strings.ToLower(r.URL.Query().Get("user"))

	url := fmt.Sprintf("%s/users?search=%s", config.APIURL, search)

	res, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		responses.ErrorHandling(w, res)
		return
	}

	var users []models.User
	if err = json.NewDecoder(res.Body).Decode(&users); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.Error{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "users.html", users)
}

// LoadUserProfile carrega a página de perfil do usuário consultado
func LoadUserProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Error{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	loggedUserID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	if userId == loggedUserID {
		http.Redirect(w, r, "/profile", 302)
		return
	}

	user, err := models.SearchCompleteUser(userId, r)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "user.html", struct {
		User         models.User
		LoggedUserID uint64
	}{
		User:         user,
		LoggedUserID: loggedUserID,
	})
}

// LoadLoggedUserProfile carrega a página de perfil do usuário logado
func LoadLoggedUserProfile(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)

	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)
	user, err := models.SearchCompleteUser(userId, r)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Error{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "profile.html", user)
}

// LoadUserEditPage carrega a página de edição de perfil do usuário
func LoadUserEditPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)

	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)
	channel := make(chan models.User)
	go models.SearchUserData(channel, userId, r)
	user := <-channel

	if user.ID == 0 {
		responses.JSON(
			w,
			http.StatusInternalServerError,
			responses.Error{Error: "Erro na consulta do perfil do usuário"},
		)
		return
	}

	utils.ExecuteTemplate(w, "edit-user.html", user)
}

// LoadUpdatePasswordPage carrega a página de atualização de senha
func LoadUpdatePasswordPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "update-password.html", nil)
}
