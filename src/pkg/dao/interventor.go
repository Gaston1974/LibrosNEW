package dao

import (
	"database/sql"
	"fmt"
	fun "hello/src/pkg/funciones"
	"log"
	"strconv"
)

type Interventor struct {
	Nombre string
}

func (u *Interventor) Load(pp *sql.Rows) string {

	err := pp.Scan(&u.Nombre)

	if err != nil {
		msg := "\n Error en la carga del objeto"
		fmt.Printf("%s \n %s", msg, err)
		return msg
	}

	return ""

}

func (u *Interventor) Load2(s string) {

	u.Nombre = s

}

func (p *Interventor) LoadDB(s string) (int, string) {

	db := fun.Acceso{}

	_, val0, msg := db.SetCliente()

	if val0 != 1 {
		return 0, msg
	}

	var sqlStatement string
	var res sql.Result

	var err error

	defer db.Cliente.Close()

	// Test the connection to the database
	if err := db.Cliente.Ping(); err != nil {
		log.Println(err)
		return 0, err.Error()
	} else {
		log.Println("\n Successfully Connected")
	}

	msg = "No se ha podido concretar el alta. "
	m := "Alta creada con exito"

	if s == "fiscalias" {

		sqlStatement = "INSERT INTO fiscalias  " +
			" (nombre) VALUES (?);"

	} else {

		sqlStatement = "INSERT INTO juzgados  " +
			" (nombre) VALUES (?);"

	}

	res, err = db.Cliente.Exec(sqlStatement, p.Nombre)

	if err != nil {

		fmt.Printf("\n %s error: %s", msg, err.Error())
		return 0, msg
	} else {
		i, _ := res.RowsAffected()
		fmt.Printf("\n count: %d", int(i))
		return 1, m
	}

}

func (p *Interventor) LoadBaja(s string, b string) (int, string) { // actualiza informacion del cliente

	db := fun.Acceso{}

	var sqlStatement string

	_, val0, msg := db.SetCliente()

	if val0 != 1 {
		return 0, msg
	}

	i, _ := strconv.Atoi(b)

	defer db.Cliente.Close()

	switch s {

	case "fiscalia":

		sqlStatement = "UPDATE fiscalias " +
			" SET is_active = 0 " +
			" WHERE id = ? ;"

	default:

		sqlStatement = "UPDATE juzgados " +
			" SET is_active = 0 " +
			" WHERE id = ? ;"

	}

	msg = "Falla al desactivar el interventor -"

	var res1 sql.Result
	var err3 error

	res1, err3 = db.Cliente.Exec(sqlStatement, i)

	if err3 != nil {

		fmt.Printf("\n %s error: %s", msg, err3.Error())
		return 0, msg

	} else {

		count, _ := res1.RowsAffected()
		fmt.Println("registros afectados al desactivar el interventor: ", count)
		return 1, "interventor deshabilitado"
	}

}
