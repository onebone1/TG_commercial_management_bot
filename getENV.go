package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load("./example.env")
	if err != nil {
		log.Fatal("Error ;padding .env file")
	}

	key1 := os.Getenv("key1")
	log.Printf("%t\n", os.Getenv("key2"))
	key2, err := strconv.ParseInt(os.Getenv("key2"), 10, 64)
	if err != nil {
		log.Println(err)
	}
	log.Println(key1)
	log.Printf("%t\n", key2)
}
