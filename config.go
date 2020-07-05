package main

import (
	"encoding/json"
	"os"
	"path"
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

// GeneratePath generates path from the given directory with the config specified
func GeneratePath(config *Configuration, dir string) (string, string) {
	return path.Join(dir, config.Input), path.Join(dir, config.Output)
}
