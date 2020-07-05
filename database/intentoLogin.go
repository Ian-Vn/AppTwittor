package database

import (
	"github.com/Ian-Vn/AppTwittor/models"
	"golang.org/x/crypto/bcrypt"
)
/*IntentoLogin toma como parametro un email y un password para verificar si existe em ñla base de datos*/
func IntentoLogin( email string, password string ) (models.Usuario, bool) {

	// Llamamos a la funcion de verificacion si existe un usuario
	usuario, encontrado, _ := ExisteUsuario(email)

	// Si encontrado es false entonces no existe el usuario
	if !encontrado {
		// regresamos la estructura vacia y un false
		return usuario, false
	}
	// Si si existe un usuario tomamos el password de la estructura asi como el del parametros y los
	// convertimos a un slice de bytes para usarlos compararlos
	// en donde el primer parametro es el password encriptada
	err := bcrypt.CompareHashAndPassword( []byte(usuario.Password), []byte(password) )
	// Si ha dado un error entonces devolver el usuaruo y un false
	if err != nil {
		return usuario, false
	}
	// si no ha habido error
	return usuario, true
}
