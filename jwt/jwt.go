package jwt

import (
	"time"

	"github.com/John-Santa/twittor/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJWT(user models.Usuario) (string, error) {
	mykey := []byte("MySuperSecretKey")

	payload := jwt.MapClaims{
		"email":            user.Email,
		"nombre":           user.Nombre,
		"apellidos":        user.Apellidos,
		"fecha_nacimiento": user.FechaNacimiento,
		"biografia":        user.Biografia,
		"ubicacion":        user.Ubicacion,
		"sitio_web":        user.SitioWeb,
		"_id":              user.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	myToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := myToken.SignedString(mykey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
