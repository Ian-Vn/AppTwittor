package routes

import (
	"github.com/Ian-Vn/AppTwittor/database"
	"github.com/Ian-Vn/AppTwittor/models"
	"net/http"
)

// Funcion para insertar la relacion a partir del http
func AltaRelacion( w http.ResponseWriter, r *http.Request) {

	// conseguimos el ID
	ID := r.URL.Query().Get("id")

	// Checamos el ID
	if len( ID ) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}
	// Creamos un modelo de "Relacion"
	var relacion models.Relacion

	// Modificamos algunas propiedades
	// Este ID es del usaurio
	relacion.UsuarioID = IDUsuario
	// este ID es del parametro
	relacion.UsuarioRelacionID = ID

	// Llamamos a la funcion que inserta una relacion
	status, err := database.InsertarRelacion(relacion)

	// si ha ocurrido un error
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar relacion", http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar la relacion", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

