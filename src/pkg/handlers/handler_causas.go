package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"hello/src/pkg/dao"
	funciones "hello/src/pkg/funciones"
)

func HandlerCausas(w http.ResponseWriter, r *http.Request) {

	db := funciones.Acceso{}

	_, val0, msg := db.SetCliente()
	if val0 != 1 {
		funciones.ResponseWithJSON(w, 400, msg)
	}

	// var user funciones.Identificable = &dao.Usuario{}
	// var cau funciones.Identificable = &dao.Causa{}

	var cau dao.Causa
	var hist dao.Causa_historico

	var sqlStatement string
	var object []dao.Causa
	var object2 []dao.Causa_historico
	var err error
	var rows *sql.Rows

	type parameters struct {
		Nro_causa string
	}

	b, err := io.ReadAll(r.Body)
	//fmt.Println(string(b))
	if err != nil {
		msg := "Falla interna al leer el body del mensaje"
		fmt.Printf("%s", msg)
		funciones.ResponseWithJSON(w, 400, msg)
		return
	}

	defer db.Cliente.Close()

	params := parameters{}

	path := r.URL.Path

	vector := strings.Split(path, "/")
	value := len(vector)

	if value == 4 {

		err = json.Unmarshal(b, &params)
		if err != nil {
			msg := "\nFalla durante parseo de parametros del Request: "
			fmt.Printf("%s", msg)
			funciones.ResponseWithJSON(w, 400, msg)
			return
		}

		sqlStatement = " SELECT numero_causa, accion, usuario_id, descripcion, COALESCE(ip_address, '') , creado_en " +
			" FROM historial_causas " +
			" WHERE numero_causa = ? " +
			" ORDER BY numero_causa;"

		rows, err = db.Cliente.Query(sqlStatement, params.Nro_causa)

		if err == nil {
			for rows.Next() {

				hist.Load(rows)

				object2 = append(object2, hist)

			}

			funciones.ResponseWithJSON(w, 200, object2)

		} else {

			funciones.ResponseWithJSON(w, 400, "falla en la BD al ib insertar"+"\n"+err.Error())
		}

		return

	}

	err = json.Unmarshal(b, &params)
	if err != nil {
		msg := "\nFalla durante parseo de parametros del Request: "
		fmt.Printf("%s", msg)
		funciones.ResponseWithJSON(w, 400, msg)
		return
	}

	if params.Nro_causa != "" {

		sqlStatement = " SELECT numero_causa , COALESCE(caratula, 'VACIO') , j.nombre , f.nombre , COALESCE(a_cargo_del_magistrado, 'VACIO') , " +
			" preventor , preventor_auxiliar,  provincia_id , localidad_id , COALESCE(domicilio, '')  , " +
			" COALESCE(nro_mto, 'VACIO') , COALESCE(nro_sgo, 'VACIO')  , " +
			" COALESCE(tipo_delito, 'VACIO'), COALESCE(nombre_fantasia, 'VACIO'), COALESCE(fecha_llegada, 'VACIO') , COALESCE(providencia, 'VACIO'), " +
			" COALESCE(estado, 'VACIO'), '' as ipaddress , " +
			" COALESCE(d.nombre_archivo, 'VACIO'),  COALESCE(d.ruta_archivo, 'VACIO'), COALESCE(d.tipo_documento, 'VACIO'), " +
			" COALESCE(d.tamano, 'VACIO'), " +
			" COALESCE(d.subido_por, 'VACIO'), COALESCE(n.contenido, 'VACIO') " +
			" FROM causas c " +
			" LEFT JOIN documentos_causa d ON c.id = d.causa_id " +
			" LEFT JOIN notas_causa n ON c.id = n.causa_id " +
			" INNER JOIN fiscalias f ON fiscalia_id = f.id " +
			" INNER JOIN juzgados j ON juzgado_id = j.id  " +
			" WHERE numero_causa = ? " +
			" ORDER BY numero_causa;"

		rows, err = db.Cliente.Query(sqlStatement, params.Nro_causa)

	} else {

		sqlStatement = " SELECT numero_causa , COALESCE(caratula, 'VACIO') , j.nombre , f.nombre , COALESCE(a_cargo_del_magistrado, 'VACIO') , " +
			" preventor , preventor_auxiliar,  provincia_id , localidad_id , COALESCE(domicilio, '')  , " +
			" COALESCE(nro_mto, 'VACIO') , COALESCE(nro_sgo, 'VACIO')  , " +
			" COALESCE(tipo_delito, 'VACIO'), COALESCE(nombre_fantasia, 'VACIO'), COALESCE(fecha_llegada, 'VACIO') , COALESCE(providencia, 'VACIO'), " +
			" COALESCE(estado, 'VACIO'), '' as ipaddress , " +
			" COALESCE(d.nombre_archivo, 'VACIO'),  COALESCE(d.ruta_archivo, 'VACIO'), COALESCE(d.tipo_documento, 'VACIO'), " +
			" COALESCE(d.tamano, 'VACIO'), " +
			" COALESCE(d.subido_por, 'VACIO'), COALESCE(n.contenido, 'VACIO') " +
			" FROM causas c " +
			" LEFT JOIN documentos_causa d ON c.id = d.causa_id " +
			" LEFT JOIN notas_causa n ON c.id = n.causa_id " +
			" INNER JOIN fiscalias f ON fiscalia_id = f.id " +
			" INNER JOIN juzgados j ON juzgado_id = j.id  " +
			" ORDER BY providencia;"

		rows, err = db.Cliente.Query(sqlStatement)

	}

	if err == nil {
		for rows.Next() {
			cau.Load(rows)

			object = append(object, cau)
		}

		funciones.ResponseWithJSON(w, 200, object)
	} else {

		funciones.ResponseWithJSON(w, 400, "falla en la consulta a la BD"+"\n"+err.Error())
	}

}
