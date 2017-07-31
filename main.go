package main

import (
	"github.com/heraware/cryptodev/clients"
	"github.com/heraware/cryptodev/cmd"
)

func main() {
	defer clients.DB.Close()
	cmd.Execute()
}
