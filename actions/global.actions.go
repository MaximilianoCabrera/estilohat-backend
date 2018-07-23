package actions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MaximilianoCabrera/estilohat/estilohat-backend/dao/factory"
	"github.com/MaximilianoCabrera/estilohat/estilohat-backend/dao/interfaces"
	"github.com/MaximilianoCabrera/estilohat/estilohat-backend/models"
	"github.com/MaximilianoCabrera/estilohat/estilohat-backend/utilities"
)

//Repondo uno
func response(w http.ResponseWriter, status int, results models.GlobalModel, model string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	switch model {
	case "producto":
		json.NewEncoder(w).Encode(results.Producto)
	}
}

//Responde Varios
func responses(w http.ResponseWriter, status int, results models.GlobalModels, model string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	switch model {
	case "producto":
		json.NewEncoder(w).Encode(results.Producto)
	}
}

//Func Config
func config() models.Configuration {
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}
	return config
}

//func GlobalDAO
func globalDAO() (x interfaces.GlobalDAO) {
	x = factory.GlobalFactoryDAO(config().Engine)

	return x
}

//func Errores .
func msjError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func checkErr(model string, accion string, err error) {
	if err != nil {
		switch model {
		case "producto":
			log.Println("No se pudo " + accion + " el Producto. Error: " + err)
		default:
			log.Println("Error: ", err)
		}
	}
}
