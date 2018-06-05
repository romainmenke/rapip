package main

import (
	"os"

	"github.com/romainmenke/rapip/server"
)

// TODO : cli

func main() {

	server.Run(server.Config{Port: os.Getenv("PORT")})

}
