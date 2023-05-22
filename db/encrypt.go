package db

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(pass string) (string, error) {
	cost := 8 // El costo es el número de iteraciones que se realizarán para crear el hash. Un valor de 14 es un buen valor por defecto, aunque cuanto mayor sea el valor, más lento será el proceso.
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}
