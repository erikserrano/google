// Paquete encargado de obtener un token de Google en base a un código de autorización
// Documentación: https://developers.google.com/accounts/docs/OAuth2WebServer 
package server

// Estructura encargada de almacenar código de respuesta de Google
type OAuth2WebServer struct {
    Code         string
	JsonResponse map[string]interface{}
}


// Método encargado de consultar una propiedad en la respuesta de Google
func (oauth *OAuth2WebServer) QueryProperty(property string) interface{} {
	if len(oauth.JsonResponse) > 0 {
		if val, ok := oauth.JsonResponse[property]; ok == true {
			return val
		}
	}
	return nil
}
