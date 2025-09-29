package main

import (
	"log"
	"net/http"
	"os"

	handlers "hello/src/pkg/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	//"golang.org/x/crypto/acme/autocert"
)

// mongodb://[username:password@]host1[:port1][,...hostN[:portN]][/[defaultauthdb][?options]]

func main() {
	/*
		// Domain(s) for which obtain a Let's Encrypt certificate
		domains := []string{"entertainment.ddns.net"}

		// Setup autocert manager
		m := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(domains...),
			Cache:      autocert.DirCache("./certs"), // Store certificates in this directory
		}
	*/
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("No PORT variable is found on the environment")
	}
	mongoURI := os.Getenv("DBURL")
	if mongoURI == "" {
		log.Fatal("No URL variable is found on the environment")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	routerV1 := chi.NewRouter()

	// endpoints

	routerV1.Get("/html/*", handlers.HandlerReadiness)
	routerV1.Post("/login", handlers.HandlerLogIn)
	routerV1.Get("/combos/provincias", handlers.HandlerCombos)
	routerV1.Get("/combos/localidades", handlers.HandlerCombos)
	routerV1.Get("/combos/fiscalias", handlers.HandlerCombos)
	routerV1.Get("/combos/juzgados", handlers.HandlerCombos)
	routerV1.Get("/combos/preventores", handlers.HandlerCombos)
	routerV1.Post("/altas", handlers.HandlerAltas)
	routerV1.Post("/bajas", handlers.HandlerBajas)
	routerV1.Post("/password", handlers.HandlerModifyPassword)
	routerV1.Post("/causas", handlers.HandlerCausas)           // consulta all o por numero de causa
	routerV1.Post("/causas/historico", handlers.HandlerCausas) // consulta por numero de causa
	routerV1.Post("/usuarios", handlers.HandlerUsuarios)       // consulta all o por ce
	routerV1.Post("/updates", handlers.HandlerUpdates)
	//routerV1.Get("/token", handlers.HandlerToken)
	router.Mount("/v1", routerV1)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
		//	TLSConfig: &tls.Config{
		//		GetCertificate: m.GetCertificate,
		//	},
	}

	log.Printf("Server starting on port %v", portString)
	/*
		err := srv.ListenAndServeTLS("", "")
		if err != nil {
			log.Fatal(err)
		}
	*/
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
