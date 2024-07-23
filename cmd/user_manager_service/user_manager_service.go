package main

import (
	"Expense_Manager/pkg/auth_service/controller"
	"Expense_Manager/pkg/user_manager_service/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	authService, err := controller.NewAuth()
	if err != nil {
		log.Fatal(err)
	}

	userManagerController, err := controllers.NewUserManagerController()

	r := gin.Default()

	r.GET("/login", authService.Login)

	r.PUT("/register", userManagerController.CreateNewUser)

	s := r.Group("/user", authService.VerifyJWTToken)
	{
		s.GET("ping")
		s.POST("/change-password", userManagerController.ChangePassword)
	}

	_ = r.Run()
}
