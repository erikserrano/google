// Paquete encargado de obtener la información de un usuario en base a un token
// Documentación: https://developers.google.com/accounts/docs/OAuth2
package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	GOOGLE_API_OAUTH = "https://www.googleapis.com/oauth2/v2/userinfo"
)

// Función encargada de consultar los servicios de Google
// y obtener la información del usuario
func getUserInfo(token string, user *User) error {
	res, err := http.Get(GOOGLE_API_OAUTH + "?alt=json&access_token=" + token)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &user.JsonResponse)
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
