package main

import (
	"encoding/json"
	"os"
)

// Configuration for expected input file
type Configuration struct {
	Input  string
	Output string
}

// ReadConfig get read configuration input and output filename
func ReadConfig() (*Configuration, error) {
	file, _ := os.Open("config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Configuration{}
	err := decoder.Decode(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
