package main

import (
	"log"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	routes "github.com/ecommerce/routes"
	config "github.com/ecommerce/config"
	models "github.com/ecommerce/models"

)

var router *gin.Engine

type data struct {
	ID int32 `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type payload struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data data `json:"data"`
}

func main(){
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)
	db := config.Connect()

	//create tables
	models.CreateTables(db)
	models.InitiateDB(db)
	// Set the router as the default one provided by Gin
	router = gin.Default()

	routes.Routes(router)

	// Start serving the application
	srv := &http.Server{
		Addr:	":4001",
		Handler: router,
		ReadTimeout: 20 * time.Second,
		WriteTimeout: 20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(srv.ListenAndServe())
}
