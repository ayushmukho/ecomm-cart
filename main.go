package main

import (
	"log"
	"os"

	"github.com/ayushmukho/ecomm-cart/controllers"
	"github.com/ayushmukho/ecomm-cart/database"
	"github.com/ayushmukho/ecomm-cart/middleware"
	"github.com/gin-gonic/gin"
) 

func main(){
	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.new()
	router.Use(gin.Logger())

	router.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveIten())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}