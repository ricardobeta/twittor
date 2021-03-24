package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ricardobeta/twittor/bd"
	"github.com/ricardobeta/twittor/models"
)

/*Registro es la funcion oara crear en la BD el reguistro de usuaro*/
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	// validaciones
	if len(t.Email) == 0 {
		http.Error(w, "Se requiere el Email de usuario"+err.Error(), 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "la contraseña no es segura"+err.Error(), 400)
		return
	}
	//
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		http.Error(w, "Ya exste un Usuario Registrado con ese Email", 400)
		return
	}
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un Error al intentar realizar el registro de Usuario"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Ocurrio un Error al intentar  registro de Usuario"+err.Error(), 400)
		return
	}
	// Inserto Correctamente
	w.WriteHeader(http.StatusCreated)
}
