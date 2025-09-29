package handlers

import (
	"net/http"

	funciones "hello/src/pkg/funciones"
)

func HandlerModifyPassword(w http.ResponseWriter, r *http.Request) {

	e := funciones.ErrorMsg{}

	res, err := funciones.ModifyPassword(w, r)
	if res != 1 {
		e.SetErrorMsg(err)
		funciones.ResponseWithJSON(w, 500, e)

	} else {

		funciones.ResponseWithJSON(w, 200, "password modificado")

	}

}
