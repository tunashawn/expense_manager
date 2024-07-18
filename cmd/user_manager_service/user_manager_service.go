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

	userManager, err := controllers.NewUserManager()

	r := gin.Default()

	r.GET("/login", authService.Login)

	r.PUT("/register", userManager.CreateNewUser)

	s := r.Group("/user", authService.VerifyJWTToken)
	{
		s.GET("ping")
		s.POST("/change-password", userManager.ChangePassword)
	}

	_ = r.Run()
}
