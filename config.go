package main

import (
	"io/ioutil"
	yaml "gopkg.in/yaml.v2"
	"log"
)

type Config struct {
	Length	int
	Fps	int
	Server
}

type Server struct {
	Host	string
	Port	int
}

func readConfigFile() []byte {
	config, err := ioutil.ReadFile("./config.yaml")

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return config
}

func getConfig() Config {
	file := readConfigFile()

	config := Config{}

	err := yaml.Unmarshal(file, &config)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return config
}
