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

func HandlerUpdates(w http.ResponseWriter, r *http.Request) {

	var cau dao.Causa
	var user dao.Usuario

	//var object []dao.Causa
	var err error

	type parameters1 struct {
		Id       string
		Nombre   string
		Apellido string
		Ce       string
	}

	type parameters2 struct {
		Id                 string
		Nro_causa          string
		Caratula           string
		Fiscalia_id        string
		Juzgado_id         string
		Magistrado         string
		Preventor          string
		Preventor_auxiliar string
		Provincia_id       string
		Localidad_id       string
		Domicilio          string
		Nro_sgo            string
		Nro_mto            string
		Tipo_delito        string
		Nombre_fantasia    string
		Fecha              string
		Providencia        string
		Estado             string
		IpAdress           string
		Descripcion        string
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

	countParams := strings.Count(string(b), ":")

	switch countParams {

	case 4:

		err = json.Unmarshal(b, &params1)
		if err != nil {
			msg := "\nFalla durante parseo de parametros del Request: "
			fmt.Printf("%s", msg)
			funciones.ResponseWithJSON(w, 400, msg)
		}

		res, msg := user.LoadDBUpdtData(params1.Id, params1.Ce, params1.Nombre, params1.Apellido)

		if res != 1 {
			msg := "\nFalla durante parseo de parametros del Request: "
			fmt.Printf("%s", msg)
			funciones.ResponseWithJSON(w, 400, msg)
			return

		} else {

			funciones.ResponseWithJSON(w, 200, msg)

		}

	default:

		err = json.Unmarshal(b, &params2)
		if err != nil {
			msg := "\nFalla durante parseo de parametros del Request: "
			fmt.Printf("%s", msg)
			funciones.ResponseWithJSON(w, 400, msg)
			return
		}

		cau.Load2(params2.Nro_causa, params2.Caratula, params2.Juzgado_id, params2.Fiscalia_id, params2.Magistrado, params2.Estado,
			params2.Preventor, params2.Preventor_auxiliar, params2.Provincia_id, params2.Localidad_id, params2.Domicilio, params2.IpAdress,
			params2.Nro_sgo, params2.Nro_mto, params2.Tipo_delito, params2.Nombre_fantasia, params2.Fecha, params2.Providencia,
			"", "", "", "", "")

		token, _ := funciones.GetToken(r.Header)

		res, msg := cau.LoadDBUpdt(params2.Id, token, params2.Descripcion)

		if res == 1 {

			funciones.ResponseWithJSON(w, 200, msg)

		} else {

			funciones.ResponseWithJSON(w, 400, msg)

		}

	}

}
