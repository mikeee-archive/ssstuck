package main

import (
	"fmt"

	"github.com/mikeee/ssstuck"
)

func main() {
	config := ssstuck.Config{
		Port: 2222,
	}
	fmt.Println("Starting ssstuck on port:" + config.Port)
	ssstuck.Serve(config)
}
