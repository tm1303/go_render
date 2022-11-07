package main

import (
	"log"

	"go_render/internal/service"
)


func main() {
	log.Println("Starting...")


	renderInterface := go_render.NewRenderInterface()
	go_render.StartGrpcServer(renderInterface)

	log.Println("Exiting...")
}