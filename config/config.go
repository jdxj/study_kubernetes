package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	Ins Config
)

type Config struct {
	DB    DB    `yaml:"db"`
	Redis Redis `yaml:"redis"`
	API   API   `yaml:"api"`
}

type DB struct {
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Name string `yaml:"name"`
}

type Redis struct {
	Host  string `yaml:"host"`
	Port  string `yaml:"port"`
	Pass  string `yaml:"pass"`
	DBNum int    `yaml:"dbNum"`
}

type API struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func Init(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, &Ins)
}
