package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"app.devbook/src/config"
	"app.devbook/src/requests"
)

// User representa um usuário da aplicação
type User struct {
	ID           uint64        `json:"id"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Nick         string        `json:"nick"`
	CreatedAt    time.Time     `json:"createdAt"`
	Followers    []User        `json:"followers"`
	Following    []User        `json:"following"`
	Publications []Publication `json:"publications"`
}

// SearchCompleteUser faz as requisições necessárias para montar o
// usuário com suas publicações, seguidores e quem o usuário está seguindo
func SearchCompleteUser(userId uint64, r *http.Request) (User, error) {
	userChannel := make(chan User)
	followersChannel := make(chan []User)
	followingChannel := make(chan []User)
	publicationsChannel := make(chan []Publication)

	go SearchUserData(userChannel, userId, r)
	go SearchFollowers(followersChannel, userId, r)
	go SearchFollowing(followingChannel, userId, r)
	go SearchPublications(publicationsChannel, userId, r)

	var (
		user         User
		followers    []User
		following    []User
		publications []Publication
	)

	for i := 0; i < 4; i++ {
		select {
		case loadedUser := <-userChannel:
			if loadedUser.ID == 0 {
				return User{}, errors.New("Erro na requisição de busca de usuário")
			}

			user = loadedUser

		case loadedFollowers := <-followersChannel:
			if loadedFollowers == nil {
				return User{}, errors.New("Erro na requisição de buscar de seguidores")
			}

			followers = loadedFollowers

		case loadedFollowing := <-followingChannel:
			if loadedFollowing == nil {
				return User{}, errors.New("Erro na requisição de buscar usuários que o usuário está seguindo")
			}

			following = loadedFollowing

		case loadedPublications := <-publicationsChannel:
			if loadedPublications == nil {
				return User{}, errors.New("Erro na requisição de buscar publicações")
			}

			publications = loadedPublications
		}
	}

	user.Followers = followers
	user.Following = following
	user.Publications = publications

	return user, nil
}

// SearchUserData faz uma requisição para a API buscar os dados base do usuário
func SearchUserData(channel chan<- User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.APIURL, userId)

	res, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- User{}
		return
	}
	defer res.Body.Close()

	var user User
	if err = json.NewDecoder(res.Body).Decode(&user); err != nil {
		channel <- User{}
		return
	}

	channel <- user
}

// SearchFollowers faz uma requisição para a API buscar todos os seguindores do usuário
func SearchFollowers(channel chan<- []User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.APIURL, userId)

	res, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer res.Body.Close()

	var followers []User
	if err = json.NewDecoder(res.Body).Decode(&followers); err != nil {
		channel <- nil
		return
	}

	if followers == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- followers
}

// SearchFollowing faz uma requisição para a API para buscar todos os usuários que o usuários está seguindo
func SearchFollowing(channel chan<- []User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.APIURL, userId)

	res, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer res.Body.Close()

	var following []User
	if err = json.NewDecoder(res.Body).Decode(&following); err != nil {
		channel <- nil
		return
	}

	if following == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- following
}

// SearchPublications faz a requisição para a API buscar todas as publicações de um usuário
func SearchPublications(channel chan<- []Publication, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/publications", config.APIURL, userId)

	res, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer res.Body.Close()

	var publications []Publication
	if err = json.NewDecoder(res.Body).Decode(&publications); err != nil {
		channel <- nil
		return
	}

	if publications == nil {
		channel <- make([]Publication, 0)
		return
	}

	channel <- publications
}
