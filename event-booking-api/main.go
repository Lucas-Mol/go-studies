package main

import (
	"fmt"
	"github.com/Lucas-Mol/go-studies/event-booking-api/db"
	"github.com/Lucas-Mol/go-studies/event-booking-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
