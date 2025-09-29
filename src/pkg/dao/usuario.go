package dao

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"strconv"

	fun "hello/src/pkg/funciones"
	"log"
)

type Usuario struct {
	Id        string
	Dato      Persona
	Ce        string
	Pass      string
	Is_active string
}

func (p *Usuario) Load(pp *sql.Rows) string {

	var t Persona

	err := pp.Scan(&p.Id, &t.Nombre, &t.Apellido, &p.Ce, &p.Is_active)
	p.Dato = t

	if err != nil {
		msg := "\n Error en la carga del objeto"
		fmt.Printf("%s \n %s", msg, err)
		return msg
	}

	return ""

}

func (p *Usuario) Load2(nom string, ape string, ce string, pass string) {

	var t Persona

	t.Nombre = nom
	t.Apellido = ape
	p.Ce = ce
	p.Pass = pass
	p.Dato = t

}

func (p *Usuario) Load3(ce string) {

	p.Ce = ce

}

func (p *Usuario) LoadDB() (int, string) {

	db := fun.Acceso{}

	_, val0, msg := db.SetCliente()

	if val0 != 1 {
		return 0, msg
	}

	defer db.Cliente.Close()

	// Test the connection to the database
	if err := db.Cliente.Ping(); err != nil {
		log.Println(err)
	} else {
		log.Println("\n Successfully Connected")
	}

	msg = "No se ha podido concretar el alta. "
	m := "Alta creada con exito"

	sqlStatement := "INSERT INTO usuarios" +
		" (username, first_name, last_name, ce, clave) " +
		" VALUES (?, ?, ?, ?, ?);"

	//hashedPass := hash.Checksum([]byte(p.Pass))
	hasher := sha256.New()
	//hasher.Write([]byte())
	hasher.Write([]byte(p.Pass))
	passHash := hasher.Sum(nil)
	hashed := hex.EncodeToString(passHash)

	username := p.Dato.Nombre + p.Ce

	res, err := db.Cliente.Exec(sqlStatement, username, p.Dato.Nombre, p.Dato.Apellido, p.Ce, hashed)

	if err != nil {
		fmt.Printf("\n %s error: %s", msg, err)
		return 0, msg

	} else {

		count, _ := res.RowsAffected()
		fmt.Println("\n registros afectados: ", count)
		fmt.Printf("\n %s ", m)

		return 1, m
	}
}

func (p *Usuario) LoadDBUpdtData(id string, ce string, nombre string, apellido string) (int, string) { // actualiza informacion del cliente

	db := fun.Acceso{}

	_, val0, msg := db.SetCliente()

	if val0 != 1 {
		return 0, msg
	}

	defer db.Cliente.Close()
	idd, _ := strconv.Atoi(id)

	sqlStatement := "UPDATE usuarios " +
		" SET first_name = '" + nombre + "' ," +
		" last_name = '" + apellido + "' , " +
		" ce = '" + ce + "' " +
		" WHERE id = ? ;"

	msg = "Falla al actualizar el usuario -"

	var res1 sql.Result
	var err3 error

	res1, err3 = db.Cliente.Exec(sqlStatement, idd)

	if err3 != nil {

		fmt.Printf("\n %s error: %s", msg, err3.Error())
		return 0, msg

	} else {

		count, _ := res1.RowsAffected()
		fmt.Println("registros afectados al actualizar el usuario: ", count)
		return 1, "Usuario actualizado"
	}

}

func (p *Usuario) LoadBaja() (int, string) { // actualiza informacion del cliente

	db := fun.Acceso{}

	_, val0, msg := db.SetCliente()

	if val0 != 1 {
		return 0, msg
	}

	defer db.Cliente.Close()

	sqlStatement := "UPDATE usuarios " +
		" SET is_active = 0 " +
		" WHERE ce = ? ;"

	msg = "Falla al eliminar el usuario -"

	var res1 sql.Result
	var err3 error

	res1, err3 = db.Cliente.Exec(sqlStatement, p.Ce)

	if err3 != nil {

		fmt.Printf("\n %s error: %s", msg, err3.Error())
		return 0, msg

	} else {

		count, _ := res1.RowsAffected()
		fmt.Println("registros afectados al eliminar el usuario: ", count)
		return 1, "usuario deshabilitado"
	}

}
