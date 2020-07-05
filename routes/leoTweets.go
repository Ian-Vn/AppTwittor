package routes

import (
	"encoding/json"
	"github.com/Ian-Vn/AppTwittor/database"
	"net/http"
	"strconv"
)

// funcion para leer los tweets
func LeoTweets( w http.ResponseWriter, r *http.Request ) {

	// Se obtiene el ID
	ID := r.URL.Query().Get("id")

	// Se verifica el el id
	if len( ID ) < 1 {
		http.Error(w, "Se debe enviar el ID", http.StatusBadRequest)
		return
	}

	// se verifica el tamaño de la pagina
	if len( r.URL.Query().Get("pagina")  ) < 1 {
		http.Error(w, "Se debe enviar el ID", http.StatusBadRequest)
		return
	}

	// Si la pagina ha sido correcta se convierte a entero
	pagina, err := strconv.Atoi( r.URL.Query().Get("pagina") )

	// si ha ocurrido un error
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	// se convierte a un int64
	pag := int64(pagina)

	// se llama a la funcion LeoTweets el cual regresa un slice de apuntadores al tipo DevolverTweets
	respuesta, correcto := database.LeerTweets( ID, pag )

	// si correcto es false es porque ha ocurrido un error
	if !correcto {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	// Se envian los datos
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)


}

