package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	)

// Se realiza una estructura para poder almacenar informacion del token
type Claim struct {
	Email string `json:"email"`
	ID primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	jwt.StandardClaims
}


