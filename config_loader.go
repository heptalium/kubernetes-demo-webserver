package main

import (
	"log"
	"os"

	"github.com/heptalium/envconfig"
	"gopkg.in/yaml.v2"
)

func configure() {
	configFile := os.Getenv("CONFIG_FILE")

	// Read config file
	if configFile != "" {
		log.Printf("Reading config file %s...", configFile)

		file, err := os.ReadFile(configFile)
		if err != nil {
			log.Fatalln("Could not read config file:", err)
		}

		err = yaml.Unmarshal(file, &config)
		if err != nil {
			log.Fatalln("Could not parse config file:", err)
		}
	}

	// Parse environment variables
	envconfig.MustProcess("", &config)
}
