package routes

import (
	"github.com/Ian-Vn/AppTwittor/database"
	"github.com/Ian-Vn/AppTwittor/models"
	"net/http"
)

// Funcion que borra una relacion
func BajaRelacion( w http.ResponseWriter, r *http.Request) {

	// Conseguimos el ID del query params
	ID := r.URL.Query().Get("id")

	// creamos una relacion
	var relacion models.Relacion

	// modificamos la variable anterior
	relacion.UsuarioID = IDUsuario
	relacion.UsuarioRelacionID = ID

	// llamamos a la funcion para eliminar un tweet
	status, err := database.BorrarRelacion( relacion )

	// si ha ocurrido un error
	if err != nil {
		http.Error(w, "Ocurrio un error al borrar una relacion " + err.Error(), http.StatusBadRequest)
		return
	}

	// si el status es false
	if !status {
		http.Error(w, "No se ha logrado insertar la relacion " + err.Error(), http.StatusBadRequest)
		return
	}

	// Escribimos el header
	w.WriteHeader(http.StatusCreated)

}

