package main

import (
	"net/http"

	"github.com/MaximilianoCabrera/estilohat/estilohat-backend/routes"
	"github.com/rs/cors"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
	})

	router := routes.NewMainRouter()

	http.ListenAndServe(":8080", c.Handler(router))
}
