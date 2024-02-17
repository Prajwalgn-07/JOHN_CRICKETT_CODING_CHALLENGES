package utils

import (
	"encoding/json"
	"os"
	"src/models"
)

func GetConfig() (models.Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return models.Config{}, err
	}
	defer file.Close()

	// Decode JSON from the configuration file
	var config models.Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return models.Config{}, err
	}

	return config, nil
}
