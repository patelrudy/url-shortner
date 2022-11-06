package main

import (
	"urlshortner/model"
	"urlshortner/server"
)



func main() {
	model.Setup()
	server.SetupAndListen()
}