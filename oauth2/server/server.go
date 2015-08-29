/*
* Get Google Token by Code
* Documentation: https://developers.google.com/accounts/docs/OAuth2WebServer
 */

package server

import (
    "encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	GOOGLE_API_OAUTH = "https://www.googleapis.com/oauth2/v3/token"
	GRANT_TYPE       = "authorization_code"
)

func getToken(code, clientId, clientSecret, redirectUri string, oauth *OAuth2WebServer) error {
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

	err = json.Unmarshal(body, &oauth.JsonResponse)
	if err != nil {
		return err
	}
	oauth.Code = code

	return nil
}

func GetToken(code, clientId, clientSecret, redirectUri string) (*OAuth2WebServer, error) {
	oauth := &OAuth2WebServer{}
	err := getToken(code, clientId, clientSecret, redirectUri, oauth)
	return oauth, err
}
