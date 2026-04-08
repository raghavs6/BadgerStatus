package main

import (
	"fmt"
	"os"
)

func main() {
	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("BadgerStatus backend starting on port %s\n", port)
}
