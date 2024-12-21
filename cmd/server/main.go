package main

import (
	"log"
	db "transaction/internal/adapter/persistence"

	"github.com/joho/godotenv"
)

func main() {
	err:=godotenv.Load()
	if err!=nil{
		log.Println("Error in loading the .env file")
		return
	}
	db.ConnectToDB()
}