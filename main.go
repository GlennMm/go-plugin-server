package main

import (
	"web_api_engine/engine"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Failed to load .env file.")
	}

	system := engine.NewEngine()
	system.Run()
}
