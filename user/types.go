// Package user makes request to Google API OAuth2 (https://developers.google.com/accounts/docs/OAuth2) to get user information
package user

// User represents the Google response
type User struct {
	Token        string
	JSONResponse map[string]interface{}
}

// QueryProperty queries properties in the response
func (googleUser *User) QueryProperty(property string) interface{} {
	if len(googleUser.JSONResponse) > 0 {
		if val, ok := googleUser.JSONResponse[property]; ok == true {
			return val
		}
	}
	return nil
}
