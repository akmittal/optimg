package main

import (
	"fmt"

	"github.com/akmittal/pixer/server/src/web"
)

func main() {
	err := web.Start()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Server running")
}
