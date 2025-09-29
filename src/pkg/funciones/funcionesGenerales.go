package funciones

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Puntuaciones struct {
	Nombre   string
	Apellido string
	Puntos   float64
	Fecha    string
}

func LeerArchivo(ruta string, w http.ResponseWriter) {

	f, err := os.Open(ruta)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// read 1024 bytes at a time
	buf := make([]byte, 1024)

	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			// there is no more data to read
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {
			//f.Write(buf[:n])
			//fmt.Println("leyendo ", buf[:n])
			w.Write(buf[:n])
		}
	}

}

func WriteJson(payload interface{}) (int, string) {

	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to Marshal JSON file: %v ", payload)
		return 0, ""
	}

	return 1, string(dat)
}

/*
func LeerHistorico(historico []bson.M, w http.ResponseWriter) {

	w.Header().Add("Content-Type", "text")
	w.WriteHeader(200)

	type prediccion struct {
		fecha      string
		porcentaje string
	}

	l := len(historico)
	resp := make([]prediccion, l)
	var cadena string

	for key, value := range historico {

		resp[key].fecha = string(value["fecha"].(string))
		resp[key].porcentaje = strconv.FormatFloat(value["porcentaje"].(float64), 'f', 2, 64)

		cadena += resp[key].fecha + "," + resp[key].porcentaje + ","

	}

	dat := []byte(cadena)

	w.Write(dat)

}

func Ordenar(p *[]Puntuaciones) {

	// Ordeno Resultados :

	for i := len(*p) - 1; i > 0; i-- {

		var aux Puntuaciones
		g := *p

		for j := 0; j < i; j++ {

			if g[j].Puntos > g[j+1].Puntos {
				aux = g[j+1]
				g[j+1] = g[j]
				g[j] = aux

			} else if g[j].Puntos == g[j+1].Puntos {

				if g[j].Fecha > g[j+1].Fecha {
					aux = g[j+1]
					g[j+1] = g[j]
					g[j] = aux

				}

			}

		}

	}

}


*/
