package handlers

import (
	"fmt"
	"net/http"

	"strconv"

	funciones "hello/src/pkg/funciones"
)

func HandlerLogIn(w http.ResponseWriter, r *http.Request) {

	val, msg, val2, nombre := funciones.LogIn(w, r)

	type log struct {
		Nombre string
		Id     string
		Msg    string
	}

	var a log

	a.Id = strconv.Itoa(val2)
	a.Nombre = nombre
	a.Msg = msg

	// Respondo con la pantalla de predicciones :

	if val == 1 {

		n := strconv.Itoa(val2)
		//w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Set-Cookie", "token="+n+"; SameSite=Strict")
		w.Header().Set("Authorization", n)
		w.Header().Set("Transfer-Encoding", "chunked")
		funciones.ResponseWithJSON(w, 200, a)
		//w.WriteHeader(200)

	} else {
		// w.Header().Add("Content-Type", "text")
		// w.WriteHeader(val)
		// dat := []byte(msg)
		// w.Write(dat)
		fmt.Printf("mensage: %s", msg)
		funciones.ResponseWithJSON(w, 400, msg)
	}

}
