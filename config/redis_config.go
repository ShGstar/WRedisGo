package config

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
)

type Yaml struct {
	Redis struct {
		Address       string `yaml:"address"`
		ServerAddress string `yaml:"address"`
		Password      string `yaml:"password"`
		DBNum         int    `yaml:"dbnum"`
		ThreadsNum    int    `yaml:"threadnum"`
		MaxActive     int    `yaml:"maxactive"`
		MaxIdle       int    `yaml:"maxidle"`
	}
}

var (
	Conf = new(Yaml)
)

func ConfigInit() {

	yamlFile, err := ioutil.ReadFile("config.yaml")

	log.Println("yamlFile:", yamlFile)
	if err != nil {
		log.Print("yamkFile.Get err %#v", err)
	}
	err = yaml.Unmarshal(yamlFile, Conf)
	if err != nil {
		log.Print("Unmarshal.Get err %#v", err)
	}

	log.Println("conf", Conf)
}
