package main

import (
	"log"

	"github.com/akmittal/optimg/server/pkg/application"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load env vars")
	}

	app, err := application.Get()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = app.Start()
	if err != nil {
		log.Fatal(err.Error())
	}

}
