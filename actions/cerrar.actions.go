package actions

import (
	"log"
	"net/http"
)

func Close(w http.ResponseWriter, r *http.Request) {
	log.Fatalln("Cerrando App Estilo-Hat")
}
