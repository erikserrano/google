// Package server enables make Oauth authentication into Google
// Documentation: https://developers.google.com/accounts/docs/OAuth2WebServer
package server

// OAuth2WebServer holds the code to make request to Google API
type OAuth2WebServer struct {
	Code         string
	JSONResponse map[string]interface{}
}

// QueryProperty allows query propertieses in the Google response
func (oauth *OAuth2WebServer) QueryProperty(property string) interface{} {
	if len(oauth.JSONResponse) > 0 {
		if val, ok := oauth.JSONResponse[property]; ok == true {
			return val
		}
	}
	return nil
}
