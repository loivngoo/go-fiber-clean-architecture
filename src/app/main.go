package main

import (
	accountmodel "clean-architecture-go-fiber/src/module/account/model"
	accounttrpt "clean-architecture-go-fiber/src/module/account/transport"
	database "clean-architecture-go-fiber/src/platform/database/mysql"

	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// router.SetTrustedProxies([]string{"192.168.88.55"})
	// gin.SetMode(gin.ReleaseMode)

	db, err := database.OpenConnectToMySQL()

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}
	fmt.Println("Opened connection to MySQL")

	db.AutoMigrate(&accountmodel.Account{})

	v1 := router.Group("/api/v1")
	{
		v1.POST("/accounts", accounttrpt.HandleCreateAccount(db))
	}

	router.Run(":8080")
}
