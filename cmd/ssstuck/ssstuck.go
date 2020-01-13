package main

import (
	"fmt"

	"github.com/mikeee/ssstuck"
)

func main() {
	fmt.Println("Starting ssstuck")
	config := ssstuck.Config{
		Port: 2222,
	}
	ssstuck.Serve(config)
}
