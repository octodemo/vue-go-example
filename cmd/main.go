package main

import (
	"os"

	"github.com/ndabAP/vue-go-example/internal/routes"
)

const defaultPort string = "3000"

func main() {
	port := os.Getenv("PORT")
	os.
	if port == "" {
		port = defaultPort
	}

	r := routes.SetupRoutes()

	r.Run(":" + port)
}
