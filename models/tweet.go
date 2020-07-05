package models

// Se almacena un tweet
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`

}