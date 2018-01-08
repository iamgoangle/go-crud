package config

import (
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type conf struct {
	host string `yaml:"host"`
	port string `yaml:"port"`
}

func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("./config.dev.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {
	var c conf
	c.getConf()
	fmt.Println(c)
}
