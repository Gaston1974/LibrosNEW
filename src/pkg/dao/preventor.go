package dao

import (
	"database/sql"
	"fmt"
	fun "hello/src/pkg/funciones"
	"log"
)

type Preventor struct {
	Nombre string
}

func (u *Preventor) Load(pp *sql.Rows) string {

	err := pp.Scan(&u.Nombre)

	if err != nil {
		msg := "\n Error en la carga del objeto"
		fmt.Printf("%s \n %s", msg, err)
		return msg
	}

	return ""

}

func (u *Preventor) Load2(s string) {

	u.Nombre = s

}

func (p *Preventor) LoadDB() (int, string) {

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

	sqlStatement = "INSERT INTO preventores  " +
		" (nombre) VALUES (?);"

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

func (p *Preventor) LoadBaja() (int, string) { // actualiza informacion del cliente

	db := fun.Acceso{}

	_, val0, msg := db.SetCliente()

	if val0 != 1 {
		return 0, msg
	}

	defer db.Cliente.Close()

	sqlStatement := "UPDATE preventores " +
		" SET is_active = 0 " +
		" WHERE nombre = ? ;"

	msg = "Falla al desactivar el preventor -"

	var res1 sql.Result
	var err3 error

	res1, err3 = db.Cliente.Exec(sqlStatement, p.Nombre)

	if err3 != nil {

		fmt.Printf("\n %s error: %s", msg, err3.Error())
		return 0, msg

	} else {

		count, _ := res1.RowsAffected()
		fmt.Println("registros afectados al desactivar el preventor: ", count)
		return 1, "preventor deshabilitado"
	}

}
