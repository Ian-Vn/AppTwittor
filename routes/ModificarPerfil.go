package routes

import (
	"encoding/json"
	"github.com/Ian-Vn/AppTwittor/database"
	"github.com/Ian-Vn/AppTwittor/models"
	"net/http"
)


func ModificarPerfil ( w http.ResponseWriter, r *http.Request ) {
	// Creamos una variable
	var usuario models.Usuario
	var status bool
	// Decodificamos el json hacia una estructura
	err := json.NewDecoder(r.Body).Decode(&usuario)

	// Si ha ocurrido un error
	if err != nil {
		http.Error(w, "Datos incorrectos"+ err.Error(), 400)
		return
	}

	// llamamos a la funcion  para modificar el registro y usamso el ID que fue modificado en el middleware
	status, err = database.ModificarRegistro( usuario, IDUsuario )

	// si ha ocurrido un error
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro, reintentar nuevamente" + err.Error(), 400)
		return
	}

	// Chequeo de nuevo el error
	if !status {
		http.Error(w, "No se ha logrado modificar el registro del usuario ", 400)
		return
	}
	// si no ha ocurrido un error
	w.WriteHeader(http.StatusOK)



}

