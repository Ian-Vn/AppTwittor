package routes

import (
	"encoding/json"
	"github.com/Ian-Vn/AppTwittor/database"
	"github.com/Ian-Vn/AppTwittor/models"
	"net/http"
	"time"
)

// Funcion para grabar el tweet
func GrabarTweet( w http.ResponseWriter, r *http.Request ) {
	// Creamos una variable para para almacenar el tweet
	var mensaje models.Tweet

	// Decodificamos el json a nuestra estructura
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	if err != nil {
		http.Error(w, "No se ha podido obtener el mensaje", 400)
		return
	}
	// Se crea una variable para almacenar el Tweet
	registro := models.TweetGrabar{ UserID: IDUsuario, Mensaje: mensaje.Mensaje, Fecha: time.Now() }

	// llamamos a la funcion de insertar Tweet
	_, status, err := database.InsertarTweet( registro )

	// si ha ocurrido un error
	if err != nil {
		http.Error(w, "Ocurrio un error al insertar el registro, reintentar" + err.Error(), 400)
		return
	}

	// si no se ha grabado nada
	if !status {
		http.Error(w, "No se ha logrado insertar el Tweet" , 400)
		return
	}
	w.WriteHeader(http.StatusCreated)


}

