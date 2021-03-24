package bd

import "golang.org/x/crypto/bcrypt"
/* Algoritmo 2^costo  mayor Costo mayor seguridad, mayor demora*/
func EncriptarPassword(pass string) (string, error) {
	costo := 6
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}