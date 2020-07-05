package models

// Se guarda el token generado 
type RespuestaLogin struct {
	Token string `json:"token, omitempty"`
}