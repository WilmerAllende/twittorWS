package bd

import (
	"context"
	"time"
	"github.com/WilmerAllende/twittorWS/models"
)

func InsertoRelacion(t models.Relacion) (bool,error)  {
	
	ctx, cancel := context.WithTimeout(context.Background(),15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittorws")
	col := db.Collection("relacion")

	_,err := col.InsertOne(ctx,t)
	if err != nil {
		return false,err
	}
	return true,nil
}