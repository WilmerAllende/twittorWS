package models

import "time"

/* Formato estructura del tweet*/
type GraboTweet struct {
	UserID string `bson:"userid" json:"userid,omitempty"`
	Mensaje string `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha time.Time `bson:"fecha" json:"fecha,omitempty"`
}