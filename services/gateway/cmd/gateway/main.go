package main

import (
	"fmt"
	"log"

	"github.com/alexispell/creep/gateway/internal/config"
	"github.com/alexispell/creep/gateway/internal/server"
)


func main() {
	config, err := config.LoadConfig("dev.env")
	if err != nil {
		log.Fatal("Cannot load config :>>", err)
	}
	fmt.Println("Config :>>", config)

	server := server.NewServer(config)

	server.RegisterRouter()

	server.StartServer()
}