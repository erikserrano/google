// Paquete encargado de obtener un token de Google en base a un código de autorización
// Documentación: https://developers.google.com/accounts/docs/OAuth2WebServer
package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	GOOGLE_API_OAUTH = "https://www.googleapis.com/oauth2/v4/token"
	GRANT_TYPE       = "authorization_code"
)

// Función encargada de solicitar token
func getToken(code, clientId, clientSecret, redirectUri string, reciver interface{}) error {
	resp, err := http.PostForm(
		GOOGLE_API_OAUTH,
		url.Values{
			"code":          {code},
			"client_id":     {clientId},
			"client_secret": {clientSecret},
			"redirect_uri":  {redirectUri},
			"grant_type":    {GRANT_TYPE},
		},
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, reciver)
	if err != nil {
		return err
	}
	oauth.Code = code

	return nil
}

// Función encargada de solicitar token
func GetToken(code, clientId, clientSecret, redirectUri string) (*OAuth2WebServer, error) {
	oauth := &OAuth2WebServer{}
	err := getToken(code, clientId, clientSecret, redirectUri, oauth)
	return oauth, err
}
