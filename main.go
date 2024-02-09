package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	token := os.Getenv("TOKEN")
	fmt.Println(token)
}
