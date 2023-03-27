package config

import (
	"ms_logs_go/helper"
	"os"

	"gopkg.in/yaml.v3"
)

type Database struct {
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Configuration struct {
	DB Database `yaml:"Database"`
}

func LoadConfig(filename string) (*Configuration) {

	config := &Configuration{}

	content, err := os.ReadFile(filename)
	helper.ErrorPanic(err)

	err = yaml.Unmarshal(content, config)
	helper.ErrorPanic(err)

	return config

}
