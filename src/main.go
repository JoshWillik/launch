package main

import (
	"os"

	"./config"
	"./project"
)

func main() {
	settings := config.ReadOrCreateSettings()
	project := project.Discover(os.Getwd(), conf)
	project.Launch()
}
