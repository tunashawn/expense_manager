package main

import (
	"Expense_Manager/commons/db"
	"Expense_Manager/pkg/auth_service/controller"
	"Expense_Manager/pkg/wallet_service/controllers"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"log"
)

func main() {

	mysqlDB, err := db.NewMySQLConnection()
	if err != nil {
		log.Fatal(errors.Wrap(err, "could not connect to MySQL"))
	}

	walletController, err := controllers.NewWalletController(mysqlDB)
	if err != nil {
		log.Fatal(err)
	}

	transactionController, err := controllers.NewTransactionController(mysqlDB)
	if err != nil {
		log.Fatal(err)
	}

	authService, err := controller.NewAuth()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.Use(authService.VerifyJWTToken)

	w := r.Group("/wallet")
	{
		w.POST("/create", walletController.CreateNewWallet)
		w.GET("/:id", walletController.GetWallet)
		w.GET("/list", walletController.GetWalletList)
		w.PUT("/update", walletController.UpdateWallet)
		w.DELETE("/delete", walletController.DeleteWallet)
	}

	t := r.Group("/transaction")
	{
		t.POST("/create", transactionController.CreateNewTransaction)
		t.GET("/:id", transactionController.GetTransaction)
		t.GET("/list", transactionController.GetTransactionList)
		t.PUT("/update", transactionController.UpdateTransaction)
		t.DELETE("/delete/:id", transactionController.DeleteTransaction)
	}

	_ = r.Run()
}
