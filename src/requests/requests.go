package requests

import (
	"fmt"
	"io"
	"net/http"

	"app.devbook/src/cookies"
)

// RequestWithAuth é utilizada para colocar o token no cabeçalho da requisição
func RequestWithAuth(r *http.Request, method, url string, data io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(
		method,
		url,
		data,
	)
	if err != nil {
		return nil, err
	}

	cookie, _ := cookies.Read(r)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cookie["token"]))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
