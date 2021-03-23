package middleW

import (
	"net/http"
	"github.com/ricardobeta/twittor/bd"
)


func ChequeoBD(next http.HandlerFunc)  http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión Perdida", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}