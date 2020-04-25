package bd

import (
	"context"
	"time"
	"github.com/WilmerAllende/twittorWS/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* Verifica si existe un email en la BD*/
func ChequeoYaExisteUsuario(email string) (models.Usuario,bool,string) {

	ctx, cancel := context.WithTimeout(context.Background(),15*time.Second)
	defer cancel()

	db := MongoCN.Database("wittorws")
	col := db.Collection("usuarios")

	condicion := bson.M{"email":email}
	
	var resultado  models.Usuario
	err := col.FindOne(ctx,condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado,false,ID
	}
	return resultado,false,ID

	
}