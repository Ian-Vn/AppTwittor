package middlewares

import (
	"net/http"

	"github.com/Ian-Vn/AppTwittor/database"
)

/*ChequeoBD es un middleware que toma como parametro otra funcion y si no hay un error
  entonces le pasa el control a la funcion pasada como parametro */
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	// Como handlerfunc regresa una funcion
	return func(w http.ResponseWriter, r *http.Request) {
		if database.ChequeoConexion() == 0 {
			http.Error(w, "Conexion perdida con la base de datos", 500)
			return
		}
		// llamamos a la funcion pasada como parametro pasandole el response y el request
		next.ServeHTTP(w, r)
	}
}
