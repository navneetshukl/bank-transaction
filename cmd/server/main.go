package main

import (
	"log"
	db "transaction/internal/adapter/persistence"
	routes "transaction/internal/interface"
	bankhandler "transaction/internal/interface/handler/bank"
	userhandler "transaction/internal/interface/handler/user"
	"transaction/internal/usecase/bank"
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

	dbUseCase := db.NewUserDatabase(DB)
	userUseCase := user.NewUserInterfaceImpl(dbUseCase)
	userHandler := userhandler.NewUserHandler(userUseCase)

	bankDBUsecase := db.NewBankDatabase(DB)
	bankUseCase := bank.NewBankInterfaceImpl(dbUseCase, bankDBUsecase)
	bankHandler := bankhandler.NewBankHandler(bankUseCase)
	routes := routes.SetupRoutes(userHandler, bankHandler)
	err = routes.Listen(":8080")
	if err != nil {
		log.Println("Error in listening to the server")
		return
	}
}
