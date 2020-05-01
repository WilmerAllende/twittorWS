package routers

import (
	"net/http"
	"github.com/WilmerAllende/twittorWS/bd"
	"github.com/WilmerAllende/twittorWS/models"
)

func AltaRelacion(w http.ResponseWriter, r *http.Request)  {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w,"Debe enviar el parametro id en la relacion",http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status,err := bd.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "Error al intentar realizar registro de relacion" + err.Error(),http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar la relacion" ,http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

}