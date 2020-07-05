package routes

import (
	"encoding/json"
	"github.com/Ian-Vn/AppTwittor/database"
	"github.com/Ian-Vn/AppTwittor/jwt"
	"github.com/Ian-Vn/AppTwittor/models"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request  )  {
	// Se setea el tipo del contenido
	w.Header().Add("Conten-Type", "application/json")

	// Creamos un modelo
	var usuario models.Usuario

	// Verificamos el body para obtener el email y el password
	err := json.NewDecoder(r.Body).Decode(&usuario)

	// si ha habido un error en la transformacion de los datos mandamos un error
	if err != nil {
		http.Error(w, "Usuario y/o contraseña invalidos" + err.Error(), 400)
		return
	}

	// Si no ha habido error entonces verificamos algunos datos
	// verificamos el email
	if len(usuario.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	// Verificamos si el email existe con la funcion de IntentoLogin que devuelve el usuario con el email y esl password ademas
	// de un valor booleano
	documento, existe := database.IntentoLogin(usuario.Email, usuario.Password)

	// Verificamos si existe es false en caso de que no exista en la bd
	if !existe {
		http.Error(w, "Usuario y/o contraseña invalidos", 400)
		return
	}

	// Creamos el jwt a partir del usuario (documento)
	jwtKey, err := jwt.GenerarJWT( documento )

	// SI ha ocurrido un error mientras se creaba el jwt
	if err != nil {
		http.Error(w, "Ha ocurrido un error mientras se generaba el token" + err.Error(), 400)
		return
	}

	// Si no ha ocurrido un error entonces creamos una variable del stipo RespuestaLogin
	resp := models.RespuestaLogin { Token: jwtKey }

	// Seteamos el header
	w.Header().Set("Content-Type", "application/json")
	// Seteamos el status del header
	w.WriteHeader(http.StatusCreated )

	// Para mandar el json hacia la respuesta usamos el enconder de json para codificar la estructura a JSON
	json.NewEncoder( w ).Encode(resp)


}
