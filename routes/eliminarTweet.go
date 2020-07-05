package routes

import (
	"github.com/Ian-Vn/AppTwittor/database"
	"net/http"
)

// funcion para eliminar el Tweet
func EliminarTweet( w http.ResponseWriter, r *http.Request ) {

	// Obtenemos el ID del tweet
	ID := r.URL.Query().Get("id")

	// checamos el ID
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	// Obtenemos el valor del ID del token y pasamos como parametros el id del tweet asi como el ID del usuario
	err := database.BorrarTweet( ID, IDUsuario )

	// si no ha ocurrido un error
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar el tweet " + err.Error() , http.StatusBadRequest)
		return
	}

	// Actuqlizamos la respuesta
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

