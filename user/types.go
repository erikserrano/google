// Paquete encargado de obtener la información de un usuario en base a un token
// Documentación: https://developers.google.com/accounts/docs/OAuth2
package user

// Estructura encargada de representar la respuesta de Google
type User struct {
    Token        string
	JsonResponse map[string]interface{}
}

// Método encargado de consultar una propiedad en la respuesta de Google
func (googleUser *User) QueryProperty(property string) interface{} {
	if len(googleUser.JsonResponse) > 0 {
		if val, ok := googleUser.JsonResponse[property]; ok == true {
			return val
		}
	}
	return nil
}
