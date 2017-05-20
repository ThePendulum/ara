package main

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Length int
	Fps    int
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
