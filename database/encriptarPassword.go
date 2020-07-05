package database

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword es una funcion que devuelve un string encriptado*/
func EncriptarPassword(password string) (string, error) {

	// Encriptamos el password con bcrypt
	// password es un slice de bytes
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(passwordBytes), err

}
