package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server Server `yaml:"server"`
	Mysql  Mysql  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Mysql struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	DatabaseName string `yaml:"databaseName"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

var AppConfig *Config

func InitConfig(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln("init config failed: ", err)
	}

	if err := yaml.Unmarshal(data, &AppConfig); err != nil {
		log.Fatalln("yaml unmarshal failed ", err)
	}

	log.Println("init config success")
}
