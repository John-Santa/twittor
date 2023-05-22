package routers

import (
	"encoding/json"
	"net/http"

	"github.com/John-Santa/twittor/db"
	"github.com/John-Santa/twittor/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {

	var user models.Usuario
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", 400)
		return
	}

	_, exist, _ := db.ExistUser(user.Email)
	if exist {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := db.InsertUser(user)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro de usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
