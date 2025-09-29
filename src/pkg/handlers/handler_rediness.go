package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	funciones "hello/src/pkg/funciones"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	m := "Falla interna en parseo de archivo"
	dominio := os.Getenv("DOMINIO")

	resp := funciones.Respuesta{}
	resp.Msg = m

	vector := strings.Split(path, "/") // se asumen todos los archivos dentro de la carpeta "html"
	value1 := vector[3]

	//fmt.Printf("path:  %s \n vector: %v \n longitud: %d \n", path, vector, len(vector))

	if len(vector) == 4 {

		switch value1 {

		case "dashboard":

			w.WriteHeader(200)
			t, err := template.ParseFiles("./src/Frontend/build/html/dashboard.html")
			if err != nil {
				fmt.Printf("%s %s", m, err)
				funciones.ResponseWithJSON(w, 400, resp)
			}
			t.Execute(w, nil)

		case "register":

			w.WriteHeader(200)
			t, err := template.ParseFiles("./src/Frontend/build/html/register.html")
			if err != nil {
				fmt.Printf("%s %s", m, err)
				funciones.ResponseWithJSON(w, 400, resp)
			}
			t.Execute(w, nil)

		case "index":

			w.Header().Set("Set-Cookie", "dominio="+dominio+"; SameSite=Strict")
			w.WriteHeader(200)
			t, err := template.ParseFiles("./src/Frontend/build/html/index.html")
			if err != nil {
				fmt.Printf("%s %s", m, err)
				funciones.ResponseWithJSON(w, 400, resp)
			}
			t.Execute(w, nil)

		case "result":

			w.WriteHeader(200)
			t, err := template.ParseFiles("./src/Frontend/build/html/result.html")
			if err != nil {
				fmt.Printf("%s %s", m, err)
				funciones.ResponseWithJSON(w, 400, resp)
			}
			t.Execute(w, nil)

		}

	} else if len(vector) == 5 {
		value := vector[4]

		switch value1 {

		case "img":
			w.Header().Set("Content-Type", "image/png")
			w.WriteHeader(200)

			switch value {

			case "flag.png":

				funciones.LeerArchivo("./src/Frontend/build/html/img/flag.png", w)

			case "Premio.png":

				funciones.LeerArchivo("./src/Frontend/build/html/img/Premio.png", w)

			}

		case "css":
			w.Header().Set("Content-Type", "text/css")
			w.WriteHeader(200)
			t, err := template.ParseFiles("./src/Frontend/build/html/css/style.css")
			if err != nil {
				fmt.Printf("%s %s", m, err)
				funciones.ResponseWithJSON(w, 400, resp)
				return
			}
			t.Execute(w, nil)

		case "js":
			w.Header().Set("Content-Type", "application/javascript")
			w.WriteHeader(200)

			switch value {

			case "login.js":

				t, err := template.ParseFiles("./src/Frontend/build/html/js/login.js")
				if err != nil {
					fmt.Printf("%s %s", m, err)
					funciones.ResponseWithJSON(w, 400, resp)
					return
				}
				t.Execute(w, nil)

			case "main.js":

				t, err := template.ParseFiles("./src/Frontend/build/html/js/main.js")
				if err != nil {
					fmt.Printf("%s %s", m, err)
					funciones.ResponseWithJSON(w, 400, resp)
					return
				}
				t.Execute(w, nil)

			case "register.js":

				t, err := template.ParseFiles("./src/Frontend/build/html/js/register.js")
				if err != nil {
					fmt.Printf("%s %s", m, err)
					funciones.ResponseWithJSON(w, 400, resp)
				}
				t.Execute(w, nil)

			case "password.js":

				t, err := template.ParseFiles("./src/Frontend/build/html/js/password.js")
				if err != nil {
					fmt.Printf("%s %s", m, err)
					funciones.ResponseWithJSON(w, 400, resp)
				}
				t.Execute(w, nil)

			case "predictions.js":

				t, err := template.ParseFiles("./src/Frontend/build/html/js/predictions.js")
				if err != nil {
					fmt.Printf("%s %s", m, err)
					funciones.ResponseWithJSON(w, 400, resp)
				}
				t.Execute(w, nil)

			}
		}

	} else if len(vector) == 6 {

		w.Header().Set("Content-Type", "text/x-scss")
		w.WriteHeader(200)
		t, err := template.ParseFiles("./src/Frontend/build/html/ratchet-2.0.2/sass/ratchicons.scss")
		if err != nil {
			fmt.Printf("%s %s", m, err)
			funciones.ResponseWithJSON(w, 400, resp)
		}
		t.Execute(w, nil)

	} else if len(vector) == 7 {
		value := vector[6]

		switch value1 {

		case "ratchet-2.0.2":

			switch value {

			case "ratchet.css":

				w.Header().Set("Content-Type", "text/css")
				w.WriteHeader(200)
				t, err := template.ParseFiles("./src/Frontend/build/html/ratchet-2.0.2/dist/css/ratchet.css")
				if err != nil {
					fmt.Printf("%s %s", m, err)
					funciones.ResponseWithJSON(w, 400, resp)
				}
				t.Execute(w, nil)

			case "ratchicons.woff":

				w.Header().Set("Content-Type", "font/woff")
				w.WriteHeader(200)
				t, err := template.ParseFiles("./src/Frontend/build/html/ratchet-2.0.2/dist/fonts/ratchicons.woff")
				if err != nil {
					fmt.Printf("%s %s", m, err)
					funciones.ResponseWithJSON(w, 400, resp)
				}
				t.Execute(w, nil)

			case "ratchicons.ttf":

				w.Header().Set("Content-Type", "font/ttf")
				w.WriteHeader(200)
				t, err := template.ParseFiles("./src/Frontend/build/html/ratchet-2.0.2/dist/fonts/ratchicons.ttf")
				if err != nil {
					fmt.Printf("%s %s", m, err)
					funciones.ResponseWithJSON(w, 400, resp)
				}
				t.Execute(w, nil)

			case "ratchicons.svg":

				w.Header().Set("Content-Type", "image/svg+xml")
				w.WriteHeader(200)
				t, err := template.ParseFiles("./src/Frontend/build/html/ratchet-2.0.2/dist/fonts/ratchicons.svg")
				if err != nil {
					fmt.Printf("%s %s", m, err)
					funciones.ResponseWithJSON(w, 400, resp)
				}
				t.Execute(w, nil)

			case "ratchicons.eot":

				w.Header().Set("Content-Type", "application/vnd.ms-fontobject")
				w.WriteHeader(200)
				t, err := template.ParseFiles("./src/Frontend/build/html/ratchet-2.0.2/dist/fonts/ratchicons.eot")
				if err != nil {
					fmt.Printf("%s %s", m, err)
					funciones.ResponseWithJSON(w, 400, resp)
				}
				t.Execute(w, nil)

			case "ratchet-theme-ios.css":

				w.Header().Set("Content-Type", "text/css")
				w.WriteHeader(200)
				t, err := template.ParseFiles("./src/Frontend/build/html/ratchet-2.0.2/dist/css/ratchet-theme-ios.css")
				if err != nil {
					fmt.Printf("%s %s", m, err)
					funciones.ResponseWithJSON(w, 400, resp)
				}
				t.Execute(w, nil)

			case "ratchet.js":

				w.Header().Set("Content-Type", "application/javascript")
				w.WriteHeader(200)
				t, err := template.ParseFiles("./src/Frontend/build/html/ratchet-2.0.2/dist/js/ratchet.js")
				if err != nil {
					fmt.Printf("%s %s", m, err)
					funciones.ResponseWithJSON(w, 400, resp)
				}
				t.Execute(w, nil)

			}
		}

	} else {

		fmt.Printf("URL incorrecta")
		resp.Msg = "URL incorrecta"
		funciones.ResponseWithJSON(w, 400, resp)

	}

}
