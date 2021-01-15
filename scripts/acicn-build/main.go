package main

import (
	"log"
	"os"

	"github.com/acicn/library"
)

func main() {
	var err error
	defer func(err *error) {
		if *err != nil {
			log.Println("exited with error:", (*err).Error())
			os.Exit(1)
		}
	}(&err)

	for _, task := range library.Pulls {
		log.Printf("Pull: %+v", task)
	}
	for _, task := range library.Builds {
		log.Printf("Build: %+v", task)
	}
	for _, task := range library.Mirrors {
		log.Printf("Mirror: %+v", task)
	}
}
