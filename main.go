package main

import (
	"fmt"
	"log"

	//"github.com/apex/gateway"

	"github.com/apex/gateway"
	"github.com/joho/godotenv"

	"github.com/agp/push-notifications-api/routes"
)

var err error

func main() {

	if err != nil {
		fmt.Println("Status:", err)
	}

	err := godotenv.Load()
	if err != nil {

		log.Fatal("Error loading .env file")
	}

	r := routes.SetupRouter()
	//running

	// No Serverless
	//r.Run(":3000")

	// Serverless
	if err := gateway.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
