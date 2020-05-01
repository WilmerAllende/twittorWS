package routers

import (
	"net/http"
	"github.com/WilmerAllende/twittorWS/bd"
	"github.com/WilmerAllende/twittorWS/models"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request)  {
	
	ID := r.URL.Query().Get("id")
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	if len(ID) < 1 {
		http.Error(w,"Debe enviar el parametro id en la relacion",http.StatusBadRequest)
		return
	}

	status,err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Error al intentar realizar borro de relacion" + err.Error(),http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado borrar la relacion" ,http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

}