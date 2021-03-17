package controllers

import (
	"io/ioutil"
	"log"

	"code.unitiwireless.com/uniti-wireless/harepd/models"
	"gopkg.in/yaml.v2"
)

//ReadConf to fetch the configuration
func ReadConf(conf string, c *models.Config) {
	yamlFile, err := ioutil.ReadFile(conf)

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		panic(err)
	}
}
