// Package server makes request to Google API Oauth (https://developers.google.com/accounts/docs/OAuth2WebServer)
package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// GoogleAPIOauth holds URL for make request
	GoogleAPIOauth = "https://www.googleapis.com/oauth2/v4/token"
	// GrantType holds the type of request value sends
	GrantType = "authorization_code"
)

func getToken(code, clientID, clientSecret, redirectURI string, reciver interface{}) error {
	resp, err := http.PostForm(
		GoogleAPIOauth,
		url.Values{
			"code":          {code},
			"client_id":     {clientID},
			"client_secret": {clientSecret},
			"redirect_uri":  {redirectURI},
			"grant_type":    {GrantType},
		},
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Google server response with: %s; client id: %s; code: %s", resp.Status, clientID, code)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, reciver)
	if err != nil {
		return err
	}

	return nil
}

// GetToken makes POST request to Google Oauth API and save the response
func GetToken(code, clientID, clientSecret, redirectURI string) (*OAuth2WebServer, error) {
	oauth := &OAuth2WebServer{Code: code}
	err := getToken(code, clientID, clientSecret, redirectURI, &oauth.JSONResponse)
	return oauth, err
}
