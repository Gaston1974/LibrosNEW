package dao

import (
	"database/sql"
	"fmt"
)

type Persona struct {
	Nombre   string
	Apellido string
}

func (u *Persona) Load(pp *sql.Rows) string {

	err := pp.Scan(&u.Nombre, &u.Apellido)

	if err != nil {
		msg := "\n Error en la carga del objeto"
		fmt.Printf("%s \n %s", msg, err)
		return msg
	}

	return ""

}
