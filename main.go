package main

import (
	"log"

	"capstone-project/database"
	"capstone-project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create all tables in the database
	if err := db.CreateAllTables(); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	routes.SetupUserRouter(router, db)
	router.Run()
}
