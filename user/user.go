// Package user makes request to Google API OAuth2 (https://developers.google.com/accounts/docs/OAuth2) to get user information
package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// GoogleAPIOauth is the URL for Google API
	GoogleAPIOauth = "https://www.googleapis.com/oauth2/v2/userinfo"
)

func getUserInfo(token string, reciver interface{}) error {
	res, err := http.Get(GoogleAPIOauth + "?alt=json&access_token=" + token)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Google server response with: %s", res.Status)
	}

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

// GetUser makes GET request to Google API Oauth API and save the response
func GetUser(token string) (*User, error) {
	googleUser := &User{Token: token}
	err := getUserInfo(token, &googleUser.JSONResponse)
	return googleUser, err
}
