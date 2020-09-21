package main

import (
	"fmt"

	"github.com/akmittal/optimg/server/src/web"
)

func main() {
	err := web.Start()
	if err != nil {
		fmt.Println(err)
	}
}
