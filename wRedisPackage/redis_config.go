package wRedisPackage

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
)

type Yaml struct {
	Redis struct {
		Address     string `yaml:"Address"`
		Password    string `yaml:"Password"`
		DBNum       int    `yaml:"DBNum"`
		ThreadsNum  int    `yaml:"ThreadsNum"`
		MaxActive   int    `yaml:"MaxActive"`
		MaxIdle     int    `yaml:"MaxIdle"`
		Idletimeout int    `yaml:"Idletimeout"`
	} `yaml:"Redis"`
}

var (
	Conf = new(Yaml)
)

func ConfigInit() {

	yamlFile, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		log.Print("yamkFile.Get err ", err)
		return
	}
	err = yaml.Unmarshal(yamlFile, Conf)
	if err != nil {
		log.Print("Unmarshal.Get err ", err)
		return
	}

	log.Println("conf", Conf)
}
