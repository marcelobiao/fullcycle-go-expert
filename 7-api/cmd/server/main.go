package main

import (
	"fmt"

	"github.com/marcelobiao/fullcycle-go-expert/7-api/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	fmt.Println(config)
}
