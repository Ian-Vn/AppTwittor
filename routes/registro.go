package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Ian-Vn/AppTwittor/database"
	"github.com/Ian-Vn/AppTwittor/models"
)

/*Registro es la funcion para crear un registro en la base de datos */
func Registro(w http.ResponseWriter, r *http.Request) {

	// Creando una variable del tipo usuario
	var usuario models.Usuario

	// Creamos un decoder desde el request
	// A partir del body que vendra en formato json se hace un decoder y decodifica hacia la estructura usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)

	// Si ha ocurrido un error al decodificar el json hacia la estructura entonces
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	// A este punto la estructura ya tiene la informacion del body
	// Verificamos si el email no viene vacio
	if len(usuario.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	// Si el password viene con una longitud menor a 6
	if len(usuario.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", 400)
		return
	}

	// Llamamos a la funcion para verificar que el usuario ya existe en la bd
	_, encontrado, _ := database.ExisteUsuario(usuario.Email)

	// Si es true entonces se encontro el usuario y salimos de la funcion
	if encontrado == true {
		http.Error(w, "Ya existe el usuario con ese email", 400)
		return
	}

	// Insercion de registro
	_, status, err := database.InsertarRegistro(usuario)

	// Si ha ocurrido un error
	if err != nil {
		http.Error(w, "Ocurrio un error al insertar el registro "+err.Error(), 400)
		return
	}

	// Si el estado es falso
	if status == false {
		http.Error(w, "No se ha logrado insertar el registro", 400)
		return
	}

	// Mandamos el codigo de exito
	w.WriteHeader(http.StatusOK)
}
