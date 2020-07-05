package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Ian-Vn/AppTwittor/middlewares"
	"github.com/Ian-Vn/AppTwittor/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores es la funcion para que setea el puerto ademas crea u  router para las rutas y permite el cors*/
func Manejadores() {
	// Creamos un router vacio, devuelve un apuntador a Router
	router := mux.NewRouter()

	// Ruta de registro
	router.HandleFunc("/registro", middlewares.ChequeoBD(routes.Registro)).Methods("POST")
	router.HandleFunc("/login", middlewares.ChequeoBD(routes.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlewares.ChequeoBD( middlewares.ValidoJWT(routes.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlewares.ChequeoBD( middlewares.ValidoJWT(routes.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlewares.ChequeoBD( middlewares.ValidoJWT(routes.GrabarTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middlewares.ChequeoBD( middlewares.ValidoJWT(routes.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middlewares.ChequeoBD( middlewares.ValidoJWT(routes.EliminarTweet))).Methods("DELETE")
	router.HandleFunc("/subirAvatar", middlewares.ChequeoBD( middlewares.ValidoJWT(routes.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlewares.ChequeoBD( routes.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/subirBanner", middlewares.ChequeoBD( middlewares.ValidoJWT(routes.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlewares.ChequeoBD( routes.ObtenerBanner)).Methods("GET")
	router.HandleFunc("/altaRelacion", middlewares.ChequeoBD( middlewares.ValidoJWT(routes.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlewares.ChequeoBD( middlewares.ValidoJWT(routes.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middlewares.ChequeoBD( middlewares.ValidoJWT(routes.ConsultaRelacion))).Methods("GET")
	router.HandleFunc("/listaUsuarios", middlewares.ChequeoBD( middlewares.ValidoJWT(routes.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/leoTweetsSeguidores", middlewares.ChequeoBD( middlewares.ValidoJWT(routes.LeerTweetsSeguidores))).Methods("GET")
	// Verificamos la variable de entorno PORT
	PORT := os.Getenv("PORT")
	// Si no esta definida la variable de entorno definida entonces devuelve string vacio
	if PORT == "" {
		// Actualizamos el puerto
		PORT = "8080"
	}
	// Definimos el cors
	handler := cors.AllowAll().Handler(router)
	// Imprimimos en el log si ha ocurrido un error cuando se levanta el servidor
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
