package routes

import (
	"github.com/ayushmukho/ecomm-cart/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incommingRoutes *gin.Engine){
	incommingRoutes.POST("/users/signup")
	incommingRoutes.POST("/users/login")
	incommingRoutes.POST("/admin/addproducts")
	incommingRoutes.GET("/users/productview")
	incommingRoutes.GET("/users/search")
}