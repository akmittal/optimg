package main

import (
	"log"

	"github.com/akmittal/optimg/server/cmd"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load env vars")
	}
	cmd.Execute()

}
