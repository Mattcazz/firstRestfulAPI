package main

import (
	"fmt"
	"workspace/db"
	"workspace/routes"

	"github.com/gin-gonic/gin"
	//"errors"
)

func main() {
	db.Init()
	defer db.CloseDB()

	db.MigrateDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	fmt.Println("Running server on port 8080")
	r.Run("localhost:8080")

}
