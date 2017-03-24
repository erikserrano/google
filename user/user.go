// Package user enable get google user information
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

// geUserInfo makes http request to Google API
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

// GetUser requests to Google API for user information
func GetUser(token string) (*User, error) {
	googleUser := &User{Token: token}
	err := getUserInfo(token, googleUser)
	return googleUser, err
}
