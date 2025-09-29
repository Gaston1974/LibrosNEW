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

func HandlerAltas(w http.ResponseWriter, r *http.Request) {

	var cau dao.Causa
	var user dao.Usuario
	var prev dao.Preventor
	var inter dao.Interventor

	var err error
	var res int
	var msg string

	type parameters1 struct {
		Nombre   string
		Apellido string
		Ce       string
		Password string
	}

	type parameters2 struct {
		Nro_causa          string
		Caratula           string
		Juzgado            string
		Fiscalia           string
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
		Nombre_archivo     string
		Ruta_archivo       string
		Tipo_documento     string
		Tamano             string
		Nota_causa         string
	}

	type parameters3 struct {
		Nombre string
	}

	type parameters4 struct {
		Interventor string
		Nombre      string
	}

	b, err := io.ReadAll(r.Body)
	fmt.Println(string(b))
	if err != nil {
		msg := "Falla interna al leer el body del mensaje"
		fmt.Printf("%s", msg)
		funciones.ResponseWithJSON(w, 400, msg)
		return
	}

	params1 := parameters1{}
	params2 := parameters2{}
	params3 := parameters3{}
	params4 := parameters4{}

	countParams := strings.Count(string(b), ":")

	switch countParams {

	case 1:

		err = json.Unmarshal(b, &params3)
		if err != nil {
			msg := "\nFalla durante parseo de parametros del Request: "
			fmt.Printf("%s", msg)
			funciones.ResponseWithJSON(w, 400, msg)
			return
		}

		prev.Load2(params3.Nombre)

		res, msg = prev.LoadDB()

		if res == 1 {

			funciones.ResponseWithJSON(w, 200, msg)

		} else {

			funciones.ResponseWithJSON(w, 400, msg)

		}

	case 2:

		err = json.Unmarshal(b, &params4)
		if err != nil {
			msg := "\nFalla durante parseo de parametros del Request: "
			fmt.Printf("%s", msg)
			funciones.ResponseWithJSON(w, 400, msg)
			return
		}

		inter.Load2(params4.Nombre)

		res, msg = inter.LoadDB(params4.Interventor)

		if res == 1 {

			funciones.ResponseWithJSON(w, 200, msg)

		} else {

			funciones.ResponseWithJSON(w, 400, msg)

		}

	case 4:

		err = json.Unmarshal(b, &params1)
		if err != nil {
			msg := "\nFalla durante parseo de parametros del Request: "
			fmt.Printf("%s", msg)
			funciones.ResponseWithJSON(w, 400, msg)
			return
		}

		user.Load2(params1.Nombre, params1.Apellido, params1.Ce, params1.Password)

		res, msg = user.LoadDB()

		if res == 1 {

			funciones.ResponseWithJSON(w, 200, msg)

		} else {

			funciones.ResponseWithJSON(w, 400, msg)

		}

	default:

		err = json.Unmarshal(b, &params2)
		if err != nil {
			msg := "\nFalla durante parseo de parametros del Request: "
			fmt.Printf("%s", msg)
			funciones.ResponseWithJSON(w, 400, msg)
		}

		cau.Load2(params2.Nro_causa, params2.Caratula, params2.Juzgado, params2.Fiscalia, params2.Magistrado, params2.Estado,
			params2.Preventor, params2.Preventor_auxiliar, params2.Provincia_id, params2.Localidad_id, params2.Domicilio, "",
			params2.Nro_sgo, params2.Nro_mto, params2.Tipo_delito, params2.Nombre_fantasia, params2.Fecha, params2.Providencia,
			params2.Nombre_archivo, params2.Ruta_archivo, params2.Tipo_documento, params2.Tamano, params2.Nota_causa)

		token, _ := funciones.GetToken(r.Header)

		res, msg := cau.LoadDB("alta", token, "", "")

		if res == 1 {

			funciones.ResponseWithJSON(w, 200, msg)

		} else {

			funciones.ResponseWithJSON(w, 400, msg)

		}

	}

}
