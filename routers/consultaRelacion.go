package routers

import (
	"encoding/json"
	"net/http"
	"github.com/WilmerAllende/twittorWS/bd"
	"github.com/WilmerAllende/twittorWS/models"
)

func ConsultaRelacion(w http.ResponseWriter, r *http.Request)  {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w,"Debe enviar el parametro id en la relacion",http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	var resp models.RespuestaConsultaRelacion
	status,err := bd.ConsultoRelacion(t)
	if err != nil || status == false  {
		resp.Status = false
	} else{
		resp.Status = true
	}

	w.Header().Set("Content-type","application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)
}