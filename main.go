package main

import (
	"fmt"
	"dragon-fruit-expert-system/configs"
	"dragon-fruit-expert-system/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	// Set env from .env
	configs.SetEnv()

	// Call routes
	r := routes.Routes()

	fmt.Println("Server listen at :" + os.Getenv("PORT"))

	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), r))
}