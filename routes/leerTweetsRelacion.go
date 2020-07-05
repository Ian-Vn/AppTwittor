package routes

import (
	"encoding/json"
	"github.com/Ian-Vn/AppTwittor/database"
	"net/http"
	"strconv"
)

// funcion para leer los tweets de los seguidores
func LeerTweetsSeguidores( w http.ResponseWriter, r *http.Request) {
	// Se obtienen los parametros
	if len( r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Se debe enviar el parametro de la pagina", http.StatusBadRequest)
		return
	}

	// Convertir el parametro a enteri
	pagina, err := strconv.Atoi( r.URL.Query().Get("pagina"))

	// si ha ocurrido un error
	if err != nil {
		http.Error(w, "Se debe enviar el parametro de pagina como entero", http.StatusBadRequest)
		return
	}
	// Se llama a la funcion para leer los tweets
	respuesta, correcto := database.LeerTweetsSeguidores(IDUsuario, pagina)

	// si correcto es false
	if !correcto {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	// se setea el header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}

