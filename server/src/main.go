package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akmittal/optimg/server/src/web"
	"github.com/joho/godotenv"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	fmt.Println(os.Getenv("IMAGE_PATH"))
	err := web.Start()
	if err != nil {
		fmt.Println(err)
	}
}
