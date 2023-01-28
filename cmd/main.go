package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/julianojj/digital_wallet/internal/application/usecases"
	"github.com/julianojj/digital_wallet/internal/domain/service"
	"github.com/julianojj/digital_wallet/internal/infra/api/controller"
	"github.com/julianojj/digital_wallet/internal/infra/api/route"
	"github.com/julianojj/digital_wallet/internal/infra/repository/database"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()

	connectionString := os.Getenv("BASE_URL")
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	accountRepository := database.NewAccountRepositoryDatabase(db)
	transactionRepository := database.NewTransactionRepositoryDatabase(db)

	transferService := service.NewTransferService(transactionRepository)

	getAccount := usecases.NewGetAccount(accountRepository)
	transferAccounts := usecases.NewTransferAccounts(accountRepository, transferService)

	getAccountController := controller.NewGetAccountController(getAccount)
	transferAccountsController := controller.NewTransferAccountsController(transferAccounts)

	route.NewAccountRoute(
		r,
		getAccountController,
		transferAccountsController,
	).Register()

	r.Run(":3000")
}
