package server

type OAuth2WebServer struct {
    Code         string
	JsonResponse map[string]interface{}
}

func (oauth *OAuth2WebServer) QueryProperty(property string) interface{} {
	if len(oauth.JsonResponse) > 0 {
		if val, ok := oauth.JsonResponse[property]; ok == true {
			return val
		}
	}
	return nil
}
