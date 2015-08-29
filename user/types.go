package user

type User struct {
    Token        string
	JsonResponse map[string]interface{}
}

func (googleUser *User) QueryProperty(property string) interface{} {
	if len(googleUser.JsonResponse) > 0 {
		if val, ok := googleUser.JsonResponse[property]; ok == true {
			return val
		}
	}
	return nil
}
