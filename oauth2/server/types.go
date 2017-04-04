// Package server makes request to Google API Oauth (https://developers.google.com/accounts/docs/OAuth2WebServer)
package server

// OAuth2WebServer represents the Google response
type OAuth2WebServer struct {
	Code         string
	JSONResponse map[string]interface{}
}

// QueryProperty queries properties in the response
func (oauth *OAuth2WebServer) QueryProperty(property string) interface{} {
	if len(oauth.JSONResponse) > 0 {
		if val, ok := oauth.JSONResponse[property]; ok == true {
			return val
		}
	}
	return nil
}
