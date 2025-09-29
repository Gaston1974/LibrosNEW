package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"hello/src/pkg/dao"
	funciones "hello/src/pkg/funciones"
)

func HandlerBajas(w http.ResponseWriter, r *http.Request) {

	var cau dao.Causa
	var user dao.Usuario
	var prev dao.Preventor
	var inter dao.Interventor

	var err error
	var res int
	var msg string

	type parameters1 struct {
		Ce string
	}

	type parameters2 struct {
		Nro_causa string
		Motivo    string
		Estado    string
	}

	type parameters3 struct {
		Motivo string
		Nombre string
	}

	b, err := io.ReadAll(r.Body)
	//fmt.Println(string(b))
	if err != nil {
		msg := "Falla interna al leer el body del mensaje"
		fmt.Printf("%s", msg)
		funciones.ResponseWithJSON(w, 400, msg)
	}

	params1 := parameters1{}
	params2 := parameters2{}
	params3 := parameters3{}

	countParams := strings.Count(string(b), ":")

	switch countParams {

	case 1:

		err = json.Unmarshal(b, &params1)
		if err != nil {
			msg := "\nFalla durante parseo de parametros del Request: "
			fmt.Printf("%s", msg)
			funciones.ResponseWithJSON(w, 400, msg)
			return
		}

		user.Load3(params1.Ce)

		res, msg = user.LoadBaja()

		if res == 1 {

			funciones.ResponseWithJSON(w, 200, msg)

		} else {

			funciones.ResponseWithJSON(w, 400, msg)

		}

	case 2:

		err = json.Unmarshal(b, &params3)
		if err != nil {
			msg := "\nFalla durante parseo de parametros del Request: "
			fmt.Printf("%s", msg)
			funciones.ResponseWithJSON(w, 400, msg)
			return
		}

		switch params3.Motivo {

		case "preventor":

			prev.Load2(params3.Nombre)

			res, msg = prev.LoadBaja()

			if res == 1 {

				funciones.ResponseWithJSON(w, 200, msg)

			} else {

				funciones.ResponseWithJSON(w, 400, msg)

			}

		default:

			res, msg = inter.LoadBaja(params3.Motivo, params3.Nombre)

			if res == 1 {

				funciones.ResponseWithJSON(w, 200, msg)

			} else {

				funciones.ResponseWithJSON(w, 400, msg)

			}

		}

	default:

		err = json.Unmarshal(b, &params2)
		if err != nil {
			msg := "\nFalla durante parseo de parametros del Request: "
			fmt.Printf("%s", msg)
			funciones.ResponseWithJSON(w, 400, msg)
			return
		}

		cau.Load3(params2.Nro_causa)

		token, _ := funciones.GetToken(r.Header)

		res, msg := cau.LoadBaja("", token, params2.Motivo, "eliminacion", params2.Estado)

		if res == 1 {

			funciones.ResponseWithJSON(w, 200, msg)

		} else {

			funciones.ResponseWithJSON(w, 400, msg)

		}

	}

}
