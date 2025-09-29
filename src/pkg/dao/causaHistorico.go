package dao

import (
	"database/sql"
	"fmt"
)

type Causa_historico struct {
	NroCausa    string
	Accion      string
	Usuario_id  string
	Descripcion string
	IpAddress   string
	CreadoEn    string
}

func (u *Causa_historico) Load(pp *sql.Rows) string {

	err := pp.Scan(&u.NroCausa, &u.Accion, &u.Usuario_id, &u.Descripcion, &u.IpAddress, &u.CreadoEn)

	if err != nil {
		msg := "\n Error en la carga del objeto"
		fmt.Printf("%s \n %s", msg, err)
		return msg
	}

	return ""

}
