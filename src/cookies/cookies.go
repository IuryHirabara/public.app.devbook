package cookies

import (
	"net/http"
	"time"

	"app.devbook/src/config"
	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Config configura o ambiente para com a configuração de SecureCookie
func Config() {
	s = securecookie.New([]byte(config.HashKey), []byte(config.BlockKey))
}

// Save registra as informações de autenticação
func Save(w http.ResponseWriter, id, token string) error {
	data := map[string]string{
		"id":    id,
		"token": token,
	}

	encodedData, err := s.Encode("data", data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    encodedData,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

// Read retorna os valores armazenados nos cookies
func Read(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie("data")
	if err != nil {
		return nil, err
	}

	values := make(map[string]string)
	if err = s.Decode("data", cookie.Value, &values); err != nil {
		return nil, err
	}

	return values, nil
}

// Delete deleta os dados de autenticação dos cookies
func Delete(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})
}
