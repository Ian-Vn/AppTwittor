package jwt

import (
	"github.com/Ian-Vn/AppTwittor/models"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

// Funcion para generar el jwt
func GenerarJWT( usuario models.Usuario ) (string, error) {
	clave := []byte("password")

	// Se simula la decodificacion de JSON a un mapclaims
	payload := jwt.MapClaims{ "email": usuario.Email,
		"nombre": usuario.Nombre,
		"apellidos": usuario.Apellidos,
		"fecha_nacimiento": usuario.FechaNacimiento,
		"biografia": usuario.Biografia,
		"ubicacion": usuario.Ubicacion,
		"sitioweb": usuario.SitioWeb,
		"_id": usuario.ID.Hex(),
		"exp": time.Now().Add( 24 * time.Hour ).Unix(),
	}

	// Creamos la configuracion de la firma usando el algoritmo de encriptacion y el payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	// Generamos el token
	tokenStr, err := token.SignedString(clave)

	// Si hubo un error devolvermos el tokenStr que estara vacio y el error
	if err != nil {
		return tokenStr, err
	}
	// Si no ha ocurrido un error se devuelve el nil
	return tokenStr, nil

}



