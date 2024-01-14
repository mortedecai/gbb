package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// GoBurnBits is the configuration structure for the GoBurnBits command line utility.
type GoBurnBits struct {
	// Host is the name of the host the BitBurner game is found on
	Host string `yaml:"host" json:"host"`
	// Port is the port of the host the BitBurner game is found on
	Port int `yaml:"port" json:"port"`
	// AuthToken is the token used to authenticate on any interactions with the BitBurner server
	AuthToken string `yaml:"authToken,omitempty" json:"authToken,omitempty"`
}

func Default() *GoBurnBits {
	return &GoBurnBits{Host: "localhost", Port: 9990}
}

func FromConfig() (*GoBurnBits, error) {
	cfg := Default()
	if !cfg.ConfigFileExists() {
		return cfg, fmt.Errorf("config file doesn't exist")
	}
	// Err is checked as part of ConfigFileExists
	fp, _ := cfg.configFilePath()
	data, err := os.ReadFile(fp)
	if err != nil {
		return cfg, fmt.Errorf("could not read file %s", fp)
	}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		err = fmt.Errorf("could not parse YAML file: %s", err.Error())
	}
	return cfg, err
}

func (gbb *GoBurnBits) configFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/.config/gbb.yaml", home), nil
}

// ConfigFileExists returns whether or not a configuration file exists
func (gbb *GoBurnBits) ConfigFileExists() bool {
	fp, err := gbb.configFilePath()
	if err == nil {
		_, err = os.Stat(fp)
	}
	return err == nil
}
