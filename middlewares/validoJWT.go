package middlewares

import (
	"github.com/Ian-Vn/AppTwittor/routes"
	"net/http"
)

// Esta funcion permite validad el jwt que viene de la peticion
func ValidoJWT( next http.HandlerFunc ) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request  ){

		// Llamando a la funcion para procesar el token
		_,_,_,err := routes.ProcesoToken( r.Header.Get("Authorization") )
		// Si ha habido un error al procesar el toke
		if err != nil {
			http.Error(w, "Error en el token ! " + err.Error(), http.StatusBadRequest)
			return
		}
		// si no hubo error le pasamos el control a la siguiente funcion
		next.ServeHTTP(w, r)
	}
}
