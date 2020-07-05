package routes

import (
	"errors"
	"github.com/Ian-Vn/AppTwittor/database"
	"github.com/Ian-Vn/AppTwittor/models"
	jwt "github.com/dgrijalva/jwt-go"
	"strings"
)

// Variables para almacenar los valores
var Email string
var IDUsuario string

// Funcion para procesar el toke y extraer sus valores
func ProcesoToken(token string) ( * models.Claim, bool, string, error ){
	// Escribimos la clave
	clave := []byte("password")

	// Declaramos una variable del apuntador a Claim
	claims := &models.Claim{}

	// A partir deñ token se haran algunas operaciones
	// El token como se extrae del header viene acompañador de una palabra Bearer
	// por lo cual necesitamos extraer esa palabra y el token, por medio de la funcion
	// Split se separa el token dando asi dos elementos, un elemento vacio por encontrar el separador
	// y el otro elemento la cadena restante
	splitToken := strings.Split(token, "Bearer")

	// Si es diferente de 2 entonces el token no ha venido bien
	if len(splitToken) != 2 {
		// Se retorna la estructura vacia, un false, una cadena vacia y un nuevo error
		return claims, false, "", errors.New("formato de token invalido")
	}
	// quitamos los espacios en blanco con respecto al token y devolvemos otro token
	tk := strings.TrimSpace(splitToken[1])

	// Validacion del token, lo que se hara es decodificar el token para despues insertar el payload a la estructura definida por claims
	tkn, err := jwt.ParseWithClaims(tk, claims, func( token * jwt.Token )(interface {}, error) {
		// Regresamos la clave y un nil
		return clave, nil
	})

	// Si no ha habido un error
	if err == nil {
		// Llamamos a la funcion de existir usuario
		_, encontrado, _ := database.ExisteUsuario( claims.Email )

		// Si es true entonces se encontro el usuario
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	// Si el token no ha sido valido, es decir que se codifico un token vacio
	if !tkn.Valid {
		return claims, false, "", err
	}

	return claims, true, IDUsuario, nil
}
