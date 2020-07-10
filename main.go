package main

import (
	"log"
	"time"
	// "os"
	// "encoding/json"
	// "fmt"
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
	// payloadData := data{
	// 		ID: 23344,
	// 		Username: "johnmadakin",
	// 		Email: "yonolose@gmail.com",
	// 		Token: "eye4wslfjkljskdfdsjkldsfk",
	// }
	// result := payload {
	// 	Status: "success",
	// 	Message: "User login successful",
	// 	Data: payloadData,
	// }
	// // fmt.Println("---> ", result)

	// b, err := json.Marshal(result)
	// if err != nil {
	// 	fmt.Println("error", err)
	// }

	// os.Stdout.Write(b)
	// router.GET("/", func (c *gin.Context){
	// 	// jsonData := []byte(`{"msg":"this worked"}`)
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status": "success",
	// 		"message": "User login successful",
	// 		"data": payloadData,
	// 	})
	// })

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
	// http.ListenAndServe(":6001", router)
}
// func render(c *gin.Context, data gin.H, templateName string) {
// 	loggedInInterface, _ := c.Get("is_logged_in")
// 	data["is_logged_in"] = loggedInInterface.(bool)

// 	switch c.Request.Header.Get("Accept") {
// 	case "application/json":
// 		// Respond with JSON
// 	default:
// 		c.JSON(http.StatusOK, data["payload"])
// 	}
// }