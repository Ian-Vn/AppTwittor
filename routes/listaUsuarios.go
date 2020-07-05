package routes

import (
	"encoding/json"
	"github.com/Ian-Vn/AppTwittor/database"
	"net/http"
	"strconv"
)

// Funcion para listar usuarios
func ListaUsuarios(w http.ResponseWriter, r *http.Request) {

	// Se capturan los parametros
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	// Se convierte la variable page a int
	pagTemp, err := strconv.Atoi( page )

	// si ha ocurrido un error
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina como entero mayor a 0", http.StatusBadRequest)
		return
	}
	// se convierte la pagina a int64
	pag := int64(pagTemp)

	// se llama a la funcion para leer todos los usuarios
	result, status := database.LeerTodosUusuarios(IDUsuario, pag, search, typeUser)

	// si ha ocurrido un error
	if !status {
		http.Error(w, "Errpr al leer todos los usuarios", http.StatusBadRequest)
		return
	}

	// seteamos el header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

