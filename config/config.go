package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"app"`

	Database struct {
		Mysql struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			Dbname   string `yaml:"dbname"`
		} `yaml:"mysql"`

		Redis struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			Password string `yaml:"password"`
			Db       int    `yaml:"db"`
		} `yaml:"redis"`
	} `yaml:"database"`
}

var cfg *Config

func init() {
	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("Error opening config file: %s", err)
	}

	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		log.Fatalf("Error unmarshalling config file: %s", err)
	}
}

func GetConfig() *Config {
	return cfg
}
