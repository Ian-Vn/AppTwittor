package routes

import (
	"github.com/Ian-Vn/AppTwittor/database"
	"io"
	"net/http"
	"os"
)

// Funcion para obtener el avatar
func ObtenerAvatar( w http.ResponseWriter, r *http.Request) {

	// obtenemos el ID
	ID := r.URL.Query().Get("id")

	// checamos el id
	if len( ID ) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	// Extraemos el ID del usuario
	perfil, err := database.BuscarPerfil( ID )

	// si no hay un erro
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}

	// Usmaos el usuario
	archivo, err := os.Open("uploads/avatars/" + perfil.Avatar )

	// si ha ocurrido un error mientras se encontraba el archivo
	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
		return
	}

	// Enviamos el archivo al response
	_, err = io.Copy(w, archivo)

	// si ha ocurrido un error al copiar el archivo al response
	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
		return
	}

}

