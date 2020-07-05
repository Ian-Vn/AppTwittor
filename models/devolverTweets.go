package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Estructura  para devolver los tweets
type DevovlerTweets struct {
	ID primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID string `bson:"userid" json:"userId,omitempty"`
	Mensaje string `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha time.Time `bson:"fecha" json:"fecha,omitempty"`
}



