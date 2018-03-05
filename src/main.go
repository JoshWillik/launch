package main

import (
	"fmt"
	"os"

	"./config"
	// "./project"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	settings := config.Read(dir)
	fmt.Println(settings)
	// project := project.Discover(os.Getwd(), conf)
	// project.Launch()
}
