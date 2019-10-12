package main

import (
	"restgin/atur"
	"restgin/basecon"

	"github.com/gin-gonic/gin"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/login", atur.Login)
		auth.POST("/signup", atur.CreateUser)
	}
	prod := router.Group("/v1/products")
	{
		prod.POST("/", basecon.Auth, atur.CreateProduct)
		prod.GET("/", basecon.Auth, atur.FetchAllProduct)
		prod.GET("/:id", basecon.Auth, atur.FetchSingleProduct)
		prod.PUT("/:id", basecon.Auth, atur.UpdateProduct)
		prod.DELETE("/:id", basecon.Auth, atur.DeleteProdcut)
	}
	prod2 := router.Group("/v2/products")
	{
		prod2.GET("/", basecon.Auth, atur.ProductMessage)
	}
	router.Run()

}
