package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ricardobeta/twittor/bd"
	"github.com/ricardobeta/twittor/jwt"
	"github.com/ricardobeta/twittor/models"
)


func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario
	err:= json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario o Contraseña inválidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email Requerido"+err.Error(), 400)
		return
	}

	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if existe == false {
		http.Error(w, "Usuario o Contraseña inválidos", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el Token"+err.Error(), 400)
		return
	}
	resp := models.RespuestaLogin {
		Token: jwtKey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24*time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: jwtKey,
		Expires: expirationTime,
	})

}