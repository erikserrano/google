// Package user requests for user information to Google API
// Documentation: https://developers.google.com/accounts/docs/OAuth2
package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	// GoogleAPIOauth is the URL for Google API
	GoogleAPIOauth = "https://www.googleapis.com/oauth2/v2/userinfo"
)

// Función encargada de consultar los servicios de Google
// y obtener la información del usuario
// getUserInfo request
func getUserInfo(token string, reciver interface{}) error {
	res, err := http.Get(GoogleAPIOauth + "?alt=json&access_token=" + token)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, reciver)
	if err != nil {
		return err
	}

	return nil
}

// Función encargada de consultar la información del usuario en base a un token
func GetUser(token string) (*User, error) {
	googleUser := &User{Token: token}
	err := getUserInfo(token, googleUser)
	return googleUser, err
}
