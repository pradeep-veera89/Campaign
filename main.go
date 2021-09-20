package main

import (
	"log"
	"net/http"

	"github.com/pradeep-veera89/campaign/internal/database"
	"github.com/pradeep-veera89/campaign/routes"
)

func main() {

	routes.Routes()
	log.Println("Inside Main")
	db, err := database.ConnectSQL("host=localhost port=5432 dbname=treaction user=golang password=test123")
	if err != nil {
		log.Println("Error: Failed to connect to DB", err)
		return
	}
	defer db.SQL.Close()
	log.Println("Connected to DB")
	database.InitiateDB(&database.DB{SQL: db.SQL})
	http.ListenAndServe(":8080", nil)
}
