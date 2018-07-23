package utilities

import (
	"encoding/json"
	"os"

	"github.com/MaximilianoCabrera/estilohat/estilohat-backend/models"
)

func GetConfiguration() (models.Configuration, error) {
	config := models.Configuration{}
	file, err := os.Open("./configuracion.json")

	if err != nil {
		return config, nil
	}

	defer file.Close()

	decor := json.NewDecoder(file)

	err = decor.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
