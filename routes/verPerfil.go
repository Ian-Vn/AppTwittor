package routes

import (
	"encoding/json"
	"github.com/Ian-Vn/AppTwittor/database"
	"net/http"
)

// FUncion que devulve en json el perfil de un usuario
func VerPerfil(w http.ResponseWriter, r *http.Request) {

	// Extraemos el queryparam id
	ID := r.URL.Query().Get("id")
	// Este id es una cadena con representacion del id de mongo
	// Si el id es vacio
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el paramentro ID", http.StatusBadRequest)
		return
	}

	// Le pasamsoe el id a la funcion buscarPerfil para que se obtenga el usuario y un error
	perfil, err := database.BuscarPerfil(ID)

	// si ha ocurrido un error
	if err != nil {
		http.Error(w, "Ocurrio un error el intentar buscar el registro " + err.Error(), 400)
		return
	}
	perfil.Password = ""
	// Si no ha ocurrido un error
	w.Header().Set("content-type", "application/json")
	// sobreescribimos el header y enviarmos los datos en json
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode( perfil )
}

