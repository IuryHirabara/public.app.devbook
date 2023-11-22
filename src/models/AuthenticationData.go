package models

// AuthData representa um usuário autenticado
type AuthData struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
