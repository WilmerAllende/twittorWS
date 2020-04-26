package models


/* El mensaje del tweet*/
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}