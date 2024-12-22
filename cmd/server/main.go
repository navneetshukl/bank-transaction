package main

import (
	"log"
	db "transaction/internal/adapter/persistence"
	routes "transaction/internal/interface"
	userhandler "transaction/internal/interface/handler/user"
	"transaction/internal/usecase/user"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error in loading the .env file")
		return
	}
	DB, err := db.ConnectToDB()
	if err != nil {
		log.Println("Error connecting to the database")
		return
	}

	dbUseCase := db.NewDatabase(DB)
	userUseCase := user.NewUserInterfaceImpl(dbUseCase)
	userHandler := userhandler.NewUserHandler(userUseCase)
	routes := routes.SetupRoutes(userHandler)
	err = routes.Listen(":8080")
	if err != nil {
		log.Println("Error in listening to the server")
		return
	}
}
