package routes

import (
	"encoding/json"
	"github.com/Ian-Vn/AppTwittor/database"
	"github.com/Ian-Vn/AppTwittor/models"
	"net/http"
)

// Esta funcion se encarga de realizar la consulta para saber si existe la relacion entre dos usuarios
func ConsultaRelacion( w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	// creamos una variable del tipo relacion
	var relacion models.Relacion

	// modifcamos algunas propiedades
	relacion.UsuarioID = IDUsuario
	relacion.UsuarioRelacionID = ID

	// Creamos una variable del tipo RespuestaConsultaRelacion
	var resp models.RespuestaConsultaRelacion

	// llamamos a la funcion consultarelacion para saber si dos usuarios tienen relacion
	status, err := database.ConsultarRelacion( relacion )

	// si ha ocurrido un error es decir que no exista relacion
	if err != nil || !status {
		resp.Status = false
	} else {
		resp.Status = true
	}

	// seteamos la respuesta
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader( http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

