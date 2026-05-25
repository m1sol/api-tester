package main

import (
	"github.com/joho/godotenv"
	"github.com/m1sol/api-tester/internal/app"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	if err := app.Run(); err != nil {
		os.Exit(1)
	}
}
