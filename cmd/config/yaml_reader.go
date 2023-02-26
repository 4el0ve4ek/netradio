package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	DefaultYAMLPath = "config.yaml"
)

func ReadYAML(filePath string) (Config, error) {
	config := &Config{}
	file, err := os.Open(filePath)
	fmt.Println()
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return Config{}, err
	}

	return *config, nil
}
