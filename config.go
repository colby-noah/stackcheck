package main

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
)

type Target struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

type Config struct {
	Targets []Target `yaml:"targets"`
}

func loadConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("reading %s: %w", path, err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, fmt.Errorf("parsing %s: %w", path, err)
	}

	return config, nil
}
