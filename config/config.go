package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type NodeConfig struct {
	URL    string `yaml:"url"`
	MaxBPM int32  `yaml:"max_bpm"`
	MaxRPM int32  `yaml:"max_rpm"`
}

type Config struct {
	Nodes []NodeConfig `yaml:"nodes"`
	Port  string       `yaml:"port"`
}

func LoadConfig() (*Config, error) {
	file, err := os.Open("./config/config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
