package main

import (
	"log"

	"Netology/props"
	"Netology/pythagoreancalculation/handler"
	"Netology/server"
)

func main() {
	startApplication()
}

func startApplication() {
	newCalculate := handler.NewCalculateDomainHandler()
	properties, err := props.ReadProperties("./env/application.yaml")
	if err != nil {
		log.Fatalf("Error reading configurations file: %v", err)
	}

	servers := server.NewServer(properties)
	servers.ConfigureAPI(newCalculate)
}
