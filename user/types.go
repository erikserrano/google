// Package user enable get google user information
// Documentation: https://developers.google.com/accounts/docs/OAuth2
package user

// User holds the token & the Google response
type User struct {
	Token        string
	JSONResponse map[string]interface{}
}

// QueryProperty allows query propertieses in the Google response
func (googleUser *User) QueryProperty(property string) interface{} {
	if len(googleUser.JSONResponse) > 0 {
		if val, ok := googleUser.JSONResponse[property]; ok == true {
			return val
		}
	}
	return nil
}
