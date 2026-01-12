package main

import (
	"log"
	"os"
)

func setupDataDir() {
	if config.DataDir != "" {
		err := os.MkdirAll(config.DataDir, 0777)
		if err != nil {
			log.Fatalln("Cannot create data directory:", err)
		}
	}
}
