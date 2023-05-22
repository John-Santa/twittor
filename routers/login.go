package routers

import (
	"encoding/json"
	"net/http"

	"github.com/John-Santa/twittor/db"
	"github.com/John-Santa/twittor/jwt"
	"github.com/John-Santa/twittor/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var user models.Usuario
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña invalidos "+err.Error(), 400)
		return
	}
	if len(user.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}
	document, exist := db.Login(user.Email, user.Password)
	if !exist {
		http.Error(w, "Usuario y/o contraseña invalidos", 400)
		return
	}
	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el Token correspondiente "+err.Error(), 400)
		return
	}
	resp := models.LoginResponse{
		Token: jwtKey,
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: jwtKey,
	})
}
