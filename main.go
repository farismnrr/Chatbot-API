package main

import (
	"capstone-project/database"
	"capstone-project/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	client, err := database.NewRedisConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	if err := db.CreateAllTables(); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	routes.SetupUserRouter(router, db, client)
	router.Run()
}
