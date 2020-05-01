package routers

import (
	"net/http"
	"github.com/WilmerAllende/twittorWS/bd"
)

func EliminarTweet(w http.ResponseWriter, r *http.Request)  {
	
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w,"Debe enviar el parametro id",http.StatusBadRequest)
		return
	}

	err := bd.BorroTweet(ID,IDUsuario)
	if err != nil {
		http.Error(w,"Hubo un error al intentar borrar un tweet" + err.Error(),http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusCreated)
}