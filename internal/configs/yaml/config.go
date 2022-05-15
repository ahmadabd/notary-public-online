package yaml

import (
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		Server   Server   `yaml:"server"`
		Database Database `yaml:"database"`
		App      App      `yaml:"app"`
	}

	Server struct {
		Port string `yaml:"port"`
	}

	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	}

	App struct {
		Mode string `yaml:"mode"`
	}
)

func GetConfig(path string) (*Config, error) {
	switch filepath.Ext(path) {
	case ".yaml", ".yml":
		return parseYML(path), nil
	default:
		return nil, errors.New("unknown file extension")
	}
}

func parseYML(path string) *Config {
	c := &Config{}

	yamlFile, err := ioutil.ReadFile(path)

	if err != nil {
		log.Println("Error reading YAML file: ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		log.Println("Error parsing YAML file: ", err)
	}

	return c
}
