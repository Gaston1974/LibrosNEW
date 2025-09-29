package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"hello/src/pkg/dao"
	funciones "hello/src/pkg/funciones"
)

func HandlerUsuarios(w http.ResponseWriter, r *http.Request) {

	db := funciones.Acceso{}

	_, val0, msg := db.SetCliente()
	if val0 != 1 {
		funciones.ResponseWithJSON(w, 400, msg)
	}

	// var user funciones.Identificable = &dao.Usuario{}
	// var cau funciones.Identificable = &dao.Causa{}

	var user dao.Usuario

	var sqlStatement string
	var object []dao.Usuario
	var err error
	var rows *sql.Rows
	var filtro string

	type parameters struct {
		Ce string
	}

	b, err := io.ReadAll(r.Body)
	//fmt.Println(string(b))
	if err != nil {
		msg := "Falla interna al leer el body del mensaje"
		fmt.Printf("%s", msg)
		funciones.ResponseWithJSON(w, 400, msg)
	}

	params := parameters{}

	err = json.Unmarshal(b, &params)
	if err != nil {
		msg := "\nFalla durante parseo de parametros del Request: "
		fmt.Printf("%s", msg)
		funciones.ResponseWithJSON(w, 400, msg)
	}

	if params.Ce != "" {

		filtro = " WHERE ce = ? "
		sqlStatement = " SELECT  id, first_name , last_name , COALESCE(ce, ''), is_active  FROM usuarios " +
			filtro +
			" ORDER BY first_name;"
		rows, err = db.Cliente.Query(sqlStatement, params.Ce)

	} else {

		filtro = ""
		sqlStatement = " SELECT  id, first_name , last_name , COALESCE(ce, ''), is_active  FROM usuarios " +
			filtro +
			" ORDER BY first_name;"
		rows, err = db.Cliente.Query(sqlStatement)
	}

	if err == nil {
		for rows.Next() {
			user.Load(rows)

			object = append(object, user)
		}

		//fmt.Printf("\nresultado: %v", object)
		funciones.ResponseWithJSON(w, 200, object)
	} else {

		funciones.ResponseWithJSON(w, 400, err.Error())

	}

	defer db.Cliente.Close()

}
